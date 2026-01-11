package services

import (
	"context"
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

	// Update(ctx context.Context, donation *models.Donation) error
	// Delete(ctx context.Context, id uuid.UUID) error
}

type donationService struct {
	donationRepo repository.DonationRepository
}

func NewDonationService(donationRepo repository.DonationRepository) *donationService {
	return &donationService{
		donationRepo: donationRepo,
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

	err := c.donationRepo.Create(ctx, donation)

	if err != nil {
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

// func (c *donationService) Delete(ctx context.Context, id uuid.UUID) error {
// 	err := c.donationRepo.Delete(ctx, id)
//
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }
