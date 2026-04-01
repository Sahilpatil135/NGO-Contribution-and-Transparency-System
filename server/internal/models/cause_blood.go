package models

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type CauseBlood struct {
	ID                uuid.UUID  `json:"id" db:"id"`
	UserID            *uuid.UUID `json:"user_id,omitempty" db:"user_id"`
	CauseID           *uuid.UUID `json:"cause_id,omitempty" db:"cause_id"`
	FullName          string     `json:"full_name" db:"full_name"`
	Age               int        `json:"age" db:"age"`
	BloodGroup        string     `json:"blood_group" db:"blood_group"`
	Phone             string     `json:"phone" db:"phone"`
	Email             *string    `json:"email,omitempty" db:"email"`
	Village           *string    `json:"village,omitempty" db:"village"`
	City              *string    `json:"city,omitempty" db:"city"`
	District          *string    `json:"district,omitempty" db:"district"`
	State             *string    `json:"state,omitempty" db:"state"`
	LastDonationDate  *time.Time `json:"last_donation_date,omitempty" db:"last_donation_date"`
	Availability      bool       `json:"availability" db:"availability"`
	MedicalConditions *string    `json:"medical_conditions,omitempty" db:"medical_conditions"`
	Consent           bool       `json:"consent" db:"consent"`
	Status            string     `json:"status" db:"status"`
	CreatedAt         time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at" db:"updated_at"`
}

type CreateCauseBloodRequest struct {
	CauseID           *string `json:"cause_id,omitempty"`
	FullName          string  `json:"full_name"`
	Age               int     `json:"age"`
	BloodGroup        string  `json:"blood_group"`
	Phone             string  `json:"phone"`
	Email             *string `json:"email,omitempty"`
	Village           *string `json:"village,omitempty"`
	City              *string `json:"city,omitempty"`
	District          *string `json:"district,omitempty"`
	State             *string `json:"state,omitempty"`
	LastDonationDate  *string `json:"last_donation_date,omitempty"`
	Availability      *bool   `json:"availability,omitempty"`
	MedicalConditions *string `json:"medical_conditions,omitempty"`
	Consent           bool    `json:"consent"`
}

type BloodDonationEligibilityResponse struct {
	HasVerifiedRecord          bool   `json:"has_verified_record"`
	LatestVerifiedDate         string `json:"latest_verified_date,omitempty"`
	Eligible                   bool   `json:"eligible"`
	DaysUntilEligible          int    `json:"days_until_eligible,omitempty"`
	EligibilityMessage         string `json:"eligibility_message,omitempty"`
	RequiredGapDays            int    `json:"required_gap_days"`
	HasIncompleteSubmission    bool   `json:"has_incomplete_submission,omitempty"`
}

func NormalizeBloodGroup(v string) string {
	return strings.ToUpper(strings.TrimSpace(v))
}
