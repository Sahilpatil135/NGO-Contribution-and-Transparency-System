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

func (c *CauseHandler) RegisterRoutes(r chi.Router) {
	r.Route("/api/causes", func(r chi.Router) {

		r.Group(func(protected chi.Router) {
			protected.Use(middleware.AuthMiddleware(c.jwtService))
			protected.Post("/", c.CreateCause)
			protected.Delete("/{ID}", c.DeleteCause)
		})

		r.Get("/", c.GetAll)
		r.Get("/{ID}", c.GetCauseByID)
		r.Get("/organization/{ID}", c.GetCauseByOrganizationID)

		r.Get("/domain/{ID}", c.GetCauseByDomainID)
		r.Get("/aid/{ID}", c.GetCauseByAidTypeID)

		r.Get("/domain/{ID}", c.GetCauseByDomainID)
		r.Get("/aid/{ID}", c.GetCauseByAidTypeID)
	})

	r.Get("/api/domains", c.GetDomains)
	r.Get("/api/aids", c.GetAidTypes)
	r.Get("/api/domains/{ID}", c.GetDomainByID)
	r.Get("/api/aids/{ID}", c.GetAidTypeByID)
}
func (c *CauseHandler) CreateCause(w http.ResponseWriter, r *http.Request) {
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

func (c *CauseHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	causesResult, err := c.causeService.GetAll(r.Context())
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

	if err != nil || cause.OrganizationID != organization.ID {
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
