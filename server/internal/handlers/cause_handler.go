package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"server/internal/middleware"
	"server/internal/models"
	"server/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type CauseHandler struct {
	causeService       services.CauseService
	authService        services.AuthService
	jwtService         services.JWTService
	causeVoteService   services.CauseVoteService
	causeReviewService services.CauseReviewService
	ipfsService        services.IPFSService
}

func NewCauseHandler(
	causeService services.CauseService,
	authService services.AuthService,
	jwtService services.JWTService,
	causeVoteService services.CauseVoteService,
	causeReviewService services.CauseReviewService,
	ipfsService services.IPFSService,
) *CauseHandler {
	return &CauseHandler{
		causeService:       causeService,
		authService:        authService,
		jwtService:         jwtService,
		causeVoteService:   causeVoteService,
		causeReviewService: causeReviewService,
		ipfsService:        ipfsService,
	}
}

func (c *CauseHandler) RegisterRoutes(r chi.Router) {
	r.Route("/api/causes", func(r chi.Router) {

		r.Group(func(protected chi.Router) {
			protected.Use(middleware.AuthMiddleware(c.jwtService))
			protected.Post("/", c.CreateCause)
			protected.Post("/cover/upload", c.UploadCoverImage)
			protected.Post("/products/upload", c.UploadProductImage)
			protected.Post("/updates/upload/receipt", c.UploadUpdateReceipt)
			// new {
			protected.Get("/updates/receipt-status/{id}", c.GetReceiptStatus)
			// }
			protected.Post("/{ID}/updates", c.CreateCauseUpdate)
			protected.Post("/{ID}/upvote", c.UpvoteCause)
			protected.Post("/{ID}/downvote", c.DownvoteCause)
			protected.Get("/{ID}/votes", c.GetCauseVotes)
			protected.Post("/{ID}/reviews", c.CreateCauseReview)
			protected.Delete("/{ID}", c.DeleteCause)
		})

		r.Get("/", c.GetAllCauses)
		r.Get("/{ID}", c.GetCauseByID)
		r.Get("/organization/{ID}", c.GetCauseByOrganizationID)
		r.Get("/{ID}/reviews", c.GetCauseReviews)
		r.Get("/{ID}/reviews/count", c.GetCauseReviewCount)

		r.Get("/domain/{ID}", c.GetCauseByDomainID)
		r.Get("/aid/{ID}", c.GetCauseByAidTypeID)

		r.Get("/domain/{ID}", c.GetCauseByDomainID)
		r.Get("/aid/{ID}", c.GetCauseByAidTypeID)
	})

	r.Get("/api/domains", c.GetDomains)
	r.Get("/api/aids", c.GetAidTypes)
	r.Get("/api/domains/{ID}", c.GetDomainByID)
	r.Get("/api/aids/{ID}", c.GetAidTypeByID)

	// Serves raw files stored on IPFS for UI consumption.
	// The frontend uses this endpoint when rendering update receipts.
	r.Get("/api/ipfs/{CID}", c.GetIPFSContent)
}
func (c *CauseHandler) CreateCause(w http.ResponseWriter, r *http.Request) {
	var req models.CreateCauseRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	req.CreatedAt = time.Now()
	req.CollectedAmount = 0
	req.IsActive = true

	if req.Title == "" || req.DomainID.String() == "" || req.AidTypeID.String() == "" {
		http.Error(w, "Title, domainID, aidTypeId required", http.StatusBadRequest)
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())

	if !ok {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	organization, err := c.authService.GetOrganizationByID(r.Context(), userID)

	if err != nil || organization == nil || organization.User.Role != string(models.RoleTypeOrganization) {
		http.Error(w, "Not authenticated", http.StatusUnauthorized)
		return
	}

	ctx := context.WithValue(r.Context(), "organizationID", organization.ID)

	cause, err := c.causeService.Create(ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cause.ToCauseResponse())
}

func (c *CauseHandler) GetCauseByID(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if ID.String() == "" {
		http.Error(w, "CauseID required", http.StatusBadRequest)
		return
	}

	cause, err := c.causeService.GetByID(r.Context(), ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cause.ToCauseResponse())
}

func (c *CauseHandler) GetCauseByOrganizationID(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if ID.String() == "" {
		http.Error(w, "CauseID required", http.StatusBadRequest)
		return
	}

	causesResult, err := c.causeService.GetByOrganizationID(r.Context(), ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(causesResult)
}

func (c *CauseHandler) CreateCauseReview(w http.ResponseWriter, r *http.Request) {
	causeID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		http.Error(w, "Invalid cause ID", http.StatusBadRequest)
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	var req models.CreateCauseReviewRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	reviewText := strings.TrimSpace(req.ReviewText)
	if reviewText == "" || len(reviewText) < 5 {
		http.Error(w, "review_text is required (min 5 chars)", http.StatusBadRequest)
		return
	}

	canReview, err := c.causeReviewService.UserCanReviewCause(r.Context(), causeID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !canReview {
		http.Error(w, "Only users who donated to this cause can leave a review", http.StatusForbidden)
		return
	}

	review, err := c.causeReviewService.CreateReview(r.Context(), causeID, userID, reviewText)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(review)
}

func (c *CauseHandler) GetCauseReviews(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	reviews, err := c.causeReviewService.GetReviewsByCauseID(r.Context(), ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reviews)
}

func (c *CauseHandler) GetCauseReviewCount(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	count, err := c.causeReviewService.GetReviewCountByCauseID(r.Context(), ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"count": count})
}

// GetIPFSContent streams raw bytes from IPFS to the caller.
// The UI uses this endpoint to display uploaded update receipts.
func (c *CauseHandler) GetIPFSContent(w http.ResponseWriter, r *http.Request) {
	cid := chi.URLParam(r, "CID")
	if cid == "" {
		http.Error(w, "CID is required", http.StatusBadRequest)
		return
	}

	rc, err := c.ipfsService.Cat(r.Context(), cid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer rc.Close()

	// Read a small prefix to detect content-type.
	// Then re-stream prefix + remainder to the client.
	prefix := make([]byte, 512)
	n, readErr := rc.Read(prefix)
	if readErr != nil && readErr != io.EOF {
		http.Error(w, readErr.Error(), http.StatusBadRequest)
		return
	}

	if n > 0 {
		w.Header().Set("Content-Type", http.DetectContentType(prefix[:n]))
	}
	w.Header().Set("X-Content-Source", "ipfs")

	if n > 0 {
		_, _ = io.Copy(w, io.MultiReader(bytes.NewReader(prefix[:n]), rc))
	} else {
		_, _ = io.Copy(w, rc)
	}
}

func (c *CauseHandler) UpvoteCause(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	votes, err := c.causeVoteService.ToggleVote(r.Context(), ID, userID, 1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(votes)
}

func (c *CauseHandler) DownvoteCause(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	votes, err := c.causeVoteService.ToggleVote(r.Context(), ID, userID, -1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(votes)
}

func (c *CauseHandler) GetCauseVotes(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	votes, err := c.causeVoteService.GetVotes(r.Context(), ID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(votes)
}

func (c *CauseHandler) GetCauseByDomainID(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if ID.String() == "" {
		http.Error(w, "CauseID required", http.StatusBadRequest)
		return
	}

	causesResult, err := c.causeService.GetByDomainID(r.Context(), ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(causesResult)
}

func (c *CauseHandler) GetCauseByAidTypeID(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if ID.String() == "" {
		http.Error(w, "CauseID required", http.StatusBadRequest)
		return
	}

	causesResult, err := c.causeService.GetByAidTypeID(r.Context(), ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(causesResult)
}

func (c *CauseHandler) GetAllCauses(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")

	var ctx context.Context

	if limit != "" {
		_, err := strconv.Atoi(limit)
		if err != nil {
			http.Error(w, "Limit Parameter is invalid", http.StatusBadRequest)
			return
		}
		ctx = context.WithValue(r.Context(), "limit", limit)
	} else {
		ctx = r.Context()
	}

	causesResult, err := c.causeService.GetAll(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(causesResult)
}

func (c *CauseHandler) DeleteCause(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if ID.String() == "" {
		http.Error(w, "CauseID required", http.StatusBadRequest)
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())

	if !ok {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	organization, err := c.authService.GetOrganizationByID(r.Context(), userID)
	if err != nil || organization == nil || organization.User.Role != string(models.RoleTypeOrganization) {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	cause, err := c.causeService.GetByID(r.Context(), ID)

	if err != nil || cause.Organization.ID != organization.ID {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	err = c.causeService.Delete(r.Context(), ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "deleted cause successfully",
	})
}

func (c *CauseHandler) GetDomains(w http.ResponseWriter, r *http.Request) {
	domainsResult, err := c.causeService.GetDomains(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(domainsResult)
}

func (c *CauseHandler) GetAidTypes(w http.ResponseWriter, r *http.Request) {
	aidTypesResults, err := c.causeService.GetAidTypes(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aidTypesResults)
}

func (c *CauseHandler) GetDomainByID(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if ID.String() == "" {
		http.Error(w, "CauseID required", http.StatusBadRequest)
		return
	}

	domain, err := c.causeService.GetDomainByID(r.Context(), ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(domain)
}

func (c *CauseHandler) GetAidTypeByID(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if ID.String() == "" {
		http.Error(w, "CauseID required", http.StatusBadRequest)
		return
	}

	aidType, err := c.causeService.GetAidTypeByID(r.Context(), ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aidType)
}

// UploadUpdateReceipt handles upload of receipt images for execution updates.
// Works similarly to UploadProductImage but stores under uploads/receipts.
func (c *CauseHandler) UploadUpdateReceipt(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	org, err := c.authService.GetOrganizationByID(r.Context(), userID)
	if err != nil || org == nil || org.User.Role != string(models.RoleTypeOrganization) {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}

	// new {
	claimedAmountStr := strings.TrimSpace(r.FormValue("claimed_amount"))
	if claimedAmountStr == "" {
		http.Error(w, "claimed_amount is required", http.StatusBadRequest)
		return
	}
	claimedAmount, err := strconv.ParseFloat(claimedAmountStr, 64)
	if err != nil || claimedAmount <= 0 {
		http.Error(w, "claimed_amount must be a positive number", http.StatusBadRequest)
		return
	}
	// }
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Receipt image required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	sniff := make([]byte, 512)
	n, err := file.Read(sniff)
	if err != nil && !errors.Is(err, io.EOF) {
		http.Error(w, "Failed to read image", http.StatusBadRequest)
		return
	}
	if n == 0 {
		http.Error(w, "Empty image", http.StatusBadRequest)
		return
	}

	contentType := http.DetectContentType(sniff)
	if !strings.HasPrefix(contentType, "image/") {
		http.Error(w, "Only image uploads are allowed", http.StatusBadRequest)
		return
	}

	imageBytes, err := io.ReadAll(io.MultiReader(bytes.NewReader(sniff[:n]), file))
	if err != nil {
		http.Error(w, "Failed to read full image", http.StatusBadRequest)
		return
	}

	// Store the receipt securely on IPFS and return an API URL that streams it back.
	cid, err := c.ipfsService.AddFile(r.Context(), bytes.NewReader(imageBytes))
	if err != nil {
		http.Error(w, "Failed to store on IPFS: "+err.Error(), http.StatusBadRequest)
		return
	}

	publicPath := "/api/ipfs/" + cid

	// new {
	// Persist a DB-backed receipt verification job and trigger AI analysis async.

	receiptJobID, err := c.causeService.StartReceiptVerificationJob(r.Context(), org.ID, publicPath, claimedAmount)
	if err != nil {
		http.Error(w, "Failed to create receipt verification job", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]any{
		"url":           "/" + publicPath,
		"receipt_id":    receiptJobID.String(),
		"status":        "pending",
		"receipt_score": nil,
	})
}

// new {
func (c *CauseHandler) GetReceiptStatus(w http.ResponseWriter, r *http.Request) {
	receiptJobID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid receipt id", http.StatusBadRequest)
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	org, err := c.authService.GetOrganizationByID(r.Context(), userID)
	if err != nil || org == nil || org.User.Role != string(models.RoleTypeOrganization) {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	status, err := c.causeService.GetReceiptVerificationStatus(r.Context(), org.ID, receiptJobID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

// }
// CreateCauseUpdate creates a structured engagement/execution update for a cause.
func (c *CauseHandler) CreateCauseUpdate(w http.ResponseWriter, r *http.Request) {
	causeID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		http.Error(w, "Invalid cause ID", http.StatusBadRequest)
		return
	}

	var req models.CreateCauseUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	req.Title = strings.TrimSpace(req.Title)
	req.Description = strings.TrimSpace(req.Description)
	if req.Title == "" || req.Description == "" || req.UpdateType == "" {
		http.Error(w, "title, description and update_type are required", http.StatusBadRequest)
		return
	}

	allowedTypes := map[string]bool{
		"Engagement": true,
		"Milestone":  true,
		"Execution":  true,
		"Completion": true,
	}
	if !allowedTypes[req.UpdateType] {
		http.Error(w, "invalid update_type", http.StatusBadRequest)
		return
	}

	// Execution updates require a claimed amount (used for receipt verification)
	if strings.EqualFold(req.UpdateType, "Execution") {
		if req.ClaimedAmount == nil || *req.ClaimedAmount <= 0 {
			http.Error(w, "claimed_amount is required for Execution updates", http.StatusBadRequest)
			return
		}
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	org, err := c.authService.GetOrganizationByID(r.Context(), userID)
	if err != nil || org == nil || org.User.Role != string(models.RoleTypeOrganization) {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	cause, err := c.causeService.GetByID(r.Context(), causeID)
	if err != nil || cause == nil {
		http.Error(w, "Cause not found", http.StatusBadRequest)
		return
	}
	if cause.Organization.ID != org.ID {
		http.Error(w, "Not authorized for this cause", http.StatusForbidden)
		return
	}

	// update, err := c.causeService.CreateUpdate(r.Context(), causeID, &req)
	// new {
	ctx := context.WithValue(r.Context(), "organizationID", org.ID)
	update, err := c.causeService.CreateUpdate(ctx, causeID, &req)
	// }
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(update)
}

// UploadCoverImage handles secure upload of a single campaign cover image.
// It validates image content type and stores the file under uploads/covers,
// returning a public URL path that can be saved as cover_image_url.
func (c *CauseHandler) UploadCoverImage(w http.ResponseWriter, r *http.Request) {
	// Ensure requester is an authenticated organization (same as CreateCause)
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	org, err := c.authService.GetOrganizationByID(r.Context(), userID)
	if err != nil || org == nil || org.User.Role != string(models.RoleTypeOrganization) {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Cover image required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Read a small buffer to detect content type safely
	sniff := make([]byte, 512)
	n, err := file.Read(sniff)
	if err != nil && !errors.Is(err, io.EOF) {
		http.Error(w, "Failed to read image", http.StatusBadRequest)
		return
	}
	if n == 0 {
		http.Error(w, "Empty image", http.StatusBadRequest)
		return
	}

	contentType := http.DetectContentType(sniff)
	if !strings.HasPrefix(contentType, "image/") {
		http.Error(w, "Only image uploads are allowed", http.StatusBadRequest)
		return
	}

	// Reset reader to include the sniffed bytes
	imageBytes, err := io.ReadAll(io.MultiReader(strings.NewReader(string(sniff[:n])), file))
	if err != nil {
		http.Error(w, "Failed to read full image", http.StatusBadRequest)
		return
	}

	// Derive safe extension
	exts, _ := mime.ExtensionsByType(contentType)
	ext := ""
	if len(exts) > 0 {
		ext = exts[0]
	}
	if ext == "" {
		ext = filepath.Ext(header.Filename)
	}
	if ext == "" {
		ext = ".img"
	}

	dir := "uploads/covers"
	_ = os.MkdirAll(dir, 0755)

	filename := uuid.New().String() + ext
	path := filepath.Join(dir, filename)

	if err := os.WriteFile(path, imageBytes, 0644); err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	publicPath := filepath.ToSlash(filepath.Join("uploads", "covers", filename))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"url": "/" + publicPath,
	})
}

// UploadProductImage handles secure upload of a single product image for a cause.
// It mirrors the cover image upload but stores under uploads/products and returns
// a URL that can be saved in the cause_products.image_url column.
func (c *CauseHandler) UploadProductImage(w http.ResponseWriter, r *http.Request) {
	// Reuse the same authentication as CreateCause
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	org, err := c.authService.GetOrganizationByID(r.Context(), userID)
	if err != nil || org == nil || org.User.Role != string(models.RoleTypeOrganization) {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Product image required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	sniff := make([]byte, 512)
	n, err := file.Read(sniff)
	if err != nil && !errors.Is(err, io.EOF) {
		http.Error(w, "Failed to read image", http.StatusBadRequest)
		return
	}
	if n == 0 {
		http.Error(w, "Empty image", http.StatusBadRequest)
		return
	}

	contentType := http.DetectContentType(sniff)
	if !strings.HasPrefix(contentType, "image/") {
		http.Error(w, "Only image uploads are allowed", http.StatusBadRequest)
		return
	}

	imageBytes, err := io.ReadAll(io.MultiReader(strings.NewReader(string(sniff[:n])), file))
	if err != nil {
		http.Error(w, "Failed to read full image", http.StatusBadRequest)
		return
	}

	exts, _ := mime.ExtensionsByType(contentType)
	ext := ""
	if len(exts) > 0 {
		ext = exts[0]
	}
	if ext == "" {
		ext = filepath.Ext(header.Filename)
	}
	if ext == "" {
		ext = ".img"
	}

	dir := "uploads/products"
	_ = os.MkdirAll(dir, 0755)

	filename := uuid.New().String() + ext
	path := filepath.Join(dir, filename)

	if err := os.WriteFile(path, imageBytes, 0644); err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	publicPath := filepath.ToSlash(filepath.Join("uploads", "products", filename))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"url": "/" + publicPath,
	})
}
