package models

import (
	"time"

	"github.com/google/uuid"
)

type ProofSessionResponse struct {
	SessionID string    `json:"sessionId"`
	ExpiresAt time.Time `json:"expiresAt"`
	QRURL     string    `json:"qrUrl"`
}

type ProofUploadEvent struct {
	ImagePath string    `json:"image"`
	Latitude  string    `json:"lat"`
	Longitude string    `json:"lng"`
	Timestamp time.Time `json:"timestamp"`
}

// ProofSession is a DB-backed proof session linked to a cause
type ProofSession struct {
	ID             uuid.UUID `json:"id" db:"id"`
	OrganizationID uuid.UUID `json:"organization_id" db:"organization_id"`
	CauseID       uuid.UUID `json:"cause_id" db:"cause_id"`
	IsActive      bool      `json:"is_active" db:"is_active"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}

// ProofImage is a stored proof image with hash and metadata score
type ProofImage struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	SessionID    uuid.UUID  `json:"session_id" db:"session_id"`
	ImageHash    string     `json:"image_hash" db:"image_hash"`
	IPFSCID      *string    `json:"ipfs_cid,omitempty" db:"ipfs_cid"`
	Latitude     *float64   `json:"latitude" db:"latitude"`
	Longitude    *float64   `json:"longitude" db:"longitude"`
	Timestamp    *time.Time `json:"timestamp" db:"timestamp"`
	MetadataScore int       `json:"metadata_score" db:"metadata_score"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
}

// CreateProofSessionRequest for starting a cause-backed proof session (optional causeId)
type CreateProofSessionRequest struct {
	CauseID *uuid.UUID `json:"causeId,omitempty"`
}

// CauseExecution holds cause execution window and location for validation
type CauseExecution struct {
	CauseID               uuid.UUID  `json:"cause_id" db:"cause_id"`
	ExecutionLat           *float64   `json:"execution_lat" db:"execution_lat"`
	ExecutionLng           *float64   `json:"execution_lng" db:"execution_lng"`
	ExecutionRadiusMeters   *int       `json:"execution_radius_meters" db:"execution_radius_meters"`
	ExecutionStartTime     *time.Time `json:"execution_start_time" db:"execution_start_time"`
	ExecutionEndTime       *time.Time `json:"execution_end_time" db:"execution_end_time"`
}

// UploadProofResponse returned after processing proof upload
type UploadProofResponse struct {
	Status        string `json:"status"`
	Score         int    `json:"score,omitempty"`
	IsDuplicate   bool   `json:"isDuplicate,omitempty"`
	ValidationOK  bool   `json:"validationOk,omitempty"`
}
