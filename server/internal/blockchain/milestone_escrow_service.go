package blockchain

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"

	"server/internal/blockchain/contracts"
)

// EscrowCauseView is on-chain milestone escrow state for a cause.
type EscrowCauseView struct {
	ContractAddress string   `json:"contract_address"`
	GoalWei         string   `json:"goal_wei"`
	CollectedWei    string   `json:"collected_wei"`
	ReleasedWei     string   `json:"released_wei"`
	MilestonesPaid  uint64   `json:"milestones_paid"`
	Beneficiary     string   `json:"beneficiary"`
	Exists          bool     `json:"exists"`
	MilestonePct    int      `json:"milestone_percent_each"` // 25% increments
}

type MilestoneEscrowService struct {
	client   *Client
	contract *contracts.CauseMilestoneEscrow
	address  common.Address
}

func NewMilestoneEscrowService(client *Client, contractAddress string) (*MilestoneEscrowService, error) {
	if contractAddress == "" {
		return nil, fmt.Errorf("milestone escrow contract address is empty")
	}
	addr := common.HexToAddress(contractAddress)
	instance, err := contracts.NewCauseMilestoneEscrow(addr, client.EthClient)
	if err != nil {
		return nil, err
	}
	return &MilestoneEscrowService{
		client:   client,
		contract: instance,
		address:  addr,
	}, nil
}

func (s *MilestoneEscrowService) AddressHex() string {
	return s.address.Hex()
}

// RegisterCause registers the cause on-chain (platform owner key). Beneficiary receives milestone payouts.
func (s *MilestoneEscrowService) RegisterCause(
	ctx context.Context,
	causeID uuid.UUID,
	goalWei *big.Int,
	beneficiary common.Address,
) (string, error) {
	auth, err := s.client.NewTransactor(ctx)
	if err != nil {
		return "", err
	}

	tx, err := s.contract.RegisterCause(auth, UUIDToBytes16(causeID), goalWei, beneficiary)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

// GetCauseState reads escrow balances and milestone progress from the chain.
func (s *MilestoneEscrowService) GetCauseState(ctx context.Context, causeID uuid.UUID) (*EscrowCauseView, error) {
	st, err := s.contract.GetCause(&bind.CallOpts{Context: ctx}, UUIDToBytes16(causeID))
	if err != nil {
		return nil, err
	}

	milestonePaid := uint64(0)
	if st.MilestonesPaid != nil {
		milestonePaid = st.MilestonesPaid.Uint64()
	}

	return &EscrowCauseView{
		ContractAddress: s.address.Hex(),
		GoalWei:         bigintStr(st.Goal),
		CollectedWei:    bigintStr(st.Collected),
		ReleasedWei:     bigintStr(st.Released),
		MilestonesPaid:  milestonePaid,
		Beneficiary:     st.Beneficiary.Hex(),
		Exists:          st.Exists,
		MilestonePct:    25,
	}, nil
}

func bigintStr(v *big.Int) string {
	if v == nil {
		return "0"
	}
	return v.String()
}
