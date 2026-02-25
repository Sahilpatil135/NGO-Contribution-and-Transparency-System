package models

import (
	"time"

	"github.com/google/uuid"
)

// CauseProduct represents a structured item-based need within a cause
type CauseProduct struct {
	ID             uuid.UUID `json:"id" db:"id"`
	CauseID        uuid.UUID `json:"cause_id" db:"cause_id"`
	Name           string    `json:"name" db:"name"`
	Description    string    `json:"description" db:"description"`
	PricePerUnit   float64   `json:"price_per_unit" db:"price_per_unit"`
	QuantityNeeded int       `json:"quantity_needed" db:"quantity_needed"`
	QuantityFunded int       `json:"quantity_funded" db:"quantity_funded"`
	ImageURL       string    `json:"image_url" db:"image_url"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

