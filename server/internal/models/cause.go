package models

import (
	"time"

	"github.com/google/uuid"
)

// Cause represents a cause in the system
type Cause struct {
	ID              uuid.UUID  `json:"id" db:"id"`
	OrganizationID  uuid.UUID  `json:"organization_id" db:"organization_id"`
	Title           string     `json:"title" db:"title"`
	Description     *string    `json:"description" db:"description"`
	DomainID        uuid.UUID  `json:"domain_id" db:"domain_id"`
	AidTypeID       uuid.UUID  `json:"aid_type_id" db:"aid_type_id"`
	CollectedAmount float32    `json:"collected_amount" db:"collected_amount"`
	GoalAmount      *float32   `json:"goal_amount" db:"goal_amount"`
	Deadline        *time.Time `json:"deadline" db:"deadline"`
	CreatedAt       time.Time  `json:"created_at" db:"created_at"`
	IsActive        bool       `json:"is_active" db:"is_active"`
	CoverImageURL   *string    `json:"cover_image_url" db:"cover_image_url"`
}

// CreateCauseRequest represents the request payload for creating a cause
type CreateCauseRequest struct {
	OrganizationID  *uuid.UUID `json:"organization_id"`
	Title           string     `json:"title" validate:"required"`
	Description     *string    `json:"description,omitempty"`
	DomainID        uuid.UUID  `json:"domain_id" validate:"required"`
	AidTypeID       uuid.UUID  `json:"aid_type_id" validate:"required"`
	CollectedAmount float32    `json:"collected_amount"`
	GoalAmount      *float32   `json:"goal_amount,omitempty"`
	Deadline        *time.Time `json:"deadline,omitempty"`
	CreatedAt       time.Time  `json:"created_at" validate:"required"`
	IsActive        bool       `json:"is_active" validate:"required"`
	CoverImageURL   *string    `json:"cover_image_url,omitempty"`
}

type CauseByIDRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

// CauseResponse represents a cause response without sensitive data
type CauseResponse struct {
	ID              uuid.UUID  `json:"id"`
	OrganizationID  uuid.UUID  `json:"organization_id"`
	Title           string     `json:"title"`
	Description     *string    `json:"description"`
	DomainID        uuid.UUID  `json:"domain_id"`
	AidTypeID       uuid.UUID  `json:"aid_type_id"`
	CollectedAmount float32    `json:"collected_amount"`
	GoalAmount      *float32   `json:"goal_amount"`
	Deadline        *time.Time `json:"deadline"`
	CreatedAt       time.Time  `json:"created_at"`
	IsActive        bool       `json:"is_active"`
	CoverImageURL   *string    `json:"cover_image_url"`
}

// ToCauseResponse converts a Cause to CauseResponse
func (c *Cause) ToCauseResponse() CauseResponse {
	return CauseResponse{
		ID:              c.ID,
		OrganizationID:  c.OrganizationID,
		Title:           c.Title,
		Description:     c.Description,
		DomainID:        c.DomainID,
		AidTypeID:       c.AidTypeID,
		CollectedAmount: c.CollectedAmount,
		GoalAmount:      c.GoalAmount,
		Deadline:        c.Deadline,
		CreatedAt:       c.CreatedAt,
		IsActive:        c.IsActive,
		CoverImageURL:   c.CoverImageURL,
	}
}
