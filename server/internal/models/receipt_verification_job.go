package models

import (
	"time"

	"github.com/google/uuid"
)

// ReceiptVerificationJob represents an async verification task for a single receipt image.
// It is created when a receipt is uploaded and updated when the AI service finishes.
type ReceiptVerificationJob struct {
	ID            uuid.UUID  `json:"id" db:"id"`
	OrganizationID uuid.UUID `json:"organization_id" db:"organization_id"`

	// receipt_path is a local filesystem path used by the AI service.
	ReceiptPath string `json:"-" db:"receipt_path"`

	// ClaimedAmount is required by the AI receipt analysis endpoint.
	ClaimedAmount float64 `json:"claimed_amount" db:"claimed_amount"`

	Status       string   `json:"status" db:"status"` // pending | verified | review | rejected | error
	ReceiptScore *float64 `json:"receipt_score,omitempty" db:"receipt_score"`

	ErrorMessage *string `json:"error_message,omitempty" db:"error_message"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type ReceiptStatusResponse struct {
	ReceiptJobID  uuid.UUID `json:"receipt_job_id"`
	Status         string   `json:"status"`
	ReceiptScore   *float64 `json:"receipt_score,omitempty"`
	ErrorMessage   *string  `json:"error_message,omitempty"`
	// Optional: helpful for debugging UI; do not rely on it for business logic.
	ClaimedAmount  float64  `json:"claimed_amount,omitempty"`
}

