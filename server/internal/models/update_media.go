package models

import (
	"time"

	"github.com/google/uuid"
)

// UpdateMedia represents a media asset (receipt, image, pdf) attached to a cause update.
type UpdateMedia struct {
	ID        uuid.UUID `json:"id" db:"id"`
	UpdateID  uuid.UUID `json:"update_id" db:"update_id"`
	MediaType string    `json:"media_type" db:"media_type"`
	MediaURL  string    `json:"media_url" db:"media_url"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

