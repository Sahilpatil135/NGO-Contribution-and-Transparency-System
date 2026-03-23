package services

import (
	"context"

	"server/internal/models"
	"server/internal/repository"

	"github.com/google/uuid"
)

type CauseReviewService interface {
	UserCanReviewCause(ctx context.Context, causeID uuid.UUID, userID uuid.UUID) (bool, error)
	CreateReview(ctx context.Context, causeID uuid.UUID, userID uuid.UUID, reviewText string) (*models.CauseReviewResponse, error)
	GetReviewsByCauseID(ctx context.Context, causeID uuid.UUID) (*models.CauseReviewsResponse, error)
	GetReviewCountByCauseID(ctx context.Context, causeID uuid.UUID) (int, error)
}

type causeReviewService struct {
	repo repository.CauseReviewRepository
}

func NewCauseReviewService(repo repository.CauseReviewRepository) *causeReviewService {
	return &causeReviewService{repo: repo}
}

func (s *causeReviewService) UserCanReviewCause(ctx context.Context, causeID uuid.UUID, userID uuid.UUID) (bool, error) {
	return s.repo.UserCanReviewCause(ctx, causeID, userID)
}

func (s *causeReviewService) CreateReview(ctx context.Context, causeID uuid.UUID, userID uuid.UUID, reviewText string) (*models.CauseReviewResponse, error) {
	return s.repo.CreateReview(ctx, causeID, userID, reviewText)
}

func (s *causeReviewService) GetReviewsByCauseID(ctx context.Context, causeID uuid.UUID) (*models.CauseReviewsResponse, error) {
	return s.repo.GetReviewsByCauseID(ctx, causeID)
}

func (s *causeReviewService) GetReviewCountByCauseID(ctx context.Context, causeID uuid.UUID) (int, error) {
	return s.repo.GetReviewCountByCauseID(ctx, causeID)
}

