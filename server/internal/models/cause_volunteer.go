package models

import (
	"time"

	"github.com/google/uuid"
)

type CauseVolunteer struct {
	ID                uuid.UUID  `json:"id" db:"id"`
	UserID            *uuid.UUID `json:"user_id,omitempty" db:"user_id"`
	CauseID           *uuid.UUID `json:"cause_id,omitempty" db:"cause_id"`
	FullName          string     `json:"full_name" db:"full_name"`
	Phone             string     `json:"phone" db:"phone"`
	Email             *string    `json:"email,omitempty" db:"email"`
	Village           *string    `json:"village,omitempty" db:"village"`
	City              *string    `json:"city,omitempty" db:"city"`
	District          *string    `json:"district,omitempty" db:"district"`
	State             *string    `json:"state,omitempty" db:"state"`
	Skills            string     `json:"skills" db:"skills"`
	Interests         *string    `json:"interests,omitempty" db:"interests"`
	AvailabilityType  *string    `json:"availability_type,omitempty" db:"availability_type"`
	AvailableHours    *int       `json:"available_hours,omitempty" db:"available_hours"`
	Experience        *string    `json:"experience,omitempty" db:"experience"`
	Consent           bool       `json:"consent" db:"consent"`
	Status            string     `json:"status" db:"status"`
	CreatedAt         time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at" db:"updated_at"`
}

type CreateCauseVolunteerRequest struct {
	CauseID          *string `json:"cause_id,omitempty"`
	FullName         string  `json:"full_name"`
	Phone            string  `json:"phone"`
	Email            *string `json:"email,omitempty"`
	Village          *string `json:"village,omitempty"`
	City             *string `json:"city,omitempty"`
	District         *string `json:"district,omitempty"`
	State            *string `json:"state,omitempty"`
	Skills           string  `json:"skills"`
	Interests        *string `json:"interests,omitempty"`
	AvailabilityType *string `json:"availability_type,omitempty"`
	AvailableHours   *int    `json:"available_hours,omitempty"`
	Experience       *string `json:"experience,omitempty"`
	Consent          bool    `json:"consent"`
}
