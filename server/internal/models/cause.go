package models

import (
	"time"

	"github.com/google/uuid"
)

type OrganizationInCause struct {
	ID   uuid.UUID `json:"id" db:"organization_id"`
	Name *string   `json:"name" db:"organization_name"`
}

// Cause represents a cause in the system
type Cause struct {
	ID                    uuid.UUID           `json:"id" db:"id"`
	Organization          OrganizationInCause `json:"organization" db:"organization"`
	Title                 string              `json:"title" db:"title"`
	Description           *string             `json:"description" db:"description"`
	Domain                CauseCategory       `json:"domain"`
	AidType               CauseCategory       `json:"aid_type"`
	CollectedAmount       float32             `json:"collected_amount" db:"collected_amount"`
	GoalAmount            *float32            `json:"goal_amount" db:"goal_amount"`
	Deadline              *time.Time          `json:"deadline" db:"deadline"`
	CreatedAt             time.Time           `json:"created_at" db:"created_at"`
	IsActive              bool                `json:"is_active" db:"is_active"`
	CoverImageURL         *string             `json:"cover_image_url" db:"cover_image_url"`
	ExecutionLat          *float64            `json:"execution_lat" db:"execution_lat"`
	ExecutionLng          *float64            `json:"execution_lng" db:"execution_lng"`
	ExecutionRadiusMeters *int                `json:"execution_radius_meters" db:"execution_radius_meters"`
	ExecutionStartTime    *time.Time          `json:"execution_start_time" db:"execution_start_time"`
	ExecutionEndTime      *time.Time          `json:"execution_end_time" db:"execution_end_time"`
	FundingStatus        *string             `json:"funding_status" db:"funding_status"`
}

type CauseCategory struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description *string   `json:"description" db:"description"`
	IconURL     *string   `json:"icon_url" db:"icon_url"`
}

// CreateCauseRequest represents the request payload for creating a cause
type CreateCauseRequest struct {
	OrganizationID        *uuid.UUID `json:"organization_id"`
	Title                 string     `json:"title" validate:"required"`
	Description           *string    `json:"description,omitempty"`
	DomainID              uuid.UUID  `json:"domain_id" validate:"required"`
	AidTypeID             uuid.UUID  `json:"aid_type_id" validate:"required"`
	CollectedAmount       float32    `json:"collected_amount"`
	GoalAmount            *float32   `json:"goal_amount,omitempty"`
	Deadline              *time.Time `json:"deadline,omitempty"`
	CreatedAt             time.Time  `json:"created_at" validate:"required"`
	IsActive              bool       `json:"is_active" validate:"required"`
	CoverImageURL         *string    `json:"cover_image_url,omitempty"`
	ExecutionLat          *float64   `json:"execution_lat,omitempty"`
	ExecutionLng          *float64   `json:"execution_lng,omitempty"`
	ExecutionRadiusMeters *int       `json:"execution_radius_meters,omitempty"`
	ExecutionStartTime    *time.Time `json:"execution_start_time,omitempty"`
	ExecutionEndTime      *time.Time `json:"execution_end_time,omitempty"`
	FundingStatus        *string    `json:"funding_status,omitempty"`
}

type CauseByIDRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

// CauseResponse represents a cause response without sensitive data
type CauseResponse struct {
	ID                    uuid.UUID           `json:"id"`
	Organization          OrganizationInCause `json:"organization"`
	Title                 string              `json:"title"`
	Description           *string             `json:"description"`
	Domain                CauseCategory       `json:"domain"`
	AidType               CauseCategory       `json:"aid_type"`
	CollectedAmount       float32             `json:"collected_amount"`
	GoalAmount            *float32            `json:"goal_amount"`
	Deadline              *time.Time          `json:"deadline"`
	CreatedAt             time.Time           `json:"created_at"`
	IsActive              bool                `json:"is_active"`
	CoverImageURL         *string             `json:"cover_image_url"`
	ExecutionLat          *float64            `json:"execution_lat"`
	ExecutionLng          *float64            `json:"execution_lng"`
	ExecutionRadiusMeters *int                `json:"execution_radius_meters"`
	ExecutionStartTime    *time.Time          `json:"execution_start_time"`
	ExecutionEndTime      *time.Time          `json:"execution_end_time"`
	FundingStatus         *string             `json:"funding_status"`
}

// ToCauseResponse converts a Cause to CauseResponse
func (c *Cause) ToCauseResponse() CauseResponse {
	return CauseResponse{
		ID:                    c.ID,
		Organization:          c.Organization,
		Title:                 c.Title,
		Description:           c.Description,
		Domain:                c.Domain,
		AidType:               c.AidType,
		CollectedAmount:       c.CollectedAmount,
		GoalAmount:            c.GoalAmount,
		Deadline:              c.Deadline,
		CreatedAt:             c.CreatedAt,
		IsActive:              c.IsActive,
		CoverImageURL:         c.CoverImageURL,
		ExecutionLat:          c.ExecutionLat,
		ExecutionLng:          c.ExecutionLng,
		ExecutionRadiusMeters: c.ExecutionRadiusMeters,
		ExecutionStartTime:    c.ExecutionStartTime,
		ExecutionEndTime:      c.ExecutionEndTime,
		FundingStatus:         c.FundingStatus,
	}
}
