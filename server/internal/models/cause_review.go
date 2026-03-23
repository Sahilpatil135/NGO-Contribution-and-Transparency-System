package models

import (
	"time"

	"github.com/google/uuid"
)

type CreateCauseReviewRequest struct {
	ReviewText string `json:"review_text"`
}

type CauseReviewResponse struct {
	ID        uuid.UUID `json:"id"`
	CauseID   uuid.UUID `json:"cause_id"`
	UserID    uuid.UUID `json:"user_id"`
	UserName  string    `json:"user_name"`
	ReviewText string   `json:"review_text"`
	CreatedAt time.Time `json:"created_at"`
}

type CauseReviewsResponse struct {
	Count   int                  `json:"count"`
	Reviews []*CauseReviewResponse `json:"reviews"`
}

