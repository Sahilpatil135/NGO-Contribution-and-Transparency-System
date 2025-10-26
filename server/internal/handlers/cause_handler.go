package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"server/internal/middleware"
	"server/internal/models"
	"server/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type CauseHandler struct {
	causeService services.CauseService
	authService  services.AuthService
	jwtService   services.JWTService
}

func NewCauseHandler(causeService services.CauseService, authService services.AuthService, jwtService services.JWTService) *CauseHandler {
	return &CauseHandler{
		causeService: causeService,
		authService:  authService,
		jwtService:   jwtService,
	}
}

func (h *CauseHandler) RegisterRoutes(r chi.Router) {
	r.Route("/api/causes", func(r chi.Router) {

		r.Group(func(protected chi.Router) {
			protected.Use(middleware.AuthMiddleware(h.jwtService))
			protected.Post("/", h.CreateCause)
			protected.Delete("/{ID}", h.DeleteCause)
		})

		r.Get("/", h.GetAll)
		r.Get("/{ID}", h.GetCauseByID)
		r.Get("/organization/{ID}", h.GetCauseByOrganizationID)
		r.Get("/domain/{ID}", h.GetCauseByDomainID)
		r.Get("/aid/{ID}", h.GetCauseByAidTypeID)
	})
}

func (h *CauseHandler) CreateCause(w http.ResponseWriter, r *http.Request) {
	var req models.CreateCauseRequest

	req.CreatedAt = time.Now()
	req.CollectedAmount = 0
	req.IsActive = true

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Title == "" || req.DomainID.String() == "" || req.AidTypeID.String() == "" {
		http.Error(w, "Title, domainID, aidTypeId required", http.StatusBadRequest)
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())

	if !ok {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	organization, err := h.authService.GetOrganizationByID(r.Context(), userID)

	if err != nil || organization == nil || organization.User.Role != "organization" {
		http.Error(w, "Not authenticated", http.StatusUnauthorized)
		return
	}

	ctx := context.WithValue(r.Context(), "organizationID", organization.ID)

	cause, err := h.causeService.Create(ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cause.ToCauseResponse())
}

func (h *CauseHandler) GetCauseByID(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if ID.String() == "" {
		http.Error(w, "CauseID required", http.StatusBadRequest)
		return
	}

	cause, err := h.causeService.GetByID(r.Context(), ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cause.ToCauseResponse())
}

func (h *CauseHandler) GetCauseByOrganizationID(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if ID.String() == "" {
		http.Error(w, "CauseID required", http.StatusBadRequest)
		return
	}

	causesResult, err := h.causeService.GetByOrganizationID(r.Context(), ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(causesResult)
}

func (h *CauseHandler) GetCauseByDomainID(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if ID.String() == "" {
		http.Error(w, "CauseID required", http.StatusBadRequest)
		return
	}

	causesResult, err := h.causeService.GetByDomainID(r.Context(), ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(causesResult)
}

func (h *CauseHandler) GetCauseByAidTypeID(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if ID.String() == "" {
		http.Error(w, "CauseID required", http.StatusBadRequest)
		return
	}

	causesResult, err := h.causeService.GetByAidTypeID(r.Context(), ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(causesResult)
}

func (h *CauseHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	causesResult, err := h.causeService.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(causesResult)
}

func (h *CauseHandler) DeleteCause(w http.ResponseWriter, r *http.Request) {
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

	organization, err := h.authService.GetOrganizationByID(r.Context(), userID)
	if err != nil || organization == nil || organization.User.Role != "organization" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	cause, err := h.causeService.GetByID(r.Context(), ID)

	if err != nil || cause.OrganizationID != organization.ID {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	err = h.causeService.Delete(r.Context(), ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "deleted cause successfully",
	})
}
