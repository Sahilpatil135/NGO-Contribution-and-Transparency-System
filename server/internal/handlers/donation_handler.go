package handlers

import (
	"encoding/json"
	"net/http"

	"server/internal/middleware"
	"server/internal/models"
	"server/internal/services"

	"github.com/go-chi/chi/v5"
)

type DonationHandler struct {
	donationService services.DonationService
	authService     services.AuthService
	jwtService      services.JWTService
}

func NewDonationHandler(donationService services.DonationService, authService services.AuthService, jwtService services.JWTService) *DonationHandler {
	return &DonationHandler{
		donationService: donationService,
		authService:     authService,
		jwtService:      jwtService,
	}
}

func (c *DonationHandler) RegisterRoutes(r chi.Router) {
	r.Route("/api/donations", func(r chi.Router) {

		r.Group(func(protected chi.Router) {
			protected.Use(middleware.AuthMiddleware(c.jwtService))
			protected.Post("/", c.CreateDonation)
			// protected.Delete("/{ID}", c.DeleteDonation)
		})

		// r.Get("/", c.GetAll)
		// r.Get("/{ID}", c.GetDonationByID)
		// r.Get("/cause/{ID}", c.GetDonationByCauseID)
		// r.Get("/user/{ID}", c.GetDonationByUserID)
	})
}
func (c *DonationHandler) CreateDonation(w http.ResponseWriter, r *http.Request) {
	var req models.CreateDonationRequest

	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "User not found", http.StatusUnauthorized)
	}

	req.UserID = userID

	user, err := c.authService.GetUserByID(r.Context(), userID)
	if err != nil {
		http.Error(w, "User not available", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if *req.Name == "" {
		req.Name = &user.Name
	}

	if req.CauseID.String() == "" || req.UserID.String() == "" || *req.Name == "" || req.Phone == "" || req.Amount == 0.0 {
		http.Error(w, "CauseID, UserID, Name, Phone, Amount is required", http.StatusBadRequest)
		return
	}

	donation, err := c.donationService.Create(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(donation.ToDonationResponse())
}

//
// func (c *DonationHandler) GetDonationByID(w http.ResponseWriter, r *http.Request) {
// 	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
//
// 	if ID.String() == "" {
// 		http.Error(w, "DonationID required", http.StatusBadRequest)
// 		return
// 	}
//
// 	donation, err := c.donationService.GetByID(r.Context(), ID)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
//
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(donation.ToDonationResponse())
// }
//
// func (c *DonationHandler) GetDonationByOrganizationID(w http.ResponseWriter, r *http.Request) {
// 	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
//
// 	if ID.String() == "" {
// 		http.Error(w, "DonationID required", http.StatusBadRequest)
// 		return
// 	}
//
// 	donationsResult, err := c.donationService.GetByOrganizationID(r.Context(), ID)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
//
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(donationsResult)
// }
//
// func (c *DonationHandler) GetDonationByDomainID(w http.ResponseWriter, r *http.Request) {
// 	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
//
// 	if ID.String() == "" {
// 		http.Error(w, "DonationID required", http.StatusBadRequest)
// 		return
// 	}
//
// 	donationsResult, err := c.donationService.GetByDomainID(r.Context(), ID)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
//
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(donationsResult)
// }
//
// func (c *DonationHandler) GetDonationByAidTypeID(w http.ResponseWriter, r *http.Request) {
// 	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
//
// 	if ID.String() == "" {
// 		http.Error(w, "DonationID required", http.StatusBadRequest)
// 		return
// 	}
//
// 	donationsResult, err := c.donationService.GetByAidTypeID(r.Context(), ID)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
//
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(donationsResult)
// }
//
// func (c *DonationHandler) GetAll(w http.ResponseWriter, r *http.Request) {
// 	donationsResult, err := c.donationService.GetAll(r.Context())
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
//
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(donationsResult)
// }
//
// func (c *DonationHandler) DeleteDonation(w http.ResponseWriter, r *http.Request) {
// 	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
//
// 	if ID.String() == "" {
// 		http.Error(w, "DonationID required", http.StatusBadRequest)
// 		return
// 	}
//
// 	userID, ok := middleware.GetUserIDFromContext(r.Context())
//
// 	if !ok {
// 		http.Error(w, "User not found", http.StatusUnauthorized)
// 		return
// 	}
//
// 	organization, err := c.authService.GetOrganizationByID(r.Context(), userID)
// 	if err != nil || organization == nil || organization.User.Role != string(models.RoleTypeOrganization) {
// 		http.Error(w, "Not authorized", http.StatusUnauthorized)
// 		return
// 	}
//
// 	donation, err := c.donationService.GetByID(r.Context(), ID)
//
// 	if err != nil || donation.OrganizationID != organization.ID {
// 		http.Error(w, "Not authorized", http.StatusUnauthorized)
// 		return
// 	}
//
// 	err = c.donationService.Delete(r.Context(), ID)
//
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
//
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(map[string]string{
// 		"message": "deleted donation successfully",
// 	})
// }
//
// func (c *DonationHandler) GetDomains(w http.ResponseWriter, r *http.Request) {
// 	domainsResult, err := c.donationService.GetDomains(r.Context())
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
//
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(domainsResult)
// }
//
// func (c *DonationHandler) GetAidTypes(w http.ResponseWriter, r *http.Request) {
// 	aidTypesResults, err := c.donationService.GetAidTypes(r.Context())
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
//
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(aidTypesResults)
// }
//
// func (c *DonationHandler) GetDomainByID(w http.ResponseWriter, r *http.Request) {
// 	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
//
// 	if ID.String() == "" {
// 		http.Error(w, "DonationID required", http.StatusBadRequest)
// 		return
// 	}
//
// 	domain, err := c.donationService.GetDomainByID(r.Context(), ID)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
//
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(domain)
// }
//
// func (c *DonationHandler) GetAidTypeByID(w http.ResponseWriter, r *http.Request) {
// 	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
//
// 	if ID.String() == "" {
// 		http.Error(w, "DonationID required", http.StatusBadRequest)
// 		return
// 	}
//
// 	aidType, err := c.donationService.GetAidTypeByID(r.Context(), ID)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
//
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(aidType)
// }
