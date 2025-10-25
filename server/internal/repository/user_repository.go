package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"server/internal/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetByProviderID(ctx context.Context, provider, providerID string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (id, name, email, password_hash, provider, provider_id, avatar_url, is_active, is_verified, role)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err := r.db.ExecContext(ctx, query,
		user.ID,
		user.Name,
		user.Email,
		user.PasswordHash,
		user.Provider,
		user.ProviderID,
		user.AvatarURL,
		user.IsActive,
		user.IsVerified,
		user.Role,
	)

	return err
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT id, name, email, password_hash, provider, provider_id, avatar_url, is_active, is_verified, created_at, updated_at, role
		FROM users 
		WHERE email = $1 AND is_active = true
	`

	user := &models.User{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.Provider,
		&user.ProviderID,
		&user.AvatarURL,
		&user.IsActive,
		&user.IsVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Role,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	query := `
		SELECT id, name, email, password_hash, provider, provider_id, avatar_url, is_active, is_verified, created_at, updated_at, role
		FROM users 
		WHERE id = $1 AND is_active = true
	`

	user := &models.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.Provider,
		&user.ProviderID,
		&user.AvatarURL,
		&user.IsActive,
		&user.IsVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Role,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetByProviderID(ctx context.Context, provider, providerID string) (*models.User, error) {
	query := `
		SELECT id, name, email, password_hash, provider, provider_id, avatar_url, is_active, is_verified, created_at, updated_at, role
		FROM users 
		WHERE provider = $1 AND provider_id = $2 AND is_active = true
	`

	user := &models.User{}
	err := r.db.QueryRowContext(ctx, query, provider, providerID).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.Provider,
		&user.ProviderID,
		&user.AvatarURL,
		&user.IsActive,
		&user.IsVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Role,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	query := `
		UPDATE users 
		SET name = $2, email = $3, password_hash = $4, provider = $5, provider_id = $6, 
		    avatar_url = $7, is_active = $8, is_verified = $9, updated_at = $10, role = $11
		WHERE id = $1
	`

	user.UpdatedAt = time.Now()
	_, err := r.db.ExecContext(ctx, query,
		user.ID,
		user.Name,
		user.Email,
		user.PasswordHash,
		user.Provider,
		user.ProviderID,
		user.AvatarURL,
		user.IsActive,
		user.IsVerified,
		user.UpdatedAt,
		user.Role,
	)

	return err
}

func (r *userRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE users SET is_active = false, updated_at = $2 WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id, time.Now())
	return err
}
