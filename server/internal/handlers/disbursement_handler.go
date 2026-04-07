package handlers

import (
	"encoding/json"
	"net/http"
	"server/internal/middleware"
	"server/internal/models"
	"server/internal/repository"
	"server/internal/services"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type DisbursementHandler struct {
	disbursementRepo  repository.DisbursementRepository
	organizationRepo  repository.OrganizationRepository
	jwtService        services.JWTService
}

func NewDisbursementHandler(
	disbursementRepo repository.DisbursementRepository,
	organizationRepo repository.OrganizationRepository,
	jwtService services.JWTService,
) *DisbursementHandler {
	return &DisbursementHandler{
		disbursementRepo:  disbursementRepo,
		organizationRepo:  organizationRepo,
		jwtService:        jwtService,
	}
}

func (h *DisbursementHandler) RegisterRoutes(r chi.Router) {
	r.Route("/api/disbursements", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(h.jwtService))
		r.Get("/my-organization", h.GetMyOrganizationDisbursements)
		r.Get("/cause/{causeID}", h.GetCauseDisbursements)
	})
}

// GetMyOrganizationDisbursements returns disbursements for the authenticated organization
func (h *DisbursementHandler) GetMyOrganizationDisbursements(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Get organization by user ID
	organization, err := h.organizationRepo.GetByID(r.Context(), userID)
	if err != nil || organization == nil {
		http.Error(w, "Organization not found", http.StatusNotFound)
		return
	}

	// Parse pagination parameters
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")
	
	limit := 50
	offset := 0
	
	if limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 && parsedLimit <= 100 {
			limit = parsedLimit
		}
	}
	
	if offsetStr != "" {
		if parsedOffset, err := strconv.Atoi(offsetStr); err == nil && parsedOffset >= 0 {
			offset = parsedOffset
		}
	}

	// Get disbursements
	disbursements, err := h.disbursementRepo.GetByOrganizationID(r.Context(), organization.ID, limit, offset)
	if err != nil {
		http.Error(w, "Failed to fetch disbursements", http.StatusInternalServerError)
		return
	}

	// Get total count
	total, err := h.disbursementRepo.CountByOrganizationID(r.Context(), organization.ID)
	if err != nil {
		http.Error(w, "Failed to fetch total count", http.StatusInternalServerError)
		return
	}

	// Convert to response format
	var responses []models.DisbursementResponse
	for _, d := range disbursements {
		responses = append(responses, d.ToResponse())
	}

	// Return response
	response := map[string]interface{}{
		"disbursements": responses,
		"total":         total,
		"limit":         limit,
		"offset":        offset,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetCauseDisbursements returns disbursements for a specific cause (public endpoint)
func (h *DisbursementHandler) GetCauseDisbursements(w http.ResponseWriter, r *http.Request) {
	causeIDStr := chi.URLParam(r, "causeID")
	if causeIDStr == "" {
		http.Error(w, "Cause ID is required", http.StatusBadRequest)
		return
	}

	causeID, err := uuid.Parse(causeIDStr)
	if err != nil {
		http.Error(w, "Invalid cause ID", http.StatusBadRequest)
		return
	}

	// Get disbursements for this cause
	disbursements, err := h.disbursementRepo.GetByCauseID(r.Context(), causeID)
	if err != nil {
		http.Error(w, "Failed to fetch disbursements", http.StatusInternalServerError)
		return
	}

	// Convert to response format
	var responses []models.DisbursementResponse
	totalDisbursed := 0.0
	for _, d := range disbursements {
		responses = append(responses, d.ToResponse())
		totalDisbursed += d.Amount
	}

	// Return response
	response := map[string]interface{}{
		"disbursements":   responses,
		"total_disbursed": totalDisbursed,
		"count":           len(disbursements),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
