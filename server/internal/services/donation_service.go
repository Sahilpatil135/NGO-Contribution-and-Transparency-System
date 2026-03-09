package services

import (
	"context"
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

	GetFromChainByID(ctx context.Context, id uuid.UUID) (*contracts.DonationLedgerDonation, error)
	GetFromChainByCauseID(ctx context.Context, id uuid.UUID) ([]*contracts.DonationLedgerDonation, error)

	// Update(ctx context.Context, donation *models.Donation) error
	// Delete(ctx context.Context, id uuid.UUID) error
}

type donationService struct {
	donationRepo repository.DonationRepository
	chainService blockchain.DonationChainService
}

func NewDonationService(
	donationRepo repository.DonationRepository,
	chainService blockchain.DonationChainService,
) *donationService {
	return &donationService{
		donationRepo: donationRepo,
		chainService: chainService,
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
		Status:         models.DonationStatusPending,
		PanNumber:      req.PanNumber,
		PaymentID:      req.PaymentID,
		CreatedAt:      time.Now(),
	}

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
