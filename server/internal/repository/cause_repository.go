package repository

import (
	"context"
	"database/sql"
	"fmt"

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

	// GetCauseExecution returns execution window and location for proof validation
	GetCauseExecution(ctx context.Context, causeID uuid.UUID) (*models.CauseExecution, error)

	// Update(ctx context.Context, cause *models.Cause) error
	Delete(ctx context.Context, id uuid.UUID) error

	GetDomains(ctx context.Context) ([]*models.CauseCategory, error)
	GetAidTypes(ctx context.Context) ([]*models.CauseCategory, error)
	GetDomainByID(ctx context.Context, id uuid.UUID) (*models.CauseCategory, error)
	GetAidTypeByID(ctx context.Context, id uuid.UUID) (*models.CauseCategory, error)
}

type causeRepository struct {
	db *sql.DB
}

func NewCauseRepository(db *sql.DB) CauseRepository {
	return &causeRepository{db: db}
}

func (c *causeRepository) Create(ctx context.Context, cause *models.Cause) error {
	query := `
		INSERT INTO causes (
			id, organization_id, title, description, domain_id, aid_type_id,
			collected_amount, goal_amount, deadline, is_active, cover_image_url, created_at,
			execution_lat, execution_lng, execution_radius_meters, execution_start_time, execution_end_time, funding_status
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)
	`

	_, err := c.db.ExecContext(ctx, query,
		cause.ID,
		cause.Organization.ID,
		cause.Title,
		cause.Description,
		cause.Domain.ID,
		cause.AidType.ID,
		cause.CollectedAmount,
		cause.GoalAmount,
		cause.Deadline,
		cause.IsActive,
		cause.CoverImageURL,
		cause.CreatedAt,
		cause.ExecutionLat,
		cause.ExecutionLng,
		cause.ExecutionRadiusMeters,
		cause.ExecutionStartTime,
		cause.ExecutionEndTime,
		cause.FundingStatus,
	)

	return err
}

func GetCauseByColumnID(c *causeRepository, ctx context.Context, id uuid.UUID, column string) (*models.Cause, error) {
	query := fmt.Sprintf(`
		SELECT 
			c.id, c.title, c.description, c.collected_amount,
			c.goal_amount, c.deadline, c.is_active, c.cover_image_url, c.created_at,
			c.execution_lat, c.execution_lng, c.execution_radius_meters, c.execution_start_time, c.execution_end_time, c.funding_status,
			cd.id, cd.name, cd.description, cd.icon_url,
			ca.id, ca.name, ca.description, ca.icon_url,
			o.id, o.organization_name
		FROM causes c 
		LEFT JOIN cause_domains cd on cd.id = c.domain_id
		LEFT JOIN cause_aid_types ca on ca.id = c.aid_type_id
		LEFT JOIN organizations o on o.id = c.organization_id
		WHERE c.%s = $1
		`, column)

	cause := &models.Cause{
		Domain:       models.CauseCategory{},
		AidType:      models.CauseCategory{},
		Organization: models.OrganizationInCause{},
	}

	err := c.db.QueryRowContext(ctx, query, id).Scan(
		&cause.ID,
		&cause.Title,
		&cause.Description,
		&cause.CollectedAmount,
		&cause.GoalAmount,
		&cause.Deadline,
		&cause.IsActive,
		&cause.CoverImageURL,
		&cause.CreatedAt,
		&cause.ExecutionLat,
		&cause.ExecutionLng,
		&cause.ExecutionRadiusMeters,
		&cause.ExecutionStartTime,
		&cause.ExecutionEndTime,
		&cause.FundingStatus,

		&cause.Domain.ID,
		&cause.Domain.Name,
		&cause.Domain.Description,
		&cause.Domain.IconURL,

		&cause.AidType.ID,
		&cause.AidType.Name,
		&cause.AidType.Description,
		&cause.AidType.IconURL,

		&cause.Organization.ID,
		&cause.Organization.Name,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("cause not found")
		}
		return nil, err
	}

	return cause, nil
}

func GetCausesByColumnID(c *causeRepository, ctx context.Context, id uuid.UUID, column string) ([]*models.Cause, error) {
	query := fmt.Sprintf(`
		SELECT 
			c.id, c.title, c.description, c.collected_amount,
			c.goal_amount, c.deadline, c.is_active, c.cover_image_url, c.created_at,
			c.execution_lat, c.execution_lng, c.execution_radius_meters, c.execution_start_time, c.execution_end_time, c.funding_status,
			cd.id, cd.name, cd.description, cd.icon_url,
			ca.id, ca.name, ca.description, ca.icon_url,
			o.id, o.organization_name
		FROM causes c 
		LEFT JOIN cause_domains cd on cd.id = c.domain_id
		LEFT JOIN cause_aid_types ca on ca.id = c.aid_type_id
		LEFT JOIN organizations o on o.id = c.organization_id
		WHERE %s = $1
	`, column)

	result, err := c.db.QueryContext(ctx, query, id)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	var causesResult []*models.Cause = make([]*models.Cause, 0, 5)

	for result.Next() {
		cause := &models.Cause{
			Domain:       models.CauseCategory{},
			AidType:      models.CauseCategory{},
			Organization: models.OrganizationInCause{},
		}

		err = result.Scan(
			&cause.ID,
			&cause.Title,
			&cause.Description,
			&cause.CollectedAmount,
			&cause.GoalAmount,
			&cause.Deadline,
			&cause.IsActive,
			&cause.CoverImageURL,
			&cause.CreatedAt,
			&cause.ExecutionLat,
			&cause.ExecutionLng,
			&cause.ExecutionRadiusMeters,
			&cause.ExecutionStartTime,
			&cause.ExecutionEndTime,
			&cause.FundingStatus,

			&cause.Domain.ID,
			&cause.Domain.Name,
			&cause.Domain.Description,
			&cause.Domain.IconURL,

			&cause.AidType.ID,
			&cause.AidType.Name,
			&cause.AidType.Description,
			&cause.AidType.IconURL,

			&cause.Organization.ID,
			&cause.Organization.Name,
		)

		if err != nil {
			return nil, err
		}

		causesResult = append(causesResult, cause)
	}

	return causesResult, nil
}

func (c *causeRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Cause, error) {
	return GetCauseByColumnID(c, ctx, id, "id")
}

func (c *causeRepository) GetCauseExecution(ctx context.Context, causeID uuid.UUID) (*models.CauseExecution, error) {
	query := `
		SELECT id, execution_lat, execution_lng, execution_radius_meters, execution_start_time, execution_end_time
		FROM causes
		WHERE id = $1
	`
	ex := &models.CauseExecution{}
	var id uuid.UUID
	err := c.db.QueryRowContext(ctx, query, causeID).Scan(
		&id,
		&ex.ExecutionLat,
		&ex.ExecutionLng,
		&ex.ExecutionRadiusMeters,
		&ex.ExecutionStartTime,
		&ex.ExecutionEndTime,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	ex.CauseID = id
	return ex, nil
}

func (c *causeRepository) GetByOrganizationID(ctx context.Context, organizationId uuid.UUID) ([]*models.Cause, error) {
	return GetCausesByColumnID(c, ctx, organizationId, "organization_id")
}

func (c *causeRepository) GetByDomainID(ctx context.Context, domainID uuid.UUID) ([]*models.Cause, error) {
	return GetCausesByColumnID(c, ctx, domainID, "domain_id")
}

func (c *causeRepository) GetByAidTypeID(ctx context.Context, aidTypeId uuid.UUID) ([]*models.Cause, error) {
	return GetCausesByColumnID(c, ctx, aidTypeId, "aid_type_id")
}

func (c *causeRepository) GetAll(ctx context.Context) ([]*models.Cause, error) {
	query := `
		SELECT 
			c.id, c.title, c.description, c.collected_amount,
			c.goal_amount, c.deadline, c.is_active, c.cover_image_url, c.created_at,
			c.execution_lat, c.execution_lng, c.execution_radius_meters, c.execution_start_time, c.execution_end_time, c.funding_status,
			cd.id, cd.name, cd.description, cd.icon_url,
			ca.id, ca.name, ca.description, ca.icon_url,
			o.id, o.organization_name
		FROM causes c 
		LEFT JOIN cause_domains cd on cd.id = c.domain_id
		LEFT JOIN cause_aid_types ca on ca.id = c.aid_type_id
		LEFT JOIN organizations o on o.id = c.organization_id
		WHERE is_active = true
		ORDER BY random()
	`

	limitQuery := `LIMIT $1`
	limit := ctx.Value("limit")

	var err error
	var result *sql.Rows

	if limit == nil {
		result, err = c.db.QueryContext(ctx, query)
	} else {
		result, err = c.db.QueryContext(ctx, query+limitQuery, limit)
	}

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	var causesResult []*models.Cause = make([]*models.Cause, 0, 5)

	for result.Next() {
		cause := &models.Cause{
			Domain:       models.CauseCategory{},
			AidType:      models.CauseCategory{},
			Organization: models.OrganizationInCause{},
		}

		err = result.Scan(
			&cause.ID,
			&cause.Title,
			&cause.Description,
			&cause.CollectedAmount,
			&cause.GoalAmount,
			&cause.Deadline,
			&cause.IsActive,
			&cause.CoverImageURL,
			&cause.CreatedAt,
			&cause.ExecutionLat,
			&cause.ExecutionLng,
			&cause.ExecutionRadiusMeters,
			&cause.ExecutionStartTime,
			&cause.ExecutionEndTime,
			&cause.FundingStatus,

			&cause.Domain.ID,
			&cause.Domain.Name,
			&cause.Domain.Description,
			&cause.Domain.IconURL,

			&cause.AidType.ID,
			&cause.AidType.Name,
			&cause.AidType.Description,
			&cause.AidType.IconURL,

			&cause.Organization.ID,
			&cause.Organization.Name,
		)

		if err != nil {
			return nil, err
		}

		causesResult = append(causesResult, cause)
	}

	return causesResult, nil
}

// func (r *causeRepository) Update(ctx context.Context, cause *models.Cause) error { }

func (c *causeRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE causes SET is_active = false WHERE id = $1`

	_, err := c.db.ExecContext(ctx, query, id)

	return err
}

func (c *causeRepository) GetDomains(ctx context.Context) ([]*models.CauseCategory, error) {
	query := `
		SELECT id, name, description, icon_url
		FROM cause_domains
		ORDER BY name DESC
	`

	result, err := c.db.QueryContext(ctx, query)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	var domainResults []*models.CauseCategory = make([]*models.CauseCategory, 0, 5)

	for result.Next() {
		domain := &models.CauseCategory{}

		err = result.Scan(
			&domain.ID,
			&domain.Name,
			&domain.Description,
			&domain.IconURL,
		)

		if err != nil {
			return nil, err
		}

		domainResults = append(domainResults, domain)
	}

	return domainResults, nil
}

func (c *causeRepository) GetAidTypes(ctx context.Context) ([]*models.CauseCategory, error) {
	query := `
		SELECT id, name, description, icon_url
		FROM cause_aid_types
	`

	result, err := c.db.QueryContext(ctx, query)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	var aidTypeResults []*models.CauseCategory = make([]*models.CauseCategory, 0, 5)

	for result.Next() {
		aidType := &models.CauseCategory{}

		err = result.Scan(
			&aidType.ID,
			&aidType.Name,
			&aidType.Description,
			&aidType.IconURL,
		)

		if err != nil {
			return nil, err
		}

		aidTypeResults = append(aidTypeResults, aidType)
	}

	return aidTypeResults, nil
}

func (c *causeRepository) GetDomainByID(ctx context.Context, id uuid.UUID) (*models.CauseCategory, error) {
	query := `
		SELECT id, name, description, icon_url
		FROM cause_domains
		WHERE id = $1
	`

	domain := &models.CauseCategory{}

	err := c.db.QueryRowContext(ctx, query, id).Scan(
		&domain.ID,
		&domain.Name,
		&domain.Description,
		&domain.IconURL,
	)

	if err != nil {
		return nil, err
	}

	return domain, err
}

func (c *causeRepository) GetAidTypeByID(ctx context.Context, id uuid.UUID) (*models.CauseCategory, error) {
	query := `
		SELECT id, name, description, icon_url
		FROM cause_aid_types
		WHERE id = $1
	`

	aidType := &models.CauseCategory{}

	err := c.db.QueryRowContext(ctx, query, id).Scan(
		&aidType.ID,
		&aidType.Name,
		&aidType.Description,
		&aidType.IconURL,
	)

	if err != nil {
		return nil, err
	}

	return aidType, err
}
