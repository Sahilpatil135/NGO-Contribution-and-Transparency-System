package handlers

import (
	"encoding/json"
	"net/http"

	"server/internal/repository"

	"github.com/go-chi/chi/v5"
)

type AdminHandler struct {
	adminRepo repository.AdminRepository
}

func NewAdminHandler(adminRepo repository.AdminRepository) *AdminHandler {
	return &AdminHandler{
		adminRepo: adminRepo,
	}
}

func (h *AdminHandler) RegisterRoutes(r chi.Router) {
	r.Route("/api/admin", func(r chi.Router) {
		r.Get("/dashboard", h.GetDashboardData)
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
