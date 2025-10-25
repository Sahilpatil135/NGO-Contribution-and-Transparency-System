package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"server/internal/models"

	"github.com/google/uuid"
)

type CauseRepository interface {
	Create(ctx context.Context, cause *models.Cause) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Cause, error)
	GetByOrganizationID(ctx context.Context, id uuid.UUID) ([]*models.Cause, error)
	GetByDomainID(ctx context.Context, id uuid.UUID) ([]*models.Cause, error)
	GetByAidTypeID(ctx context.Context, id uuid.UUID) ([]*models.Cause, error)
	GetAll(ctx context.Context) ([]*models.Cause, error)
	// Update(ctx context.Context, cause *models.Cause) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type causeRepository struct {
	db *sql.DB
}

func NewCauseRepository(db *sql.DB) CauseRepository {
	return &causeRepository{db: db}
}

func (r *causeRepository) Create(ctx context.Context, cause *models.Cause) error {
	query := `
		INSERT INTO causes (
			id, organization_id, title, description, domain_id, aid_type_id,
			collected_amount, goal_amount, deadline, is_active, cover_image_url, created_at 
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`

	_, err := r.db.ExecContext(ctx, query,
		cause.ID,
		cause.OrganizationID,
		cause.Title,
		cause.Description,
		cause.DomainID,
		cause.AidTypeID,
		cause.CollectedAmount,
		cause.GoalAmount,
		cause.Deadline,
		cause.IsActive,
		cause.CoverImageURL,
		cause.CreatedAt,
	)

	return err
}

func (r *causeRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Cause, error) {
	query := `
		SELECT 
			id, organization_id, title, description, domain_id, aid_type_id,
			collected_amount, goal_amount, deadline, is_active, cover_image_url, created_at 
		FROM causes
		WHERE id = $1
	`

	cause := &models.Cause{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&cause.ID,
		&cause.OrganizationID,
		&cause.Title,
		&cause.Description,
		&cause.DomainID,
		&cause.AidTypeID,
		&cause.CollectedAmount,
		&cause.GoalAmount,
		&cause.Deadline,
		&cause.IsActive,
		&cause.CoverImageURL,
		&cause.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("cause not found")
		}
		return nil, err
	}

	return cause, nil
}

func (r *causeRepository) GetByOrganizationID(ctx context.Context, organizationId uuid.UUID) ([]*models.Cause, error) {
	query := `
		SELECT 
			id, organization_id, title, description, domain_id, aid_type_id,
			collected_amount, goal_amount, deadline, is_active, cover_image_url, created_at 
		FROM causes
		WHERE organization_id = $1
	`

	result, err := r.db.QueryContext(ctx, query, organizationId)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	var causesResult []*models.Cause = make([]*models.Cause, 0, 5)

	for result.Next() {
		cause := &models.Cause{}

		err = result.Scan(
			&cause.ID,
			&cause.OrganizationID,
			&cause.Title,
			&cause.Description,
			&cause.DomainID,
			&cause.AidTypeID,
			&cause.CollectedAmount,
			&cause.GoalAmount,
			&cause.Deadline,
			&cause.IsActive,
			&cause.CoverImageURL,
			&cause.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		causesResult = append(causesResult, cause)
	}

	return causesResult, nil
}

func (r *causeRepository) GetByDomainID(ctx context.Context, domainID uuid.UUID) ([]*models.Cause, error) {
	query := `
		SELECT 
			id, organization_id, title, description, domain_id, aid_type_id,
			collected_amount, goal_amount, deadline, is_active, cover_image_url, created_at 
		FROM causes
		WHERE domain_id = $1
	`

	result, err := r.db.QueryContext(ctx, query, domainID)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	var causesResult []*models.Cause = make([]*models.Cause, 0, 5)

	for result.Next() {
		cause := &models.Cause{}

		err = result.Scan(
			&cause.ID,
			&cause.OrganizationID,
			&cause.Title,
			&cause.Description,
			&cause.DomainID,
			&cause.AidTypeID,
			&cause.CollectedAmount,
			&cause.GoalAmount,
			&cause.Deadline,
			&cause.IsActive,
			&cause.CoverImageURL,
			&cause.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		causesResult = append(causesResult, cause)
	}

	return causesResult, nil
}

func (r *causeRepository) GetByAidTypeID(ctx context.Context, aidTypeId uuid.UUID) ([]*models.Cause, error) {
	query := `
		SELECT 
			id, organization_id, title, description, domain_id, aid_type_id,
			collected_amount, goal_amount, deadline, is_active, cover_image_url, created_at 
		FROM causes
		WHERE aid_type_id = $1
	`

	result, err := r.db.QueryContext(ctx, query, aidTypeId)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	var causesResult []*models.Cause = make([]*models.Cause, 0, 5)

	for result.Next() {
		cause := &models.Cause{}

		err = result.Scan(
			&cause.ID,
			&cause.OrganizationID,
			&cause.Title,
			&cause.Description,
			&cause.DomainID,
			&cause.AidTypeID,
			&cause.CollectedAmount,
			&cause.GoalAmount,
			&cause.Deadline,
			&cause.IsActive,
			&cause.CoverImageURL,
			&cause.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		causesResult = append(causesResult, cause)
	}

	return causesResult, nil
}

func (r *causeRepository) GetAll(ctx context.Context) ([]*models.Cause, error) {
	query := `
		SELECT 
			id, organization_id, title, description, domain_id, aid_type_id,
			collected_amount, goal_amount, deadline, is_active, cover_image_url, created_at 
		FROM causes
		WHERE is_active = true
	`

	result, err := r.db.QueryContext(ctx, query)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	var causesResult []*models.Cause = make([]*models.Cause, 0, 5)

	for result.Next() {
		cause := &models.Cause{}

		err = result.Scan(
			&cause.ID,
			&cause.OrganizationID,
			&cause.Title,
			&cause.Description,
			&cause.DomainID,
			&cause.AidTypeID,
			&cause.CollectedAmount,
			&cause.GoalAmount,
			&cause.Deadline,
			&cause.IsActive,
			&cause.CoverImageURL,
			&cause.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		causesResult = append(causesResult, cause)
	}

	return causesResult, nil
}

// func (r *causeRepository) Update(ctx context.Context, cause *models.Cause) error { }

func (r *causeRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE causes SET is_active = false, updated_at = $2 WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id, time.Now())

	return err
}
