package services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"time"

	"server/internal/models"
	"server/internal/repository"

	"github.com/google/uuid"
)

const (
	scoreTimeValid     = 40
	scoreLocationValid = 40
	scoreUnique        = 20
)

type ProofService interface {
	CreateSession(ctx context.Context, causeID, organizationID uuid.UUID) (*models.ProofSession, error)
	ProcessUpload(ctx context.Context, sessionID uuid.UUID, lat, lng float64, timestamp time.Time, imageBytes []byte) (*models.ProofImage, int, bool, bool, error)
}

type proofService struct {
	sessionRepo repository.ProofSessionRepository
	imageRepo   repository.ProofImageRepository
	causeRepo   repository.CauseRepository
}

func NewProofService(
	sessionRepo repository.ProofSessionRepository,
	imageRepo repository.ProofImageRepository,
	causeRepo repository.CauseRepository,
) ProofService {
	return &proofService{
		sessionRepo: sessionRepo,
		imageRepo:   imageRepo,
		causeRepo:   causeRepo,
	}
}

func (s *proofService) CreateSession(ctx context.Context, causeID, organizationID uuid.UUID) (*models.ProofSession, error) {
	session := &models.ProofSession{
		ID:             uuid.New(),
		OrganizationID: organizationID,
		CauseID:        causeID,
		IsActive:       true,
		CreatedAt:      time.Now(),
	}
	if err := s.sessionRepo.Create(ctx, session); err != nil {
		return nil, fmt.Errorf("create proof session: %w", err)
	}
	return session, nil
}

func (s *proofService) ProcessUpload(ctx context.Context, sessionID uuid.UUID, lat, lng float64, timestamp time.Time, imageBytes []byte) (*models.ProofImage, int, bool, bool, error) {
	session, err := s.sessionRepo.GetByID(ctx, sessionID)
	if err != nil || session == nil {
		return nil, 0, false, false, fmt.Errorf("session not found")
	}

	if len(imageBytes) == 0 {
		return nil, 0, false, false, fmt.Errorf("empty image")
	}

	hash := hashImage(imageBytes)

	// Duplicate detection
	exists, err := s.imageRepo.ExistsBySessionIDAndHash(ctx, sessionID, hash)
	if err != nil {
		return nil, 0, false, false, err
	}
	if exists {
		return nil, 0, true, false, nil // duplicate, score 0, validation not applicable
	}

	// Metadata validation and scoring
	exec, err := s.causeRepo.GetCauseExecution(ctx, session.CauseID)
	if err != nil {
		return nil, 0, false, false, err
	}

	score := 0
	timeValid := false
	locationValid := false

	if exec != nil {
		if exec.ExecutionStartTime != nil && exec.ExecutionEndTime != nil {
			timeValid = !timestamp.Before(*exec.ExecutionStartTime) && !timestamp.After(*exec.ExecutionEndTime)
		} else {
			timeValid = true // no window set => accept any time
		}
		if timeValid {
			score += scoreTimeValid
		}

		if exec.ExecutionLat != nil && exec.ExecutionLng != nil && exec.ExecutionRadiusMeters != nil {
			radius := float64(*exec.ExecutionRadiusMeters)
			if radius <= 0 {
				radius = 200
			}
			dist := haversineMeters(lat, lng, *exec.ExecutionLat, *exec.ExecutionLng)
			locationValid = dist <= radius
		} else {
			locationValid = true // no geo set => accept any location
		}
		if locationValid {
			score += scoreLocationValid
		}
	} else {
		timeValid = true
		locationValid = true
		score = scoreTimeValid + scoreLocationValid
	}

	// Unique image bonus
	score += scoreUnique

	img := &models.ProofImage{
		ID:            uuid.New(),
		SessionID:     sessionID,
		ImageHash:     hash,
		Latitude:      &lat,
		Longitude:     &lng,
		Timestamp:     &timestamp,
		MetadataScore: score,
		CreatedAt:     time.Now(),
	}
	if err := s.imageRepo.Create(ctx, img); err != nil {
		return nil, 0, false, false, fmt.Errorf("store proof image: %w", err)
	}

	validationOK := timeValid && locationValid
	return img, score, false, validationOK, nil
}

func hashImage(data []byte) string {
	h := sha256.Sum256(data)
	return hex.EncodeToString(h[:])
}

// haversineMeters returns distance in meters between two points (lat/lng in degrees).
func haversineMeters(lat1, lng1, lat2, lng2 float64) float64 {
	const earthRadiusM = 6371000
	dLat := (lat2 - lat1) * math.Pi / 180
	dLng := (lng2 - lng1) * math.Pi / 180
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*math.Pi/180)*math.Cos(lat2*math.Pi/180)*
			math.Sin(dLng/2)*math.Sin(dLng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return earthRadiusM * c
}
