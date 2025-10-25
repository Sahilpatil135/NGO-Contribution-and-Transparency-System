package services

import (
	"context"
	"github.com/google/uuid"
	"server/internal/models"
	"server/internal/repository"
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
	cause := &models.Cause{
		ID:              uuid.New(),
		OrganizationID:  req.OrganizationID,
		Title:           req.Title,
		Description:     req.Description,
		DomainID:        req.DomainID,
		AidTypeID:       req.AidTypeID,
		CollectedAmount: req.CollectedAmount,
		GoalAmount:      req.GoalAmount,
		Deadline:        req.Deadline,
		CreatedAt:       req.CreatedAt,
		IsActive:        req.IsActive,
		CoverImageURL:   req.CoverImageURL,
	}

	err := c.causeRepo.Create(ctx, cause)

	if err != nil {
		return nil, err
	}

	return cause, nil
}

func (c *causeService) GetByID(ctx context.Context, id uuid.UUID) (*models.Cause, error) {
	cause, err := c.causeRepo.GetByID(ctx, id)

	if err != nil {
		return nil, err
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
	causesResult, err := c.causeRepo.GetByOrganizationID(ctx, aidTypeID)

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
