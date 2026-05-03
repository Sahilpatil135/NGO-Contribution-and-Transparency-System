package handlers

import (
	"encoding/json"
	"net/http"

	"server/internal/middleware"
	"server/internal/repository"
	"server/internal/services"

	"github.com/go-chi/chi/v5"
)

type AdminHandler struct {
	adminRepo  repository.AdminRepository
	jwtService services.JWTService
}

func NewAdminHandler(adminRepo repository.AdminRepository, jwtService services.JWTService) *AdminHandler {
	return &AdminHandler{
		adminRepo:  adminRepo,
		jwtService: jwtService,
	}
}

func (h *AdminHandler) RegisterRoutes(r chi.Router) {
	r.Route("/api/admin", func(r chi.Router) {
		// All admin routes require a valid JWT and the "admin" role
		r.Group(func(protected chi.Router) {
			protected.Use(middleware.AuthMiddleware(h.jwtService))
			protected.Use(middleware.RequireRole("admin"))
			protected.Get("/dashboard", h.GetDashboardData)
		})
	})
}

func (h *AdminHandler) GetDashboardData(w http.ResponseWriter, r *http.Request) {
	data, err := h.adminRepo.GetDashboardData(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch dashboard data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
