package models

import (
	"time"

	"github.com/google/uuid"
)

type Disbursement struct {
	ID              uuid.UUID  `json:"id" db:"id"`
	OrganizationID  uuid.UUID  `json:"organization_id" db:"organization_id"`
	CauseID         uuid.UUID  `json:"cause_id" db:"cause_id"`
	MilestoneNumber int        `json:"milestone_number" db:"milestone_number"`
	Amount          float64    `json:"amount" db:"amount"`
	TransactionHash *string    `json:"transaction_hash,omitempty" db:"transaction_hash"`
	DisbursedAt     time.Time  `json:"disbursed_at" db:"disbursed_at"`
	CreatedAt       time.Time  `json:"created_at" db:"created_at"`
	
	// Optional joined data
	Cause        *Cause        `json:"cause,omitempty" db:"-"`
	Organization *Organization `json:"organization,omitempty" db:"-"`
}

type DisbursementResponse struct {
	ID              uuid.UUID  `json:"id"`
	OrganizationID  uuid.UUID  `json:"organization_id"`
	CauseID         uuid.UUID  `json:"cause_id"`
	CauseName       string     `json:"cause_name"`
	MilestoneNumber int        `json:"milestone_number"`
	MilestoneLabel  string     `json:"milestone_label"`
	Amount          float64    `json:"amount"`
	TransactionHash *string    `json:"transaction_hash,omitempty"`
	DisbursedAt     time.Time  `json:"disbursed_at"`
}

func (d *Disbursement) ToResponse() DisbursementResponse {
	milestoneLabel := ""
	switch d.MilestoneNumber {
	case 1:
		milestoneLabel = "25% Milestone"
	case 2:
		milestoneLabel = "50% Milestone"
	case 3:
		milestoneLabel = "75% Milestone"
	case 4:
		milestoneLabel = "100% Milestone (Complete)"
	}
	
	causeName := ""
	if d.Cause != nil {
		causeName = d.Cause.Title
	}
	
	return DisbursementResponse{
		ID:              d.ID,
		OrganizationID:  d.OrganizationID,
		CauseID:         d.CauseID,
		CauseName:       causeName,
		MilestoneNumber: d.MilestoneNumber,
		MilestoneLabel:  milestoneLabel,
		Amount:          d.Amount,
		TransactionHash: d.TransactionHash,
		DisbursedAt:     d.DisbursedAt,
	}
}
