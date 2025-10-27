package models

import (
	"time"

	"github.com/google/uuid"
)

type DonationStatus string

const (
	DonationStatusPending   DonationStatus = "pending"
	DonationStatusCompleted DonationStatus = "paid"
	DonationStatusFailed    DonationStatus = "failed"
)

type Donation struct {
	ID             uuid.UUID      `json:"id" db:"id"`
	CauseID        uuid.UUID      `json:"cause_id" db:"cause_id"`
	UserID         uuid.UUID      `json:"user_id" db:"user_id"`
	Name           string         `json:"name" db:"name"`
	Phone          string         `json:"phone" db:"phone"`
	BillingAddress *string        `json:"billing_address,omitempty" db:"billing_address"`
	Pincode        *string        `json:"pincode,omitempty" db:"pincode"`
	Amount         float32        `json:"amount" db:"amount"`
	Status         DonationStatus `json:"status" db:"status"`
	PanNumber      *string        `json:"pan_number,omitempty" db:"pan_number"`
	PaymentID      *string        `json:"payment_id,omitempty" db:"payment_id"`
	CreatedAt      time.Time      `json:"created_at" db:"created_at"`
}

type CreateDonationRequest struct {
	CauseID        uuid.UUID `json:"cause_id" validate:"required,uuid"`
	UserID         uuid.UUID `json:"user_id" validate:"required,uuid"`
	Name           *string   `json:"name,omitempty"`
	Phone          string    `json:"phone" validate:"required"`
	BillingAddress *string   `json:"billing_address,omitempty"`
	Pincode        *string   `json:"pincode,omitempty"`
	Amount         float32   `json:"amount" validate:"required,gt=0"`
	PanNumber      *string   `json:"pan_number,omitempty"`
	PaymentID      *string   `json:"payment_id,omitempty"`
}

type CreateDonationResponse struct {
	ID             uuid.UUID      `json:"id"`
	CauseID        uuid.UUID      `json:"cause_id"`
	UserID         uuid.UUID      `json:"user_id"`
	Name           string         `json:"name"`
	Phone          string         `json:"phone"`
	BillingAddress *string        `json:"billing_address,omitempty"`
	Pincode        *string        `json:"pincode,omitempty"`
	Amount         float32        `json:"amount"`
	Status         DonationStatus `json:"status"`
}

func (d *Donation) ToDonationResponse() *CreateDonationResponse {
	return &CreateDonationResponse{
		ID:             d.ID,
		CauseID:        d.CauseID,
		UserID:         d.UserID,
		Name:           d.Name,
		Phone:          d.Phone,
		BillingAddress: d.BillingAddress,
		Pincode:        d.Pincode,
		Amount:         d.Amount,
		Status:         d.Status,
	}
}
