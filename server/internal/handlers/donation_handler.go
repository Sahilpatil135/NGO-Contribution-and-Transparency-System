package handlers

import (
	"encoding/json"
	"net/http"

	"server/internal/middleware"
	"server/internal/models"
	"server/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
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
			protected.Get("/user/me", c.GetDonationByUserID)
			// protected.Delete("/{ID}", c.DeleteDonation)
		})

		r.Get("/{ID}", c.GetDonationByID)
		r.Get("/cause/{ID}", c.GetDonationByCauseID)
		r.Get("/payment/{ID}", c.GetDonationByPaymentID)

		r.Get("/chain/{ID}", c.GetDonationFromChainByID)
		r.Get("/chain/cause/{ID}", c.GetDonationFromChainByCauseID)
	})
}

func GetUserFromContext(w http.ResponseWriter, r *http.Request, c *DonationHandler) (*models.User, error) {
	userID, ok := middleware.GetUserIDFromContext(r.Context())

	if !ok {
		http.Error(w, "User not found", http.StatusUnauthorized)
	}

	user, err := c.authService.GetUserByID(r.Context(), userID)
	if err != nil {
		http.Error(w, "User not available", http.StatusBadRequest)
		return nil, err
	}
	return user, nil
}

func (c *DonationHandler) CreateDonation(w http.ResponseWriter, r *http.Request) {
	var req models.CreateDonationRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := GetUserFromContext(w, r, c)

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

func GetIDFromURL(w http.ResponseWriter, r *http.Request) (*uuid.UUID, error) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	if ID.String() == "" {
		http.Error(w, "DonationID required", http.StatusBadRequest)
		return nil, err
	}

	return &ID, err
}

func (c *DonationHandler) GetDonationByID(w http.ResponseWriter, r *http.Request) {
	ID, err := GetIDFromURL(w, r)

	donation, err := c.donationService.GetByID(r.Context(), *ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(donation.ToDonationResponse())
}

func (c *DonationHandler) GetDonationByCauseID(w http.ResponseWriter, r *http.Request) {
	ID, err := GetIDFromURL(w, r)

	donationsResult, err := c.donationService.GetByCauseID(r.Context(), *ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var donationsResponse []*models.CreateDonationResponse
	for _, donation := range donationsResult {
		donationsResponse = append(donationsResponse, donation.ToDonationResponse())
	}

	json.NewEncoder(w).Encode(donationsResponse)
}

func (c *DonationHandler) GetDonationByPaymentID(w http.ResponseWriter, r *http.Request) {
	ID, err := GetIDFromURL(w, r)

	donation, err := c.donationService.GetByPaymentID(r.Context(), *ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(donation.ToDonationResponse())
}

func (c *DonationHandler) GetDonationByUserID(w http.ResponseWriter, r *http.Request) {
	user, err := GetUserFromContext(w, r, c)

	donationsResult, err := c.donationService.GetByUserID(r.Context(), user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var donationsResponse []*models.CreateDonationResponse
	for _, donation := range donationsResult {
		donationsResponse = append(donationsResponse, donation.ToDonationResponse())
	}

	json.NewEncoder(w).Encode(donationsResponse)
}

func (c *DonationHandler) GetDonationFromChainByID(w http.ResponseWriter, r *http.Request) {
	ID, err := GetIDFromURL(w, r)

	donation, err := c.donationService.GetFromChainByID(r.Context(), *ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.ToDonationLedgerResponse(donation))
}

func (c *DonationHandler) GetDonationFromChainByCauseID(w http.ResponseWriter, r *http.Request) {
	ID, err := GetIDFromURL(w, r)

	donations, err := c.donationService.GetFromChainByCauseID(r.Context(), *ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var donationLedgerResponses []*models.DonationLedgerResponse
	for _, donation := range donations {
		donationLedgerResponses = append(
			donationLedgerResponses,
			models.ToDonationLedgerResponse(donation),
		)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(donationLedgerResponses)
}
