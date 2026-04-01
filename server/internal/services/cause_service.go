package services

import (
	"context"
	// new {
	"fmt"
	// }
	"os"
	"path/filepath"
	"strings"
	"time"

	"server/internal/models"
	"server/internal/repository"

	"github.com/google/uuid"
)

type CauseService interface {
	Create(ctx context.Context, req *models.CreateCauseRequest) (*models.Cause, error)
	CreateCauseBlood(ctx context.Context, userID uuid.UUID, req *models.CreateCauseBloodRequest) (*models.CauseBlood, error)
	CheckBloodDonationEligibility(ctx context.Context, userID uuid.UUID) (*models.BloodDonationEligibilityResponse, error)

	GetByID(ctx context.Context, id uuid.UUID) (*models.Cause, error)
	GetByOrganizationID(ctx context.Context, id uuid.UUID) ([]*models.Cause, error)
	GetByDomainID(ctx context.Context, id uuid.UUID) ([]*models.Cause, error)
	GetByAidTypeID(ctx context.Context, id uuid.UUID) ([]*models.Cause, error)
	GetAll(ctx context.Context) ([]*models.Cause, error)

	// Update(ctx context.Context, cause *models.Cause) error
	Delete(ctx context.Context, id uuid.UUID) error

	// Structured updates
	CreateUpdate(ctx context.Context, causeID uuid.UUID, req *models.CreateCauseUpdateRequest) (*models.CauseUpdate, error)
	// new {
	// Receipt verification jobs (async)
	StartReceiptVerificationJob(ctx context.Context, organizationID uuid.UUID, receiptPath string, claimedAmount float64) (uuid.UUID, error)
	GetReceiptVerificationStatus(ctx context.Context, organizationID uuid.UUID, receiptJobID uuid.UUID) (*models.ReceiptStatusResponse, error)
	// }
	GetDomains(ctx context.Context) ([]*models.CauseCategory, error)
	GetAidTypes(ctx context.Context) ([]*models.CauseCategory, error)
	GetDomainByID(ctx context.Context, id uuid.UUID) (*models.CauseCategory, error)
	GetAidTypeByID(ctx context.Context, id uuid.UUID) (*models.CauseCategory, error)
}

type causeService struct {
	causeRepo repository.CauseRepository
}

func NewCauseService(causeRepo repository.CauseRepository) *causeService {
	return &causeService{
		causeRepo: causeRepo,
	}
}

// new {
func (c *causeService) StartReceiptVerificationJob(
	ctx context.Context,
	organizationID uuid.UUID,
	receiptPath string,
	claimedAmount float64,
) (uuid.UUID, error) {
	jobID := uuid.New()

	job := &models.ReceiptVerificationJob{
		ID:             jobID,
		OrganizationID: organizationID,
		ReceiptPath:    receiptPath,
		ClaimedAmount:  claimedAmount,
		Status:         "pending",
	}

	if err := c.causeRepo.CreateReceiptVerificationJob(ctx, job); err != nil {
		return uuid.Nil, err
	}

	// Run AI verification asynchronously. We use a background context so the
	// goroutine is not cancelled when the HTTP request completes.
	go func() {
		aiCtx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
		defer cancel()

		ai, err := CallAIReceiptService(receiptPath, claimedAmount)
		if err != nil {
			errMsg := err.Error()
			_ = c.causeRepo.UpdateReceiptVerificationJobResult(aiCtx, organizationID, jobID, "error", nil, &errMsg)
			return
		}

		var (
			status string
			score  *float64
		)

		if s, ok := ai["status"].(string); ok && strings.TrimSpace(s) != "" {
			status = s
		}
		if rs, ok := ai["receipt_score"].(float64); ok {
			score = &rs
		}

		if strings.TrimSpace(status) == "" {
			status = "review"
		}

		_ = c.causeRepo.UpdateReceiptVerificationJobResult(aiCtx, organizationID, jobID, status, score, nil)
	}()

	return jobID, nil
}

func (c *causeService) GetReceiptVerificationStatus(
	ctx context.Context,
	organizationID uuid.UUID,
	receiptJobID uuid.UUID,
) (*models.ReceiptStatusResponse, error) {
	job, err := c.causeRepo.GetReceiptVerificationJob(ctx, organizationID, receiptJobID)
	if err != nil || job == nil {
		return nil, err
	}

	return &models.ReceiptStatusResponse{
		ReceiptJobID:  job.ID,
		Status:        job.Status,
		ReceiptScore:  job.ReceiptScore,
		ErrorMessage:  job.ErrorMessage,
		ClaimedAmount: job.ClaimedAmount,
	}, nil
}

// }
func (c *causeService) Create(ctx context.Context, req *models.CreateCauseRequest) (*models.Cause, error) {
	domain, err := c.GetDomainByID(ctx, req.DomainID)

	if err != nil {
		return nil, err
	}

	aidType, err := c.GetAidTypeByID(ctx, req.AidTypeID)

	if err != nil {
		return nil, err
	}

	organization := &models.OrganizationInCause{
		ID: ctx.Value("organizationID").(uuid.UUID),
	}

	now := req.CreatedAt

	cause := &models.Cause{
		ID:                    uuid.New(),
		Organization:          *organization,
		Title:                 req.Title,
		Description:           req.Description,
		Domain:                *domain,
		AidType:               *aidType,
		CollectedAmount:       req.CollectedAmount,
		GoalAmount:            req.GoalAmount,
		Deadline:              req.Deadline,
		CreatedAt:             now,
		IsActive:              req.IsActive,
		CoverImageURL:         req.CoverImageURL,
		ExecutionLat:          req.ExecutionLat,
		ExecutionLng:          req.ExecutionLng,
		ExecutionRadiusMeters: req.ExecutionRadiusMeters,
		ExecutionStartTime:    req.ExecutionStartTime,
		ExecutionEndTime:      req.ExecutionEndTime,
		FundingStatus:         req.FundingStatus,

		BeneficiariesCount: valueOrDefaultInt(req.BeneficiariesCount, 0),
		ExecutionLocation:  req.ExecutionLocation,
		ImpactGoal:         req.ImpactGoal,
		ProblemStatement:   req.ProblemStatement,
		ExecutionPlan:      req.ExecutionPlan,
		DonorCount:         0,
		UpdatedAt:          now,
	}

	err = c.causeRepo.Create(ctx, cause)

	if err != nil {
		return nil, err
	}

	// Optionally create structured products if provided
	if len(req.Products) > 0 {
		for _, p := range req.Products {
			if p == nil {
				continue
			}
			product := &models.CauseProduct{
				ID:             uuid.New(),
				CauseID:        cause.ID,
				Name:           p.Name,
				Description:    p.Description,
				PricePerUnit:   p.PricePerUnit,
				QuantityNeeded: p.QuantityNeeded,
				QuantityFunded: 0,
				ImageURL:       valueOrDefaultString(p.ImageURL, ""),
				CreatedAt:      now,
			}
			if err := c.causeRepo.CreateProduct(ctx, product); err != nil {
				return nil, err
			}
			cause.Products = append(cause.Products, product)
		}
	}

	return cause, nil
}

func (c *causeService) CreateCauseBlood(ctx context.Context, userID uuid.UUID, req *models.CreateCauseBloodRequest) (*models.CauseBlood, error) {
	if userID == uuid.Nil {
		return nil, fmt.Errorf("user_id is required")
	}

	fullName := strings.TrimSpace(req.FullName)
	phone := strings.TrimSpace(req.Phone)
	bloodGroup := models.NormalizeBloodGroup(req.BloodGroup)

	if fullName == "" || phone == "" || bloodGroup == "" {
		return nil, fmt.Errorf("full_name, phone and blood_group are required")
	}

	eligibility, err := c.CheckBloodDonationEligibility(ctx, userID)
	if err != nil {
		return nil, err
	}
	if !eligibility.Eligible {
		return nil, fmt.Errorf(eligibility.EligibilityMessage)
	}
	if req.Age <= 0 {
		return nil, fmt.Errorf("age must be greater than zero")
	}
	if !req.Consent {
		return nil, fmt.Errorf("consent must be accepted")
	}

	var causeID *uuid.UUID
	if req.CauseID != nil {
		s := strings.TrimSpace(*req.CauseID)
		if s != "" {
			id, err := uuid.Parse(s)
			if err != nil {
				return nil, fmt.Errorf("invalid cause_id")
			}
			causeID = &id
		}
	}

	var lastDonation *time.Time
	if req.LastDonationDate != nil {
		s := strings.TrimSpace(*req.LastDonationDate)
		if s != "" {
			t, err := time.Parse(time.RFC3339, s)
			if err != nil {
				t, err = time.ParseInLocation("2006-01-02", s, time.UTC)
			}
			if err != nil {
				return nil, fmt.Errorf("invalid last_donation_date")
			}
			lastDonation = &t
		}
	}

	now := time.Now()
	availability := true
	if req.Availability != nil {
		availability = *req.Availability
	}

	uid := userID
	blood := &models.CauseBlood{
		ID:                uuid.New(),
		UserID:            &uid,
		CauseID:           causeID,
		FullName:          fullName,
		Age:               req.Age,
		BloodGroup:        bloodGroup,
		Phone:             phone,
		Email:             req.Email,
		Village:           req.Village,
		City:              req.City,
		District:          req.District,
		State:             req.State,
		LastDonationDate:  lastDonation,
		Availability:      availability,
		MedicalConditions: req.MedicalConditions,
		Consent:           req.Consent,
		Status:            "pending",
		CreatedAt:         now,
		UpdatedAt:         now,
	}

	if err := c.causeRepo.CreateCauseBlood(ctx, blood); err != nil {
		return nil, err
	}

	return blood, nil
}

func (c *causeService) CheckBloodDonationEligibility(ctx context.Context, userID uuid.UUID) (*models.BloodDonationEligibilityResponse, error) {
	if userID == uuid.Nil {
		return nil, fmt.Errorf("user_id is required")
	}

	const requiredGapDays = 90

	incomplete, err := c.causeRepo.UserHasIncompleteBloodDonationSubmission(ctx, userID)
	if err != nil {
		return nil, err
	}
	if incomplete {
		return &models.BloodDonationEligibilityResponse{
			HasIncompleteSubmission: true,
			Eligible:                false,
			RequiredGapDays:           requiredGapDays,
			EligibilityMessage:        "You already have a blood donation submission in progress. Please wait until it is completed before submitting again.",
		}, nil
	}

	latestVerifiedDate, err := c.causeRepo.GetLatestVerifiedBloodDonationDateByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if latestVerifiedDate == nil {
		return &models.BloodDonationEligibilityResponse{
			HasVerifiedRecord:  false,
			Eligible:           true,
			RequiredGapDays:    requiredGapDays,
			EligibilityMessage: "No verified donation record found. You can submit your donation details.",
		}, nil
	}

	lastDateUTC := latestVerifiedDate.UTC()
	nextEligibleDate := lastDateUTC.AddDate(0, 0, requiredGapDays)
	nowUTC := time.Now().UTC()

	if nowUTC.Before(nextEligibleDate) {
		daysRemaining := int(nextEligibleDate.Sub(nowUTC).Hours() / 24)
		if nextEligibleDate.Sub(nowUTC).Hours() > float64(daysRemaining*24) {
			daysRemaining++
		}

		return &models.BloodDonationEligibilityResponse{
			HasVerifiedRecord:  true,
			LatestVerifiedDate: lastDateUTC.Format("2006-01-02"),
			Eligible:           false,
			DaysUntilEligible:  daysRemaining,
			RequiredGapDays:    requiredGapDays,
			EligibilityMessage: fmt.Sprintf("You are not eligible yet. Please wait %d more day(s) before donating again.", daysRemaining),
		}, nil
	}

	return &models.BloodDonationEligibilityResponse{
		HasVerifiedRecord:  true,
		LatestVerifiedDate: lastDateUTC.Format("2006-01-02"),
		Eligible:           true,
		RequiredGapDays:    requiredGapDays,
		EligibilityMessage: "You are eligible to donate.",
	}, nil
}

func (c *causeService) GetByID(ctx context.Context, id uuid.UUID) (*models.Cause, error) {
	cause, err := c.causeRepo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	// Attach products and updates for detailed campaign page
	if products, err := c.causeRepo.GetProductsByCauseID(ctx, id); err == nil {
		cause.Products = products
	}
	if updates, err := c.causeRepo.GetUpdatesByCauseID(ctx, id); err == nil {
		cause.Updates = updates
	}

	return cause, nil
}

func (c *causeService) GetByOrganizationID(ctx context.Context, organizationID uuid.UUID) ([]*models.Cause, error) {
	causesResult, err := c.causeRepo.GetByOrganizationID(ctx, organizationID)

	if err != nil {
		return nil, err
	}

	return causesResult, nil
}

func (c *causeService) GetByDomainID(ctx context.Context, domainID uuid.UUID) ([]*models.Cause, error) {
	causesResult, err := c.causeRepo.GetByDomainID(ctx, domainID)

	if err != nil {
		return nil, err
	}

	return causesResult, nil
}

func (c *causeService) GetByAidTypeID(ctx context.Context, aidTypeID uuid.UUID) ([]*models.Cause, error) {
	causesResult, err := c.causeRepo.GetByAidTypeID(ctx, aidTypeID)

	if err != nil {
		return nil, err
	}

	return causesResult, nil
}

func (c *causeService) GetAll(ctx context.Context) ([]*models.Cause, error) {
	causesResult, err := c.causeRepo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return causesResult, nil
}

func (c *causeService) Delete(ctx context.Context, id uuid.UUID) error {
	err := c.causeRepo.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}

func (c *causeService) GetDomains(ctx context.Context) ([]*models.CauseCategory, error) {
	domainResults, err := c.causeRepo.GetDomains(ctx)

	if err != nil {
		return nil, err
	}

	return domainResults, err
}

func (c *causeService) GetAidTypes(ctx context.Context) ([]*models.CauseCategory, error) {
	aidTypeResults, err := c.causeRepo.GetAidTypes(ctx)

	if err != nil {
		return nil, err
	}

	return aidTypeResults, err
}

func (c *causeService) GetDomainByID(ctx context.Context, id uuid.UUID) (*models.CauseCategory, error) {
	domain, err := c.causeRepo.GetDomainByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return domain, err
}

func (c *causeService) GetAidTypeByID(ctx context.Context, id uuid.UUID) (*models.CauseCategory, error) {
	aidType, err := c.causeRepo.GetAidTypeByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return aidType, err
}

func (c *causeService) CreateUpdate(ctx context.Context, causeID uuid.UUID, req *models.CreateCauseUpdateRequest) (*models.CauseUpdate, error) {
	now := time.Now()

	update := &models.CauseUpdate{
		ID:                 uuid.New(),
		CauseID:            causeID,
		Title:              strings.TrimSpace(req.Title),
		Description:        strings.TrimSpace(req.Description),
		UpdateType:         req.UpdateType,
		FundingPercentage:  req.FundingPercentage,
		ClaimedAmount:      req.ClaimedAmount,
		VerificationStatus: "pending",
		ProofSessionID:     req.ProofSessionID,
		CreatedAt:          now,
	}

	// new {
	// If this is an Execution update and we have receipt_job_ids, use their
	// async AI results (recommended flow).
	if strings.EqualFold(req.UpdateType, "Execution") && len(req.ReceiptJobIDs) > 0 {
		orgIDAny := ctx.Value("organizationID")
		orgID, ok := orgIDAny.(uuid.UUID)
		if !ok || orgID == uuid.Nil {
			return nil, fmt.Errorf("organizationID missing from context")
		}

		receiptJobUUIDs := make([]uuid.UUID, 0, len(req.ReceiptJobIDs))
		for _, idStr := range req.ReceiptJobIDs {
			idStr = strings.TrimSpace(idStr)
			if idStr == "" {
				continue
			}
			id, err := uuid.Parse(idStr)
			if err != nil {
				return nil, fmt.Errorf("invalid receipt_job_id: %w", err)
			}
			receiptJobUUIDs = append(receiptJobUUIDs, id)
		}

		jobs, err := c.causeRepo.GetReceiptVerificationJobsByIDs(ctx, orgID, receiptJobUUIDs)
		if err != nil {
			return nil, err
		}
		if len(jobs) != len(receiptJobUUIDs) {
			return nil, fmt.Errorf("one or more receipt verification jobs not found")
		}

		// Guard: don't allow update creation while receipts are still processing.
		for _, j := range jobs {
			switch strings.TrimSpace(strings.ToLower(j.Status)) {
			case "pending", "processing":
				return nil, fmt.Errorf("receipt verification is still in progress")
			case "error":
				return nil, fmt.Errorf("receipt verification failed")
			}
		}

		// Worst-case status aggregation.
		verificationStatus := "verified"
		for _, j := range jobs {
			switch strings.TrimSpace(strings.ToLower(j.Status)) {
			case "rejected":
				verificationStatus = "rejected"
			case "review":
				if verificationStatus != "rejected" {
					verificationStatus = "review"
				}
			}
		}

		update.VerificationStatus = verificationStatus

		// Score aggregation: average of available receipt scores.
		var (
			sum   float64
			count int
		)
		for _, j := range jobs {
			if j.ReceiptScore != nil {
				sum += *j.ReceiptScore
				count++
			}
		}
		if count > 0 {
			avg := sum / float64(count)
			update.VerificationScore = &avg
		}
	} else if strings.EqualFold(req.UpdateType, "Execution") && req.ClaimedAmount != nil && len(req.ReceiptURLs) > 0 {
		// Backward-compatible fallback: older clients may only send receipt URLs.
		// In this mode, we synchronously call the Python AI service.
	// }
		first := strings.TrimSpace(req.ReceiptURLs[0])
		if first != "" {
			// The upload endpoint returns a public URL like "/uploads/receipts/<file>".
			// Convert it back to a local filesystem path for the AI service.
			// We run the Go API from the "server" directory, so "uploads/..." is relative.
			localRel := strings.TrimPrefix(first, "/")
			localPath := filepath.FromSlash(localRel)
			if _, err := os.Stat(localPath); err == nil {
				if ai, err := CallAIReceiptService(localPath, *req.ClaimedAmount); err == nil && ai != nil {
					if score, ok := ai["receipt_score"].(float64); ok {
						update.VerificationScore = &score
					}
					if status, ok := ai["status"].(string); ok && strings.TrimSpace(status) != "" {
						update.VerificationStatus = status
					}
				}
			}
		}
	}
	update.DeriveVerificationFields()

	if err := c.causeRepo.CreateUpdate(ctx, update); err != nil {
		return nil, err
	}

	// Attach receipt media only for execution updates (business rule)
	if strings.EqualFold(req.UpdateType, "Execution") && len(req.ReceiptURLs) > 0 {
		for _, url := range req.ReceiptURLs {
			u := strings.TrimSpace(url)
			if u == "" {
				continue
			}
			media := &models.UpdateMedia{
				ID:        uuid.New(),
				UpdateID:  update.ID,
				MediaType: "receipt",
				MediaURL:  u,
				CreatedAt: now,
			}
			if err := c.causeRepo.AddUpdateMedia(ctx, media); err != nil {
				return nil, err
			}
			update.Media = append(update.Media, media)
		}
	}

	return update, nil
}

func valueOrDefaultInt(v *int, def int) int {
	if v == nil {
		return def
	}
	return *v
}

func valueOrDefaultString(v *string, def string) string {
	if v == nil {
		return def
	}
	return *v
}
