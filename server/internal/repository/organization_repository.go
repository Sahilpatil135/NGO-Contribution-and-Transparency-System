package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"server/internal/models"

	"github.com/google/uuid"
)

type OrganizationRepository interface {
	Create(ctx context.Context, organization *models.Organization) error
	GetByEmail(ctx context.Context, email string) (*models.Organization, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Organization, error)
	GetByProviderID(ctx context.Context, provider, providerID string) (*models.Organization, error)
	// GetAll(ctx context.Context, provider, providerID string) (*models.Organization, error)
	// Update(ctx context.Context, organization *models.Organization) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type organizationRepository struct {
	db *sql.DB
}

func NewOrganizationRepository(db *sql.DB) OrganizationRepository {
	return &organizationRepository{db: db}
}

func (r *organizationRepository) Create(ctx context.Context, organization *models.Organization) error {
	userQuery := `
		INSERT INTO users (
			id, name, email, password_hash, provider, provider_id, avatar_url, is_active, is_verified, role
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err := r.db.ExecContext(ctx, userQuery,
		organization.User.ID,
		organization.User.Name,
		organization.User.Email,
		organization.User.PasswordHash,
		organization.User.Provider,
		organization.User.ProviderID,
		organization.User.AvatarURL,
		organization.User.IsActive,
		organization.User.IsVerified,
		string(models.RoleTypeOrganization),
	)

	organizationQuery := `
		INSERT INTO organizations (
			id, user_id, organization_name, registration_number, organization_type, about, website_url, is_approved, address
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err = r.db.ExecContext(ctx, organizationQuery,
		organization.ID,
		organization.User.ID,
		organization.OrganizationName,
		organization.OrganizationType,
		organization.RegistrationNumber,
		organization.About,
		organization.WebsiteUrl,
		organization.IsApproved,
		organization.Address,
	)

	return err
}

func (r *organizationRepository) GetByEmail(ctx context.Context, email string) (*models.Organization, error) {
	query := `
		SELECT 
		u.id as user_id, u.name, u.email, u.password_hash, u.provider, u.provider_id, u.avatar_url, u.is_active, u.is_verified, u.created_at, u.updated_at, u.role,
		o.id as id, o.organization_name, o.organization_type, o.registration_number, o.about, o.website_url, o.address, o.is_approved
		FROM users u
		FULL JOIN organizations o ON o.user_id = u.id
		WHERE email = $1 AND is_active = true AND u.role = 'organization'
	`

	organization := &models.Organization{
		User: &models.User{},
	}
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&organization.User.ID,
		&organization.User.Name,
		&organization.User.Email,
		&organization.User.PasswordHash,
		&organization.User.Provider,
		&organization.User.ProviderID,
		&organization.User.AvatarURL,
		&organization.User.IsActive,
		&organization.User.IsVerified,
		&organization.User.CreatedAt,
		&organization.User.UpdatedAt,
		string(models.RoleTypeOrganization),

		&organization.ID,
		&organization.OrganizationName,
		&organization.OrganizationType,
		&organization.RegistrationNumber,
		&organization.About,
		&organization.WebsiteUrl,
		&organization.Address,
		&organization.IsApproved,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("organization not found")
		}
		return nil, err
	}

	return organization, nil
}

func (r *organizationRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Organization, error) {
	query := `
		SELECT 
		u.id as user_id, u.name, u.email, u.password_hash, u.provider, u.provider_id, u.avatar_url, u.is_active, u.is_verified, u.created_at, u.updated_at, u.role,
		o.id as id, o.organization_name, o.organization_type, o.registration_number, o.about, o.website_url, o.address, o.is_approved
		FROM users u
		FULL JOIN organizations o ON o.user_id = u.id
		WHERE u.id = $1 AND u.role = 'organization'
	`

	organization := &models.Organization{
		User: &models.User{},
	}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&organization.User.ID,
		&organization.User.Name,
		&organization.User.Email,
		&organization.User.PasswordHash,
		&organization.User.Provider,
		&organization.User.ProviderID,
		&organization.User.AvatarURL,
		&organization.User.IsActive,
		&organization.User.IsVerified,
		&organization.User.CreatedAt,
		&organization.User.UpdatedAt,
		&organization.User.Role,

		&organization.ID,
		&organization.OrganizationName,
		&organization.OrganizationType,
		&organization.RegistrationNumber,
		&organization.About,
		&organization.WebsiteUrl,
		&organization.Address,
		&organization.IsApproved,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("organization not found")
		}
		return nil, err
	}

	return organization, nil
}

func (r *organizationRepository) GetByProviderID(ctx context.Context, provider, providerID string) (*models.Organization, error) {
	query := `
		SELECT 
		u.id as user_id, u.name, u.email, u.password_hash, u.provider, u.provider_id, u.avatar_url, u.is_active, u.is_verified, u.created_at, u.updated_at, u.role,
		o.id as id, o.organization_name, o.organization_type, o.registration_number, o.about, o.website_url, o.address, o.is_approved
		FROM users u
		FULL JOIN organizations o ON o.user_id = u.id
		WHERE u.provider_id = $1 AND u.role = 'organization'
	`

	organization := &models.Organization{}
	err := r.db.QueryRowContext(ctx, query, provider, providerID).Scan(
		&organization.User.ID,
		&organization.User.Name,
		&organization.User.Email,
		&organization.User.PasswordHash,
		&organization.User.Provider,
		&organization.User.ProviderID,
		&organization.User.AvatarURL,
		&organization.User.IsActive,
		&organization.User.IsVerified,
		&organization.User.CreatedAt,
		&organization.User.UpdatedAt,
		&organization.User.Role,

		&organization.ID,
		&organization.OrganizationName,
		&organization.OrganizationType,
		&organization.RegistrationNumber,
		&organization.About,
		&organization.WebsiteUrl,
		&organization.Address,
		&organization.IsApproved,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("organization not found")
		}
		return nil, err
	}

	return organization, nil
}

// func (r *organizationRepository) Update(ctx context.Context, organization *models.Organization) error { }

func (r *organizationRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE users SET is_active = false, updated_at = $2 WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id, time.Now())

	query = `DELETE FROM organizations WHERE users_id = $1`

	_, err = r.db.ExecContext(ctx, query, id, time.Now())
	return err
}
