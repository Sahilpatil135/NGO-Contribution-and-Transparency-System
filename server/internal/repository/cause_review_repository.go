package repository

import (
	"context"
	"database/sql"
	"fmt"

	"server/internal/models"

	"github.com/google/uuid"
)

type CauseReviewRepository interface {
	UserCanReviewCause(ctx context.Context, causeID uuid.UUID, userID uuid.UUID) (bool, error)
	CreateReview(ctx context.Context, causeID uuid.UUID, userID uuid.UUID, reviewText string) (*models.CauseReviewResponse, error)
	GetReviewsByCauseID(ctx context.Context, causeID uuid.UUID) (*models.CauseReviewsResponse, error)
	GetReviewCountByCauseID(ctx context.Context, causeID uuid.UUID) (int, error)
}

type causeReviewRepository struct {
	db *sql.DB
}

func NewCauseReviewRepository(db *sql.DB) CauseReviewRepository {
	return &causeReviewRepository{db: db}
}

func (r *causeReviewRepository) UserCanReviewCause(
	ctx context.Context,
	causeID uuid.UUID,
	userID uuid.UUID,
) (bool, error) {

	var canReview bool

	err := r.db.QueryRowContext(
		ctx,
		`
		SELECT EXISTS (
			SELECT 1 
			FROM donations d
			WHERE d.cause_id = $1 
			  AND d.user_id = $2
		)
		AND NOT EXISTS (
			SELECT 1 
			FROM cause_reviews cr
			WHERE cr.cause_id = $1 
			  AND cr.user_id = $2
		)
		`,
		causeID,
		userID,
	).Scan(&canReview)

	if err != nil {
		return false, err
	}

	return canReview, nil
}

func (r *causeReviewRepository) CreateReview(ctx context.Context, causeID uuid.UUID, userID uuid.UUID, reviewText string) (*models.CauseReviewResponse, error) {
	var res models.CauseReviewResponse
	err := r.db.QueryRowContext(
		ctx,
		`INSERT INTO cause_reviews (cause_id, user_id, review_text)
		 VALUES ($1, $2, $3)
		 ON CONFLICT (cause_id, user_id)
		 DO UPDATE SET review_text = EXCLUDED.review_text
		 RETURNING id, cause_id, user_id, review_text, created_at`,
		causeID,
		userID,
		reviewText,
	).Scan(
		&res.ID,
		&res.CauseID,
		&res.UserID,
		&res.ReviewText,
		&res.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	// Attach user name for UI.
	err = r.db.QueryRowContext(
		ctx,
		"SELECT name FROM users WHERE id = $1",
		userID,
	).Scan(&res.UserName)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *causeReviewRepository) GetReviewsByCauseID(ctx context.Context, causeID uuid.UUID) (*models.CauseReviewsResponse, error) {
	var count int
	err := r.db.QueryRowContext(
		ctx,
		"SELECT COUNT(*) FROM cause_reviews WHERE cause_id = $1",
		causeID,
	).Scan(&count)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(
		ctx,
		`SELECT
			cr.id,
			cr.cause_id,
			cr.user_id,
			cr.review_text,
			cr.created_at,
			u.name
		FROM cause_reviews cr
		JOIN users u ON u.id = cr.user_id
		WHERE cr.cause_id = $1
		ORDER BY cr.created_at DESC`,
		causeID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return &models.CauseReviewsResponse{Count: 0, Reviews: []*models.CauseReviewResponse{}}, nil
		}
		return nil, err
	}
	defer rows.Close()

	reviews := make([]*models.CauseReviewResponse, 0, 5)
	for rows.Next() {
		rv := &models.CauseReviewResponse{}
		if err := rows.Scan(
			&rv.ID,
			&rv.CauseID,
			&rv.UserID,
			&rv.ReviewText,
			&rv.CreatedAt,
			&rv.UserName,
		); err != nil {
			return nil, err
		}
		reviews = append(reviews, rv)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("failed to iterate reviews: %w", rows.Err())
	}

	return &models.CauseReviewsResponse{
		Count:   count,
		Reviews: reviews,
	}, nil
}

func (r *causeReviewRepository) GetReviewCountByCauseID(ctx context.Context, causeID uuid.UUID) (int, error) {
	var count int
	err := r.db.QueryRowContext(
		ctx,
		"SELECT COUNT(*) FROM cause_reviews WHERE cause_id = $1",
		causeID,
	).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
