package services

import (
	"context"
	"server/internal/models"
	"server/internal/repository"

	"github.com/google/uuid"
)

type CauseService interface {
	Create(ctx context.Context, req *models.CreateCauseRequest) (*models.Cause, error)

	GetByID(ctx context.Context, id uuid.UUID) (*models.Cause, error)
	GetByOrganizationID(ctx context.Context, id uuid.UUID) ([]*models.Cause, error)
	GetByDomainID(ctx context.Context, id uuid.UUID) ([]*models.Cause, error)
	GetByAidTypeID(ctx context.Context, id uuid.UUID) ([]*models.Cause, error)
	GetAll(ctx context.Context) ([]*models.Cause, error)

	// Update(ctx context.Context, cause *models.Cause) error
	Delete(ctx context.Context, id uuid.UUID) error

	GetDomains(ctx context.Context) ([]*models.CauseCategory, error)
	GetAidTypes(ctx context.Context) ([]*models.CauseCategory, error)
	GetDomainByID(ctx context.Context, id uuid.UUID) (*models.CauseCategory, error)
	GetAidTypeByID(ctx context.Context, id uuid.UUID) (*models.CauseCategory, error)
}

type causeService struct {
	causeRepo repository.CauseRepository
}

func NewCauseService(causeRepo repository.CauseRepository) *causeService {
	return &causeService{
		causeRepo: causeRepo,
	}
}

func (c *causeService) Create(ctx context.Context, req *models.CreateCauseRequest) (*models.Cause, error) {
	domain, err := c.GetDomainByID(ctx, req.DomainID)

	if err != nil {
		return nil, err
	}

	aidType, err := c.GetAidTypeByID(ctx, req.AidTypeID)

	if err != nil {
		return nil, err
	}

	organization := &models.OrganizationInCause{
		ID: ctx.Value("organizationID").(uuid.UUID),
	}

	now := req.CreatedAt

	cause := &models.Cause{
		ID:                    uuid.New(),
		Organization:          *organization,
		Title:                 req.Title,
		Description:           req.Description,
		Domain:                *domain,
		AidType:               *aidType,
		CollectedAmount:       req.CollectedAmount,
		GoalAmount:            req.GoalAmount,
		Deadline:              req.Deadline,
		CreatedAt:             now,
		IsActive:              req.IsActive,
		CoverImageURL:         req.CoverImageURL,
		ExecutionLat:          req.ExecutionLat,
		ExecutionLng:          req.ExecutionLng,
		ExecutionRadiusMeters: req.ExecutionRadiusMeters,
		ExecutionStartTime:    req.ExecutionStartTime,
		ExecutionEndTime:      req.ExecutionEndTime,
		FundingStatus:         req.FundingStatus,

		BeneficiariesCount: valueOrDefaultInt(req.BeneficiariesCount, 0),
		ExecutionLocation:  req.ExecutionLocation,
		ImpactGoal:         req.ImpactGoal,
		ProblemStatement:   req.ProblemStatement,
		ExecutionPlan:      req.ExecutionPlan,
		DonorCount:         0,
		UpdatedAt:          now,
	}

	err = c.causeRepo.Create(ctx, cause)

	if err != nil {
		return nil, err
	}

	// Optionally create structured products if provided
	if len(req.Products) > 0 {
		for _, p := range req.Products {
			if p == nil {
				continue
			}
			product := &models.CauseProduct{
				ID:             uuid.New(),
				CauseID:        cause.ID,
				Name:           p.Name,
				Description:    p.Description,
				PricePerUnit:   p.PricePerUnit,
				QuantityNeeded: p.QuantityNeeded,
				QuantityFunded: 0,
				ImageURL:       valueOrDefaultString(p.ImageURL, ""),
				CreatedAt:      now,
			}
			if err := c.causeRepo.CreateProduct(ctx, product); err != nil {
				return nil, err
			}
			cause.Products = append(cause.Products, product)
		}
	}

	return cause, nil
}

func (c *causeService) GetByID(ctx context.Context, id uuid.UUID) (*models.Cause, error) {
	cause, err := c.causeRepo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	// Attach products and updates for detailed campaign page
	if products, err := c.causeRepo.GetProductsByCauseID(ctx, id); err == nil {
		cause.Products = products
	}
	if updates, err := c.causeRepo.GetUpdatesByCauseID(ctx, id); err == nil {
		cause.Updates = updates
	}

	return cause, nil
}

func (c *causeService) GetByOrganizationID(ctx context.Context, organizationID uuid.UUID) ([]*models.Cause, error) {
	causesResult, err := c.causeRepo.GetByOrganizationID(ctx, organizationID)

	if err != nil {
		return nil, err
	}

	return causesResult, nil
}

func (c *causeService) GetByDomainID(ctx context.Context, domainID uuid.UUID) ([]*models.Cause, error) {
	causesResult, err := c.causeRepo.GetByDomainID(ctx, domainID)

	if err != nil {
		return nil, err
	}

	return causesResult, nil
}

func (c *causeService) GetByAidTypeID(ctx context.Context, aidTypeID uuid.UUID) ([]*models.Cause, error) {
	causesResult, err := c.causeRepo.GetByAidTypeID(ctx, aidTypeID)

	if err != nil {
		return nil, err
	}

	return causesResult, nil
}

func (c *causeService) GetAll(ctx context.Context) ([]*models.Cause, error) {
	causesResult, err := c.causeRepo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return causesResult, nil
}

func (c *causeService) Delete(ctx context.Context, id uuid.UUID) error {
	err := c.causeRepo.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}

func (c *causeService) GetDomains(ctx context.Context) ([]*models.CauseCategory, error) {
	domainResults, err := c.causeRepo.GetDomains(ctx)

	if err != nil {
		return nil, err
	}

	return domainResults, err
}

func (c *causeService) GetAidTypes(ctx context.Context) ([]*models.CauseCategory, error) {
	aidTypeResults, err := c.causeRepo.GetAidTypes(ctx)

	if err != nil {
		return nil, err
	}

	return aidTypeResults, err
}

func (c *causeService) GetDomainByID(ctx context.Context, id uuid.UUID) (*models.CauseCategory, error) {
	domain, err := c.causeRepo.GetDomainByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return domain, err
}

func (c *causeService) GetAidTypeByID(ctx context.Context, id uuid.UUID) (*models.CauseCategory, error) {
	aidType, err := c.causeRepo.GetAidTypeByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return aidType, err
}

func valueOrDefaultInt(v *int, def int) int {
	if v == nil {
		return def
	}
	return *v
}

func valueOrDefaultString(v *string, def string) string {
	if v == nil {
		return def
	}
	return *v
}

