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
	FundingStatus         *string             `json:"funding_status" db:"funding_status"`

	// Extended project & execution metadata
	BeneficiariesCount int       `json:"beneficiaries_count" db:"beneficiaries_count"`
	ExecutionLocation  *string   `json:"execution_location" db:"execution_location"`
	ImpactGoal         *string   `json:"impact_goal" db:"impact_goal"`
	ProblemStatement   *string   `json:"problem_statement" db:"problem_statement"`
	ExecutionPlan      *string   `json:"execution_plan" db:"execution_plan"`
	DonorCount         int       `json:"donor_count" db:"donor_count"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`

	// Optional related aggregates for campaign page
	Products []*CauseProduct `json:"products,omitempty"`
	Updates  []*CauseUpdate  `json:"updates,omitempty"`
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
	FundingStatus         *string    `json:"funding_status,omitempty"`

	// Project details (mandatory in UI, optional in API for backward compatibility)
	BeneficiariesCount *int    `json:"beneficiaries_count,omitempty"`
	ExecutionLocation  *string `json:"execution_location,omitempty"`
	ImpactGoal         *string `json:"impact_goal,omitempty"`
	ProblemStatement   *string `json:"problem_statement,omitempty"`
	ExecutionPlan      *string `json:"execution_plan,omitempty"`

	// Optional initial products for structured campaigns
	Products []*CreateCauseProductInput `json:"products,omitempty"`
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

	BeneficiariesCount int       `json:"beneficiaries_count"`
	ExecutionLocation  *string   `json:"execution_location"`
	ImpactGoal         *string   `json:"impact_goal"`
	ProblemStatement   *string   `json:"problem_statement"`
	ExecutionPlan      *string   `json:"execution_plan"`
	DonorCount         int       `json:"donor_count"`
	UpdatedAt          time.Time `json:"updated_at"`

	Products []*CauseProduct `json:"products,omitempty"`
	Updates  []*CauseUpdate  `json:"updates,omitempty"`
}

// ToCauseResponse converts a Cause to CauseResponse
func (c *Cause) ToCauseResponse() CauseResponse {
	// Derive funding status for donor UI based on live data
	computedStatus := computeFundingStatus(c)

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
		FundingStatus:         &computedStatus,

		BeneficiariesCount: c.BeneficiariesCount,
		ExecutionLocation:  c.ExecutionLocation,
		ImpactGoal:         c.ImpactGoal,
		ProblemStatement:   c.ProblemStatement,
		ExecutionPlan:      c.ExecutionPlan,
		DonorCount:         c.DonorCount,
		UpdatedAt:          c.UpdatedAt,

		Products: c.Products,
		Updates:  c.Updates,
	}
}

// computeFundingStatus applies donor-facing funding state rules:
// - collected_amount == 0 → "Not Started"
// - collected_amount > 0 and < goal_amount → "Active"
// - collected_amount >= goal_amount → "Fully Funded"
// - deadline passed and not fully funded → "Closed"
func computeFundingStatus(c *Cause) string {
	// Defensive defaults
	if c.GoalAmount == nil || *c.GoalAmount <= 0 {
		if c.CollectedAmount <= 0 {
			return "Not Started"
		}
		return "Active"
	}

	now := time.Now()
	hasDeadline := c.Deadline != nil
	collected := float64(c.CollectedAmount)
	goal := float64(*c.GoalAmount)

	if collected <= 0 {
		if hasDeadline && c.Deadline.Before(now) {
			return "Closed"
		}
		return "Not Started"
	}

	if collected >= goal {
		return "Fully Funded"
	}

	if hasDeadline && c.Deadline.Before(now) {
		return "Closed"
	}

	return "Active"
}

// CreateCauseProductInput represents a product payload when creating a cause
type CreateCauseProductInput struct {
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	PricePerUnit   float64 `json:"price_per_unit"`
	QuantityNeeded int     `json:"quantity_needed"`
	ImageURL       *string `json:"image_url,omitempty"`
}
