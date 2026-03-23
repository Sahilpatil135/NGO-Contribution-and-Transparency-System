package models

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

// CauseUpdate represents a structured update for a cause (for the Updates tab)
type CauseUpdate struct {
	ID                 uuid.UUID      `json:"id" db:"id"`
	CauseID            uuid.UUID      `json:"cause_id" db:"cause_id"`
	Title              string         `json:"title" db:"title"`
	Description        string         `json:"description" db:"description"`
	UpdateType         string         `json:"update_type" db:"update_type"`
	FundingPercentage  *int           `json:"funding_percentage,omitempty" db:"funding_percentage"`
	ClaimedAmount      *float64       `json:"claimed_amount,omitempty" db:"claimed_amount"`
	VerificationScore  *float64       `json:"verification_score,omitempty" db:"verification_score"`
	VerificationStatus string         `json:"verification_status,omitempty" db:"verification_status"`
	// Backward-compat for existing frontend (CampaignPage.jsx) expecting is_verified.
	// This is derived from VerificationStatus and is not stored in DB.
	IsVerified bool `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	Media      []*UpdateMedia `json:"media,omitempty"`
}

// CreateCauseUpdateRequest is used by the API when NGOs post an update.
type CreateCauseUpdateRequest struct {
	Title             string   `json:"title"`
	Description       string   `json:"description"`
	UpdateType        string   `json:"update_type"`
	FundingPercentage *int     `json:"funding_percentage,omitempty"`
	ReceiptURLs       []string `json:"receipt_urls,omitempty"`
	// new {
	ReceiptJobIDs    []string `json:"receipt_job_ids,omitempty"`
	// }
	ClaimedAmount     *float64 `json:"claimed_amount,omitempty"`
}

func (u *CauseUpdate) DeriveVerificationFields() {
	if strings.EqualFold(strings.TrimSpace(u.VerificationStatus), "verified") {
		u.IsVerified = true
	} else {
		u.IsVerified = false
	}
}


