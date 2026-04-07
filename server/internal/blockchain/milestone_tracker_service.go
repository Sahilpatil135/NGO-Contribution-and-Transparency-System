package blockchain

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"

	"server/internal/blockchain/contracts"
)

// MilestoneTrackerView is on-chain milestone tracking state for a cause.
type MilestoneTrackerView struct {
	ContractAddress string `json:"contract_address"`
	Goal            string `json:"goal"`
	Collected       string `json:"collected"`
	MilestonesPaid  uint8  `json:"milestones_paid"`
	Exists          bool   `json:"exists"`
	MilestonePct    int    `json:"milestone_percent_each"` // 25% increments
}

type MilestoneTrackerService struct {
	client   *Client
	contract *contracts.MilestoneTracker
	address  common.Address
}

func NewMilestoneTrackerService(client *Client, contractAddress string) (*MilestoneTrackerService, error) {
	if contractAddress == "" {
		return nil, fmt.Errorf("milestone tracker contract address is empty")
	}
	addr := common.HexToAddress(contractAddress)
	instance, err := contracts.NewMilestoneTracker(addr, client.EthClient)
	if err != nil {
		return nil, err
	}
	return &MilestoneTrackerService{
		client:   client,
		contract: instance,
		address:  addr,
	}, nil
}

func (s *MilestoneTrackerService) AddressHex() string {
	return s.address.Hex()
}

// RegisterCause registers a cause with its funding goal and initial collected amount
func (s *MilestoneTrackerService) RegisterCause(
	ctx context.Context,
	causeID uuid.UUID,
	goal *big.Int,
	initialCollected *big.Int,
) (string, error) {
	auth, err := s.client.NewTransactor(ctx)
	if err != nil {
		return "", err
	}

	tx, err := s.contract.RegisterOrUpdateCause(auth, UUIDToBytes16(causeID), goal, initialCollected)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

// GetCauseState reads milestone tracking state from the chain
func (s *MilestoneTrackerService) GetCauseState(ctx context.Context, causeID uuid.UUID) (*MilestoneTrackerView, error) {
	st, err := s.contract.GetCause(&bind.CallOpts{Context: ctx}, UUIDToBytes16(causeID))
	if err != nil {
		return nil, err
	}

	return &MilestoneTrackerView{
		ContractAddress: s.address.Hex(),
		Goal:            bigintStr(st.Goal),
		Collected:       bigintStr(st.Collected),
		MilestonesPaid:  st.MilestonesPaid,
		Exists:          st.Exists,
		MilestonePct:    25,
	}, nil
}

// IsCauseRegistered checks if a cause is already registered
func (s *MilestoneTrackerService) IsCauseRegistered(ctx context.Context, causeID uuid.UUID) (bool, error) {
	state, err := s.GetCauseState(ctx, causeID)
	if err != nil {
		return false, err
	}
	return state.Exists, nil
}

// EnsureCauseRegistered checks if cause is registered, and registers it if not
func (s *MilestoneTrackerService) EnsureCauseRegistered(
	ctx context.Context,
	causeID uuid.UUID,
	goal *big.Int,
	initialCollected *big.Int,
) error {
	// Check if already registered
	registered, err := s.IsCauseRegistered(ctx, causeID)
	if err != nil {
		return fmt.Errorf("failed to check registration: %w", err)
	}

	if registered {
		// Already registered, nothing to do
		return nil
	}

	// Not registered, register it now with initial collected amount
	txHash, err := s.RegisterCause(ctx, causeID, goal, initialCollected)
	if err != nil {
		return fmt.Errorf("failed to register cause: %w", err)
	}

	log.Printf("Registered cause %v with goal=%s, initialCollected=%s (tx: %s)", 
		causeID, goal.String(), initialCollected.String(), txHash)

	return nil
}

// RecordDonation records a donation amount for milestone tracking
// This does NOT transfer ETH, just updates the collected amount
func (s *MilestoneTrackerService) RecordDonation(
	ctx context.Context,
	causeID uuid.UUID,
	amount *big.Int,
) (string, error) {
	auth, err := s.client.NewTransactor(ctx)
	if err != nil {
		return "", err
	}

	// Call recordDonation on the tracker contract (no ETH sent)
	tx, err := s.contract.RecordDonation(auth, UUIDToBytes16(causeID), amount)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

// IsMilestoneReached checks if a specific milestone has been reached
func (s *MilestoneTrackerService) IsMilestoneReached(ctx context.Context, causeID uuid.UUID, milestone uint8) (bool, error) {
	reached, err := s.contract.IsMilestoneReached(&bind.CallOpts{Context: ctx}, UUIDToBytes16(causeID), milestone)
	if err != nil {
		return false, err
	}
	return reached, nil
}

func bigintStr(v *big.Int) string {
	if v == nil {
		return "0"
	}
	return v.String()
}
