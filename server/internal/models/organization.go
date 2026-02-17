package models

import (
	"github.com/google/uuid"
)

type RoleType string

const (
	RoleTypeUser         RoleType = "user"
	RoleTypeOrganization RoleType = "organization"
	RoleTypeAdmin        RoleType = "admin"
)

// Organization represents a organization in the system
type Organization struct {
	ID                 uuid.UUID `json:"id" db:"id"`
	UserID             uuid.UUID `json:"user_id" db:"user_id"`
	User               *User     `json:"user,omitempty" db:"-"`
	OrganizationName   string    `json:"organization_name" db:"organization_name"`
	RegistrationNumber *string   `json:"registration_number" db:"registration_number"`
	OrganizationType   *string   `json:"organization_type" db:"organization_type"`
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
	RegistrationNumber *string `json:"registration_number,omitempty"`
	OrganizationType   *string `json:"organization_type,omitempty"`	
	About              *string `json:"about,omitempty"`
	WebsiteUrl         *string `json:"website_url,omitempty"`
	Address            *string `json:"address,omitempty"`
	ContactRole        string  `json:"contact_role,omitempty"`
	ContactPhone       string  `json:"contact_phone,omitempty"`
}

// OrganizationResponse represents a user response without sensitive data
type OrganizationResponse struct {
	ID                 uuid.UUID `json:"id" db:"id"`
	UserID             uuid.UUID `json:"user_id" db:"user_id"`
	User               *User     `json:"user"`
	OrganizationName   string    `json:"organization_name" db:"organization_name"`
	RegistrationNumber *string   `json:"registration_number" db:"registration_number"`
	OrganizationType   *string   `json:"organization_type" db:"organization_type"`	
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
		RegistrationNumber: o.RegistrationNumber,
		OrganizationType:   o.OrganizationType,		
		About:              o.About,
		WebsiteUrl:         o.WebsiteUrl,
		Address:            o.Address,
		IsApproved:         o.IsApproved,
	}
}
