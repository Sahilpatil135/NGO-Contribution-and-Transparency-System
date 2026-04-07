package repository

import (
	"context"
	"database/sql"
	"server/internal/models"

	"github.com/google/uuid"
)

type DisbursementRepository interface {
	Create(ctx context.Context, disbursement *models.Disbursement) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Disbursement, error)
	GetByOrganizationID(ctx context.Context, organizationID uuid.UUID, limit, offset int) ([]*models.Disbursement, error)
	GetByCauseID(ctx context.Context, causeID uuid.UUID) ([]*models.Disbursement, error)
	GetByCauseAndMilestone(ctx context.Context, causeID uuid.UUID, milestone int) (*models.Disbursement, error)
	CountByOrganizationID(ctx context.Context, organizationID uuid.UUID) (int, error)
}

type disbursementRepository struct {
	db *sql.DB
}

func NewDisbursementRepository(db *sql.DB) DisbursementRepository {
	return &disbursementRepository{db: db}
}

func (r *disbursementRepository) Create(ctx context.Context, disbursement *models.Disbursement) error {
	query := `
		INSERT INTO disbursements (organization_id, cause_id, milestone_number, amount, transaction_hash, disbursed_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at
	`

	return r.db.QueryRowContext(
		ctx,
		query,
		disbursement.OrganizationID,
		disbursement.CauseID,
		disbursement.MilestoneNumber,
		disbursement.Amount,
		disbursement.TransactionHash,
		disbursement.DisbursedAt,
	).Scan(&disbursement.ID, &disbursement.CreatedAt)
}

func (r *disbursementRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Disbursement, error) {
	query := `
		SELECT id, organization_id, cause_id, milestone_number, amount, transaction_hash, disbursed_at, created_at
		FROM disbursements
		WHERE id = $1
	`

	disbursement := &models.Disbursement{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&disbursement.ID,
		&disbursement.OrganizationID,
		&disbursement.CauseID,
		&disbursement.MilestoneNumber,
		&disbursement.Amount,
		&disbursement.TransactionHash,
		&disbursement.DisbursedAt,
		&disbursement.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return disbursement, err
}

func (r *disbursementRepository) GetByOrganizationID(ctx context.Context, organizationID uuid.UUID, limit, offset int) ([]*models.Disbursement, error) {
	query := `
		SELECT 
			d.id, d.organization_id, d.cause_id, d.milestone_number, d.amount, d.transaction_hash, d.disbursed_at, d.created_at,
			c.id as cause_id, c.title as cause_title
		FROM disbursements d
		JOIN causes c ON d.cause_id = c.id
		WHERE d.organization_id = $1
		ORDER BY d.disbursed_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.QueryContext(ctx, query, organizationID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var disbursements []*models.Disbursement
	for rows.Next() {
		d := &models.Disbursement{
			Cause: &models.Cause{},
		}

		err := rows.Scan(
			&d.ID,
			&d.OrganizationID,
			&d.CauseID,
			&d.MilestoneNumber,
			&d.Amount,
			&d.TransactionHash,
			&d.DisbursedAt,
			&d.CreatedAt,
			&d.Cause.ID,
			&d.Cause.Title,
		)
		if err != nil {
			return nil, err
		}

		disbursements = append(disbursements, d)
	}

	return disbursements, rows.Err()
}

func (r *disbursementRepository) GetByCauseID(ctx context.Context, causeID uuid.UUID) ([]*models.Disbursement, error) {
	query := `
		SELECT id, organization_id, cause_id, milestone_number, amount, transaction_hash, disbursed_at, created_at
		FROM disbursements
		WHERE cause_id = $1
		ORDER BY milestone_number ASC
	`

	rows, err := r.db.QueryContext(ctx, query, causeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var disbursements []*models.Disbursement
	for rows.Next() {
		d := &models.Disbursement{}
		err := rows.Scan(
			&d.ID,
			&d.OrganizationID,
			&d.CauseID,
			&d.MilestoneNumber,
			&d.Amount,
			&d.TransactionHash,
			&d.DisbursedAt,
			&d.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		disbursements = append(disbursements, d)
	}

	return disbursements, rows.Err()
}

func (r *disbursementRepository) GetByCauseAndMilestone(ctx context.Context, causeID uuid.UUID, milestone int) (*models.Disbursement, error) {
	query := `
		SELECT id, organization_id, cause_id, milestone_number, amount, transaction_hash, disbursed_at, created_at
		FROM disbursements
		WHERE cause_id = $1 AND milestone_number = $2
	`

	disbursement := &models.Disbursement{}
	err := r.db.QueryRowContext(ctx, query, causeID, milestone).Scan(
		&disbursement.ID,
		&disbursement.OrganizationID,
		&disbursement.CauseID,
		&disbursement.MilestoneNumber,
		&disbursement.Amount,
		&disbursement.TransactionHash,
		&disbursement.DisbursedAt,
		&disbursement.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return disbursement, err
}

func (r *disbursementRepository) CountByOrganizationID(ctx context.Context, organizationID uuid.UUID) (int, error) {
	query := `SELECT COUNT(*) FROM disbursements WHERE organization_id = $1`

	var count int
	err := r.db.QueryRowContext(ctx, query, organizationID).Scan(&count)
	return count, err
}
