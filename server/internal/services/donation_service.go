package services

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"server/internal/blockchain"
	"server/internal/blockchain/contracts"
	"server/internal/models"
	"server/internal/repository"
	"time"

	"github.com/google/uuid"
)

type DonationService interface {
	Create(ctx context.Context, req *models.CreateDonationRequest) (*models.Donation, error)

	GetByID(ctx context.Context, id uuid.UUID) (*models.Donation, error)
	GetByCauseID(ctx context.Context, id uuid.UUID) ([]*models.Donation, error)
	GetByPaymentID(ctx context.Context, id uuid.UUID) (*models.Donation, error)
	GetByUserID(ctx context.Context, id uuid.UUID) ([]*models.Donation, error)

	GetFromChainByID(ctx context.Context, id uuid.UUID) (*contracts.DonationLedgerDonation, error)
	GetFromChainByCauseID(ctx context.Context, id uuid.UUID) ([]*contracts.DonationLedgerDonation, error)

	// Update(ctx context.Context, donation *models.Donation) error
	// Delete(ctx context.Context, id uuid.UUID) error
}

type donationService struct {
	donationRepo   repository.DonationRepository
	chainService   blockchain.DonationChainService
	trackerService *blockchain.MilestoneTrackerService
	causeRepo      repository.CauseRepository
}

func NewDonationService(
	donationRepo repository.DonationRepository,
	chainService blockchain.DonationChainService,
	trackerService *blockchain.MilestoneTrackerService,
	causeRepo repository.CauseRepository,
) *donationService {
	return &donationService{
		donationRepo:   donationRepo,
		chainService:   chainService,
		trackerService: trackerService,
		causeRepo:      causeRepo,
	}
}

func (c *donationService) Create(ctx context.Context, req *models.CreateDonationRequest) (*models.Donation, error) {
	donation := &models.Donation{
		ID:             uuid.New(),
		CauseID:        req.CauseID,
		UserID:         req.UserID,
		Name:           *req.Name,
		Phone:          req.Phone,
		BillingAddress: req.BillingAddress,
		Pincode:        req.Pincode,
		Amount:         req.Amount,
		Status:         models.DonationStatusCompleted,
		PanNumber:      req.PanNumber,
		PaymentID:      req.PaymentID,
		CreatedAt:      time.Now(),
	}

	// Record in DonationLedger (for record-keeping)
	txHash, err := c.chainService.RecordDonation(
		ctx,
		donation.ID,
		donation.CauseID,
		donation.UserID,
		big.NewInt(int64(donation.Amount)),
		*donation.PaymentID,
	)
	if err != nil {
		return nil, err
	}

	donation.TxHash = &txHash

	// Get cause for milestone tracking
	cause, err := c.causeRepo.GetByID(ctx, donation.CauseID)
	if err != nil {
		return nil, fmt.Errorf("failed to get cause: %w", err)
	}

	// If tracker service is available, record donation for milestone tracking
	if c.trackerService != nil && cause != nil && cause.GoalAmount != nil {
		// Calculate total collected amount for this cause from database
		// This ensures the contract starts with the correct baseline when registering
		collectedInDB := float32(cause.CollectedAmount) // Already includes all donations
		
		// Convert rupees to integers for contract (database stores with 2 decimals)
		goalAmount := big.NewInt(int64(*cause.GoalAmount))
		collectedAmount := big.NewInt(int64(collectedInDB))

		err := c.trackerService.EnsureCauseRegistered(
			ctx,
			donation.CauseID,
			goalAmount,
			collectedAmount,
		)
		if err != nil {
			log.Printf("Warning: Failed to ensure cause registration: %v", err)
			// Continue anyway - we'll try again on next donation
		}

		// Record the donation amount for milestone calculation
		donationAmount := big.NewInt(int64(donation.Amount))

		_, err = c.trackerService.RecordDonation(ctx, donation.CauseID, donationAmount)
		if err != nil {
			log.Printf("Warning: Failed to record donation in milestone tracker: %v", err)
			// Don't fail the whole donation if milestone tracking fails
		} else {
			log.Printf("Successfully recorded donation in milestone tracker for cause %v", donation.CauseID)
		}
	}

	if err := c.donationRepo.Create(ctx, donation); err != nil {
		return nil, err
	}

	return donation, nil
}

func (c *donationService) GetByID(ctx context.Context, id uuid.UUID) (*models.Donation, error) {
	donation, err := c.donationRepo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return donation, nil
}

func (c *donationService) GetByCauseID(ctx context.Context, id uuid.UUID) ([]*models.Donation, error) {
	donationsResult, err := c.donationRepo.GetByCauseID(ctx, id)

	if err != nil {
		return nil, err
	}

	return donationsResult, nil
}

func (c *donationService) GetByPaymentID(ctx context.Context, id uuid.UUID) (*models.Donation, error) {
	donation, err := c.donationRepo.GetByPaymentID(ctx, id)

	if err != nil {
		return nil, err
	}

	return donation, nil
}

func (c *donationService) GetByUserID(ctx context.Context, id uuid.UUID) ([]*models.Donation, error) {
	donationsResult, err := c.donationRepo.GetByUserID(ctx, id)

	if err != nil {
		return nil, err
	}

	return donationsResult, nil
}

func (c *donationService) GetFromChainByID(ctx context.Context, id uuid.UUID) (*contracts.DonationLedgerDonation, error) {
	return c.chainService.GetDonation(ctx, id)
}

func (c *donationService) GetFromChainByCauseID(ctx context.Context, id uuid.UUID) ([]*contracts.DonationLedgerDonation, error) {
	donations, err := c.chainService.GetDonationsByCause(ctx, id)
	if err != nil {
		return nil, err
	}

	var result []*contracts.DonationLedgerDonation

	for _, donation := range donations {
		donationLedger, err := c.chainService.GetDonation(ctx, donation)
		if err != nil {
			return nil, err
		}
		result = append(result, donationLedger)
	}

	return result, nil
}

// func (c *donationService) Delete(ctx context.Context, id uuid.UUID) error {
// 	err := c.donationRepo.Delete(ctx, id)
//
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }
