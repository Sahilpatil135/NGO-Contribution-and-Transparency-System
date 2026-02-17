package repository

import (
	"context"
	"database/sql"

	"server/internal/models"

	"github.com/google/uuid"
)

type ProofSessionRepository interface {
	Create(ctx context.Context, session *models.ProofSession) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.ProofSession, error)
}

type ProofImageRepository interface {
	Create(ctx context.Context, img *models.ProofImage) error
	ExistsBySessionIDAndHash(ctx context.Context, sessionID uuid.UUID, imageHash string) (bool, error)
}

type proofSessionRepository struct {
	db *sql.DB
}

type proofImageRepository struct {
	db *sql.DB
}

func NewProofSessionRepository(db *sql.DB) ProofSessionRepository {
	return &proofSessionRepository{db: db}
}

func NewProofImageRepository(db *sql.DB) ProofImageRepository {
	return &proofImageRepository{db: db}
}

func (r *proofSessionRepository) Create(ctx context.Context, session *models.ProofSession) error {
	query := `
		INSERT INTO proof_sessions (id, organization_id, cause_id, is_active, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.ExecContext(ctx, query,
		session.ID,
		session.OrganizationID,
		session.CauseID,
		session.IsActive,
		session.CreatedAt,
	)
	return err
}

func (r *proofSessionRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.ProofSession, error) {
	query := `
		SELECT id, organization_id, cause_id, is_active, created_at
		FROM proof_sessions
		WHERE id = $1 AND is_active = true
	`
	session := &models.ProofSession{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&session.ID,
		&session.OrganizationID,
		&session.CauseID,
		&session.IsActive,
		&session.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return session, nil
}

func (r *proofImageRepository) Create(ctx context.Context, img *models.ProofImage) error {
	query := `
		INSERT INTO proof_images (id, session_id, image_hash, ipfs_cid, latitude, longitude, timestamp, metadata_score, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := r.db.ExecContext(ctx, query,
		img.ID,
		img.SessionID,
		img.ImageHash,
		img.IPFSCID,
		img.Latitude,
		img.Longitude,
		img.Timestamp,
		img.MetadataScore,
		img.CreatedAt,
	)
	return err
}

func (r *proofImageRepository) ExistsBySessionIDAndHash(ctx context.Context, sessionID uuid.UUID, imageHash string) (bool, error) {
	query := `SELECT 1 FROM proof_images WHERE session_id = $1 AND image_hash = $2 LIMIT 1`
	var one int
	err := r.db.QueryRowContext(ctx, query, sessionID, imageHash).Scan(&one)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
