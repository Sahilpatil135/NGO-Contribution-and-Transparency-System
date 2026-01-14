package models

import "time"

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
