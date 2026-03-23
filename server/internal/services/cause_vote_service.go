package services

import (
	"context"

	"server/internal/models"
	"server/internal/repository"

	"github.com/google/uuid"
)

type CauseVoteService interface {
	ToggleVote(ctx context.Context, causeID uuid.UUID, userID uuid.UUID, voteValue int) (*models.CauseVotesResponse, error)
	GetVotes(ctx context.Context, causeID uuid.UUID, userID uuid.UUID) (*models.CauseVotesResponse, error)
}

type causeVoteService struct {
	repo repository.CauseVoteRepository
}

func NewCauseVoteService(repo repository.CauseVoteRepository) *causeVoteService {
	return &causeVoteService{repo: repo}
}

func (s *causeVoteService) ToggleVote(ctx context.Context, causeID uuid.UUID, userID uuid.UUID, voteValue int) (*models.CauseVotesResponse, error) {
	return s.repo.ToggleVote(ctx, causeID, userID, voteValue)
}

func (s *causeVoteService) GetVotes(ctx context.Context, causeID uuid.UUID, userID uuid.UUID) (*models.CauseVotesResponse, error) {
	return s.repo.GetVotes(ctx, causeID, userID)
}

