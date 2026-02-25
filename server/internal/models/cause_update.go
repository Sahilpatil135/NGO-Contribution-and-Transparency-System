package models

import (
	"time"

	"github.com/google/uuid"
)

// CauseUpdate represents a structured update for a cause (for the Updates tab)
type CauseUpdate struct {
	ID                uuid.UUID `json:"id" db:"id"`
	CauseID           uuid.UUID `json:"cause_id" db:"cause_id"`
	Title             string    `json:"title" db:"title"`
	Description       string    `json:"description" db:"description"`
	UpdateType        string    `json:"update_type" db:"update_type"`
	FundingPercentage *int      `json:"funding_percentage,omitempty" db:"funding_percentage"`
	IsVerified        bool      `json:"is_verified" db:"is_verified"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
}

