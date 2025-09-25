package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"server/internal/middleware"
	"server/internal/models"
	"server/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

type AuthHandler struct {
	authService services.AuthService
	jwtService  services.JWTService
}

func NewAuthHandler(authService services.AuthService, jwtService services.JWTService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (h *AuthHandler) RegisterRoutes(r chi.Router) {
	r.Route("/api/auth", func(r chi.Router) {
		r.Post("/register", h.Register)
		r.Post("/login", h.Login)
		r.Get("/me", middleware.AuthMiddleware(h.jwtService)(http.HandlerFunc(h.GetMe)).ServeHTTP)
		r.Post("/logout", h.Logout)

		// Dynamic provider routes to work with chi and gothic
		r.Get("/{provider}", h.BeginAuth)
		r.Get("/{provider}/callback", h.CompleteAuth)
	})
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Basic validation
	if req.Name == "" || req.Email == "" || req.Password == "" {
		http.Error(w, "Name, email, and password are required", http.StatusBadRequest)
		return
	}

	if len(req.Password) < 6 {
		http.Error(w, "Password must be at least 6 characters long", http.StatusBadRequest)
		return
	}

	authResp, err := h.authService.Register(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authResp)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Basic validation
	if req.Email == "" || req.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	authResp, err := h.authService.Login(r.Context(), &req)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authResp)
}

func (h *AuthHandler) GetMe(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	user, err := h.authService.GetUserByID(r.Context(), userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	userResp := user.ToUserResponse()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userResp)
}

// BeginAuth starts the OAuth flow for any provider
func (h *AuthHandler) BeginAuth(w http.ResponseWriter, r *http.Request) {
	// Resolve provider via gothic (which we configured to read chi param)
	prov, err := gothic.GetProviderName(r)
	if err != nil || prov == "" {
		http.Error(w, "provider not specified", http.StatusBadRequest)
		return
	}
	if _, err := goth.GetProvider(prov); err != nil {
		http.Error(w, "unsupported provider", http.StatusBadRequest)
		return
	}
	gothic.BeginAuthHandler(w, r)
}

// CompleteAuth completes the OAuth flow and logs the user in
func (h *AuthHandler) CompleteAuth(w http.ResponseWriter, r *http.Request) {
	// Complete the OAuth process
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		http.Error(w, "OAuth authentication failed", http.StatusInternalServerError)
		return
	}

	provider := chi.URLParam(r, "provider")
	if provider == "" {
		provider = "google"
	}

	fmt.Println(user)

	// Derive a friendly display name if provider didn't populate Name
	displayName := user.Name
	if displayName == "" {
		if user.FirstName != "" || user.LastName != "" {
			if user.FirstName != "" && user.LastName != "" {
				displayName = user.FirstName + " " + user.LastName
			} else if user.FirstName != "" {
				displayName = user.FirstName
			} else {
				displayName = user.LastName
			}
		} else if user.NickName != "" {
			displayName = user.NickName
		} else if user.Email != "" {
			// fallback to email local-part
			for i := 0; i < len(user.Email); i++ {
				if user.Email[i] == '@' {
					displayName = user.Email[:i]
					break
				}
			}
			if displayName == "" {
				displayName = user.Email
			}
		} else {
			displayName = "User"
		}
	}

	// Create or update user in database
	authResp, err := h.authService.CreateOrUpdateOAuthUser(
		r.Context(),
		provider,
		user.UserID,
		displayName,
		user.Email,
		user.AvatarURL,
	)
	if err != nil {
		http.Error(w, "Failed to create/update user", http.StatusInternalServerError)
		return
	}

	fmt.Println(user)

	// Redirect back to frontend with token
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:5173"
	}

	redirectURL, _ := url.Parse(frontendURL)
	redirectURL.Path = "/auth/callback"
	q := redirectURL.Query()
	q.Set("token", authResp.Token)
	redirectURL.RawQuery = q.Encode()

	redirectURLString := redirectURL.String()
	fmt.Println(redirectURLString)

	http.Redirect(w, r, redirectURLString, http.StatusFound)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// In a real application, you might want to blacklist the token
	// For now, we'll just return a success message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Logged out successfully"})
}
