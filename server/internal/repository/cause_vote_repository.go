package repository

import (
	"context"
	"database/sql"
	"fmt"

	"server/internal/models"

	"github.com/google/uuid"
)

type CauseVoteRepository interface {
	ToggleVote(ctx context.Context, causeID uuid.UUID, userID uuid.UUID, voteValue int) (*models.CauseVotesResponse, error)
	GetVotes(ctx context.Context, causeID uuid.UUID, userID uuid.UUID) (*models.CauseVotesResponse, error)
}

type causeVoteRepository struct {
	db *sql.DB
}

func NewCauseVoteRepository(db *sql.DB) CauseVoteRepository {
	return &causeVoteRepository{db: db}
}

func (r *causeVoteRepository) ToggleVote(ctx context.Context, causeID uuid.UUID, userID uuid.UUID, voteValue int) (*models.CauseVotesResponse, error) {
	if voteValue != 1 && voteValue != -1 {
		return nil, fmt.Errorf("invalid vote value")
	}

	var existing sql.NullInt64
	err := r.db.QueryRowContext(
		ctx,
		"SELECT vote_value FROM cause_votes WHERE cause_id = $1 AND user_id = $2",
		causeID,
		userID,
	).Scan(&existing)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		// No existing vote: insert.
		_, err = r.db.ExecContext(
			ctx,
			"INSERT INTO cause_votes (cause_id, user_id, vote_value) VALUES ($1, $2, $3)",
			causeID,
			userID,
			voteValue,
		)
		if err != nil {
			return nil, err
		}
	} else if existing.Valid && int(existing.Int64) == voteValue {
		// Same vote clicked again: toggle off.
		_, err = r.db.ExecContext(
			ctx,
			"DELETE FROM cause_votes WHERE cause_id = $1 AND user_id = $2",
			causeID,
			userID,
		)
		if err != nil {
			return nil, err
		}
	} else {
		// Existing vote is different: update.
		_, err = r.db.ExecContext(
			ctx,
			"UPDATE cause_votes SET vote_value = $3, created_at = NOW() WHERE cause_id = $1 AND user_id = $2",
			causeID,
			userID,
			voteValue,
		)
		if err != nil {
			return nil, err
		}
	}

	return r.GetVotes(ctx, causeID, userID)
}

func (r *causeVoteRepository) GetVotes(ctx context.Context, causeID uuid.UUID, userID uuid.UUID) (*models.CauseVotesResponse, error) {
	var upvotes int
	var downvotes int

	err := r.db.QueryRowContext(
		ctx,
		`SELECT
			COALESCE(SUM(CASE WHEN vote_value = 1 THEN 1 ELSE 0 END), 0) AS upvotes,
			COALESCE(SUM(CASE WHEN vote_value = -1 THEN 1 ELSE 0 END), 0) AS downvotes
		FROM cause_votes
		WHERE cause_id = $1`,
		causeID,
	).Scan(&upvotes, &downvotes)
	if err != nil {
		return nil, err
	}

	var myVote sql.NullInt64
	err = r.db.QueryRowContext(
		ctx,
		"SELECT vote_value FROM cause_votes WHERE cause_id = $1 AND user_id = $2",
		causeID,
		userID,
	).Scan(&myVote)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	var myVoteStr *string
	if err == nil && myVote.Valid {
		v := int(myVote.Int64)
		if v == 1 {
			s := "up"
			myVoteStr = &s
		} else if v == -1 {
			s := "down"
			myVoteStr = &s
		}
	}

	return &models.CauseVotesResponse{
		Upvotes:   upvotes,
		Downvotes: downvotes,
		MyVote:    myVoteStr,
	}, nil
}

