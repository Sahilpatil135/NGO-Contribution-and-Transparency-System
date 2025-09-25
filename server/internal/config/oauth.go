package config

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func ConfigureOAuth() {
	// Session store for gothic (required)
	sessionSecret := os.Getenv("GOTH_SESSION_SECRET")
	if sessionSecret == "" {
		// Fallback to JWT secret in dev if not provided
		sessionSecret = os.Getenv("JWT_SECRET")
		if sessionSecret == "" {
			sessionSecret = "dev-session-secret"
		}
	}

	store := sessions.NewCookieStore([]byte(sessionSecret))
	store.MaxAge(int((24 * time.Hour).Seconds()))
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	// In production set Secure=true when serving over HTTPS
	gothic.Store = store

	// Make gothic provider resolution work with chi
	gothic.GetProviderName = func(r *http.Request) (string, error) {
		// Prefer chi URL param: /api/auth/{provider}
		if p := chi.URLParam(r, "provider"); p != "" {
			return p, nil
		}
		// Fallback to query string ?provider=
		if p := r.URL.Query().Get("provider"); p != "" {
			return p, nil
		}
		return "", errors.New("provider not specified")
	}

	// Configure Google OAuth
	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	if googleClientID != "" && googleClientSecret != "" {
		goth.UseProviders(
			// Callback path should include provider for chi routing
			google.New(googleClientID, googleClientSecret, baseURL+"/api/auth/google/callback"),
		)
	}
}
