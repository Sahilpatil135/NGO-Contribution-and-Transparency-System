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

	// GetByID(ctx context.Context, id uuid.UUID) (*models.Donation, error)
	// GetByOrganizationID(ctx context.Context, id uuid.UUID) ([]*models.Donation, error)
	// GetByDomainID(ctx context.Context, id uuid.UUID) ([]*models.Donation, error)
	// GetByAidTypeID(ctx context.Context, id uuid.UUID) ([]*models.Donation, error)
	// GetAll(ctx context.Context) ([]*models.Donation, error)

	// Update(ctx context.Context, donation *models.Donation) error
	// Delete(ctx context.Context, id uuid.UUID) error

	// GetDomains(ctx context.Context) ([]*models.DonationCategory, error)
	// GetAidTypes(ctx context.Context) ([]*models.DonationCategory, error)
	// GetDomainByID(ctx context.Context, id uuid.UUID) (*models.DonationCategory, error)
	// GetAidTypeByID(ctx context.Context, id uuid.UUID) (*models.DonationCategory, error)
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

// func (c *donationService) GetByID(ctx context.Context, id uuid.UUID) (*models.Donation, error) {
// 	donation, err := c.donationRepo.GetByID(ctx, id)
//
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return donation, nil
// }
//
// func (c *donationService) GetByOrganizationID(ctx context.Context, organizationID uuid.UUID) ([]*models.Donation, error) {
// 	donationsResult, err := c.donationRepo.GetByOrganizationID(ctx, organizationID)
//
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return donationsResult, nil
// }
//
// func (c *donationService) GetByDomainID(ctx context.Context, domainID uuid.UUID) ([]*models.Donation, error) {
// 	donationsResult, err := c.donationRepo.GetByDomainID(ctx, domainID)
//
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return donationsResult, nil
// }
//
// func (c *donationService) GetByAidTypeID(ctx context.Context, aidTypeID uuid.UUID) ([]*models.Donation, error) {
// 	donationsResult, err := c.donationRepo.GetByAidTypeID(ctx, aidTypeID)
//
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return donationsResult, nil
// }
//
// func (c *donationService) GetAll(ctx context.Context) ([]*models.Donation, error) {
// 	donationsResult, err := c.donationRepo.GetAll(ctx)
//
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return donationsResult, nil
// }
//
// func (c *donationService) Delete(ctx context.Context, id uuid.UUID) error {
// 	err := c.donationRepo.Delete(ctx, id)
//
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }
//
// func (c *donationService) GetDomains(ctx context.Context) ([]*models.DonationCategory, error) {
// 	domainResults, err := c.donationRepo.GetDomains(ctx)
//
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return domainResults, err
// }
//
// func (c *donationService) GetAidTypes(ctx context.Context) ([]*models.DonationCategory, error) {
// 	aidTypeResults, err := c.donationRepo.GetAidTypes(ctx)
//
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return aidTypeResults, err
// }
//
// func (c *donationService) GetDomainByID(ctx context.Context, id uuid.UUID) (*models.DonationCategory, error) {
// 	domain, err := c.donationRepo.GetDomainByID(ctx, id)
//
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return domain, err
// }
//
// func (c *donationService) GetAidTypeByID(ctx context.Context, id uuid.UUID) (*models.DonationCategory, error) {
// 	aidType, err := c.donationRepo.GetAidTypeByID(ctx, id)
//
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return aidType, err
// }
