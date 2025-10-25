package models

import (
	"github.com/google/uuid"
)

// Organization represents a organization in the system
type Organization struct {
	ID                 uuid.UUID `json:"id" db:"id"`
	UserID             uuid.UUID `json:"user_id" db:"user_id"`
	User               *User     `json:"user,omitempty" db:"-"`
	OrganizationName   string    `json:"organization_name" db:"organization_name"`
	OrganizationType   *string   `json:"organization_type" db:"organization_type"`
	RegistrationNumber *string   `json:"registration_number" db:"registration_number"`
	About              *string   `json:"about" db:"about"`
	WebsiteUrl         *string   `json:"website_url" db:"website_url"`
	Address            *string   `json:"address" db:"address"`
	IsApproved         bool      `json:"is_approved" db:"is_approved"`
}

// CreateOrganizationRequest represents the request payload for creating a organization
type CreateOrganizationRequest struct {
	Name               string  `json:"name" validate:"required,min=2,max=255"`
	Email              string  `json:"email" validate:"required,email"`
	Password           string  `json:"password" validate:"required,min=6"`
	OrganizationName   string  `json:"organization_name" validate:"required"`
	OrganizationType   *string `json:"organization_type,omitempty"`
	RegistrationNumber *string `json:"registration_number,omitempty"`
	About              *string `json:"about,omitempty"`
	WebsiteUrl         *string `json:"website_url,omitempty"`
	Address            *string `json:"address,omitempty"`
}

// OrganizationResponse represents a user response without sensitive data
type OrganizationResponse struct {
	ID                 uuid.UUID `json:"id" db:"id"`
	UserID             uuid.UUID `json:"user_id" db:"user_id"`
	User               *User     `json:"user"`
	OrganizationName   string    `json:"organization_name" db:"organization_name"`
	OrganizationType   *string   `json:"organization_type" db:"organization_type"`
	RegistrationNumber *string   `json:"registration_number" db:"registration_number"`
	About              *string   `json:"about" db:"about"`
	WebsiteUrl         *string   `json:"website_url" db:"website_url"`
	Address            *string   `json:"address" db:"address"`
	IsApproved         bool      `json:"is_approved" db:"is_approved"`
}

// ToOrganizationResponse converts a Organization to OrganizationResponse
func (o *Organization) ToOrganizationResponse() OrganizationResponse {
	return OrganizationResponse{
		ID:                 o.ID,
		User:               o.User,
		UserID:             o.UserID,
		OrganizationName:   o.OrganizationName,
		OrganizationType:   o.OrganizationType,
		RegistrationNumber: o.RegistrationNumber,
		About:              o.About,
		WebsiteUrl:         o.WebsiteUrl,
		Address:            o.Address,
		IsApproved:         o.IsApproved,
	}
}
