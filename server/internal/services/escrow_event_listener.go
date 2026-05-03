package services

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/uuid"

	"server/internal/blockchain"
	"server/internal/blockchain/contracts"
	"server/internal/models"
	"server/internal/repository"
)

type EscrowEventListener struct {
	client           *blockchain.Client
	contract         *contracts.MilestoneTracker
	contractAddress  common.Address
	disbursementRepo repository.DisbursementRepository
	organizationRepo repository.OrganizationRepository
	causeRepo        repository.CauseRepository
}

func NewEscrowEventListener(
	client *blockchain.Client,
	contractAddress string,
	disbursementRepo repository.DisbursementRepository,
	organizationRepo repository.OrganizationRepository,
	causeRepo repository.CauseRepository,
) (*EscrowEventListener, error) {
	addr := common.HexToAddress(contractAddress)
	instance, err := contracts.NewMilestoneTracker(addr, client.EthClient)
	if err != nil {
		return nil, err
	}

	return &EscrowEventListener{
		client:           client,
		contract:         instance,
		contractAddress:  addr,
		disbursementRepo: disbursementRepo,
		organizationRepo: organizationRepo,
		causeRepo:        causeRepo,
	}, nil
}

// Start begins listening for MilestoneReached events
func (l *EscrowEventListener) Start(ctx context.Context) error {
	log.Println("Starting milestone tracker event listener...")

	// First, check for any past MilestoneReached events that we might have missed
	// This handles the case where events were emitted while the server was down
	log.Println("Checking for historical MilestoneReached events...")
	if err := l.processPastEvents(ctx); err != nil {
		log.Printf("Warning: Failed to process past events: %v", err)
		// Continue anyway - we'll at least catch new events
	}

	// Create a filter for events from the tracker contract
	// We'll receive all events and filter by parsing
	query := ethereum.FilterQuery{
		Addresses: []common.Address{l.contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := l.client.EthClient.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return fmt.Errorf("failed to subscribe to logs: %w", err)
	}

	log.Println("Subscribed to MilestoneTracker events")

	// Listen for events
	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Printf("Event subscription error: %v", err)
				return
			case vLog := <-logs:
				// Try to parse as MilestoneReached event
				event, err := l.contract.ParseMilestoneReached(vLog)
				if err != nil {
					// Not a MilestoneReached event, ignore it
					// This is normal - could be CauseRegistered or DonationRecorded
					log.Printf("Skipping non-MilestoneReached event (topic: %s)", vLog.Topics[0].Hex())
					continue
				}

				// Process the milestone
				if err := l.processMilestone(ctx, event); err != nil {
					log.Printf("Failed to process milestone: %v", err)
					continue
				}

				log.Printf("[SUCCESS] Successfully processed milestone %d for cause %v", event.Milestone, event.CauseId)
			case <-ctx.Done():
				log.Println("Stopping event listener")
				return
			}
		}
	}()

	return nil
}

// processPastEvents queries and processes historical MilestoneReached events
// This catches events that were emitted while the server was down
func (l *EscrowEventListener) processPastEvents(ctx context.Context) error {
	// Query all MilestoneReached events from contract deployment to now
	opts := &bind.FilterOpts{
		Start:   0,
		End:     nil,
		Context: ctx,
	}

	iterator, err := l.contract.MilestoneTrackerFilterer.FilterMilestoneReached(opts, nil)
	if err != nil {
		return fmt.Errorf("failed to create event filter: %w", err)
	}
	defer iterator.Close()

	processedCount := 0
	skippedCount := 0

	for iterator.Next() {
		event := iterator.Event

		// Check if this disbursement already exists (to avoid duplicates)
		causeID := blockchain.Bytes16ToUUID(event.CauseId)
		existing, err := l.disbursementRepo.GetByCauseAndMilestone(ctx, causeID, int(event.Milestone))
		if err == nil && existing != nil {
			skippedCount++
			continue
		}

		// Process the milestone
		if err := l.processMilestone(ctx, event); err != nil {
			log.Printf("[ERROR] Failed to process historical milestone: %v", err)
			continue
		}

		processedCount++
	}

	if err := iterator.Error(); err != nil {
		return fmt.Errorf("error iterating events: %w", err)
	}

	log.Printf("Historical event processing complete: %d processed, %d skipped (already exists)", processedCount, skippedCount)
	return nil
}

// processMilestone creates a disbursement record and updates organization amount
// This is triggered by MilestoneReached events from the tracker contract
func (l *EscrowEventListener) processMilestone(ctx context.Context, event *contracts.MilestoneTrackerMilestoneReached) error {
	causeID := blockchain.Bytes16ToUUID(event.CauseId)
	log.Printf("[MILESTONE] Processing event: cause=%v, milestone=%d, amount=%s", causeID, event.Milestone, event.AmountToDiburse.String())

	// Get the cause to find the organization
	cause, err := l.causeRepo.GetByID(ctx, causeID)
	if err != nil {
		log.Printf("[ERROR] Failed to get cause %v: %v", causeID, err)
		return fmt.Errorf("failed to get cause: %w", err)
	}
	log.Printf("[MILESTONE] Found cause: %s (org: %v)", cause.Title, cause.Organization.ID)

	// Check if this disbursement already exists
	existing, err := l.disbursementRepo.GetByCauseAndMilestone(ctx, causeID, int(event.Milestone))
	if err == nil && existing != nil {
		log.Printf("[MILESTONE] Disbursement already exists for cause %v milestone %d, skipping", causeID, event.Milestone)
		return nil
	}
	log.Printf("[MILESTONE] No existing disbursement found, creating new one...")

	// Convert amount to float
	// event.AmountToDiburse contains the milestone disbursement amount (25% of goal)
	amountFloat := weiToFloat(event.AmountToDiburse)

	// Create the disbursement record
	// NOTE: This is a VIRTUAL disbursement - actual fund transfer happens off-chain
	disbursement := &models.Disbursement{
		ID:              uuid.New(),
		OrganizationID:  cause.Organization.ID,
		CauseID:         causeID,
		MilestoneNumber: int(event.Milestone),
		Amount:          amountFloat,
		TransactionHash: ptrString(event.Raw.TxHash.Hex()), // Tracker contract tx hash
		DisbursedAt:     time.Now(),
		CreatedAt:       time.Now(),
	}

	log.Printf("[MILESTONE] Creating disbursement: id=%v, org=%v, amount=%.2f", disbursement.ID, disbursement.OrganizationID, amountFloat)
	if err := l.disbursementRepo.Create(ctx, disbursement); err != nil {
		log.Printf("[ERROR] Failed to create disbursement: %v", err)
		return fmt.Errorf("failed to create disbursement: %w", err)
	}
	log.Printf("[MILESTONE] Disbursement created successfully")

	// Update organization's total approved disbursement amount
	// This amount should be paid off-chain via traditional banking/Razorpay
	log.Printf("[MILESTONE] Updating organization %v amount by %.2f", cause.Organization.ID, amountFloat)
	if err := l.organizationRepo.AddToAmount(ctx, cause.Organization.ID, amountFloat); err != nil {
		log.Printf("[ERROR] Failed to update organization amount: %v", err)
		return fmt.Errorf("failed to update organization amount: %w", err)
	}
	log.Printf("[MILESTONE] Organization amount updated successfully")

	log.Printf("[SUCCESS] Created VIRTUAL disbursement %v: %.2f for organization %v (to be paid off-chain)",
		disbursement.ID, amountFloat, cause.Organization.ID)

	return nil
}

// weiToFloat converts wei (big.Int) to float64
func weiToFloat(wei *big.Int) float64 {
	if wei == nil {
		return 0.0
	}
	f := new(big.Float).SetInt(wei)
	result, _ := f.Float64()
	return result
}

func ptrString(s string) *string {
	return &s
}
