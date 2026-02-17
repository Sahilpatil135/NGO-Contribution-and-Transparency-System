package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"server/internal/middleware"
	"server/internal/models"
	"server/internal/repository"
	"server/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type ProofHandler struct {
	jwtService     services.JWTService
	proofService   services.ProofService
	organizationRepo repository.OrganizationRepository
}

func NewProofHandler(jwt services.JWTService, proofService services.ProofService, organizationRepo repository.OrganizationRepository) *ProofHandler {
	return &ProofHandler{
		jwtService:       jwt,
		proofService:     proofService,
		organizationRepo: organizationRepo,
	}
}

// WebSocket + Session Store(in-memory)
var (
	wsUpgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	wsClients = make(map[string][]*websocket.Conn)
	wsMutex   sync.Mutex
)

func (h *ProofHandler) RegisterRoutes(r chi.Router) {

	r.Route("/api/proof", func(r chi.Router) {
		// Legacy: no auth, returns random sessionId for QR (unchanged)
		r.Post("/session", h.CreateProofSession)
		// NGO cause session: auth required, body { "causeId": "uuid" } -> DB-backed session
		r.Group(func(protected chi.Router) {
			protected.Use(middleware.AuthMiddleware(h.jwtService))
			protected.Post("/session/cause", h.CreateCauseProofSession)
		})

		r.Post("/upload/{sessionID}", h.UploadProofImage)
	})

	r.Get("/ws/proof/{sessionID}", h.ProofWebSocket)
}

// CreateProofSession creates a legacy session (random ID, no DB). Unchanged existing process.
func (h *ProofHandler) CreateProofSession(w http.ResponseWriter, r *http.Request) {

	sessionID := uuid.New().String()
	expiresAt := time.Now().Add(15 * time.Minute)

	resp := models.ProofSessionResponse{
		SessionID: sessionID,
		ExpiresAt: expiresAt,
		QRURL:     sessionID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// CreateCauseProofSession creates a DB-backed proof session for a cause (NGO flow).
func (h *ProofHandler) CreateCauseProofSession(w http.ResponseWriter, r *http.Request) {

	var req models.CreateProofSessionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if req.CauseID == nil {
		http.Error(w, "causeId is required", http.StatusBadRequest)
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	org, err := h.organizationRepo.GetByID(r.Context(), userID)
	if err != nil || org == nil {
		http.Error(w, "Organization not found", http.StatusForbidden)
		return
	}

	session, err := h.proofService.CreateSession(r.Context(), *req.CauseID, org.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	expiresAt := time.Now().Add(15 * time.Minute)
	resp := models.ProofSessionResponse{
		SessionID: session.ID.String(),
		ExpiresAt: expiresAt,
		QRURL:     session.ID.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// WebSocket for Proof Image Upload Notifications (Laptop Listener)
func (h *ProofHandler) ProofWebSocket(w http.ResponseWriter, r *http.Request) {
	sessionID := chi.URLParam(r, "sessionID")

	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	wsMutex.Lock()
	wsClients[sessionID] = append(wsClients[sessionID], conn)
	wsMutex.Unlock()

	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			return
		}
	}
}

// Upload Image From Mobile. If sessionID is a DB session: validate metadata, hash, duplicate check, store proof, return score. Existing file save + WebSocket notify unchanged.
func (h *ProofHandler) UploadProofImage(w http.ResponseWriter, r *http.Request) {
	sessionIDStr := chi.URLParam(r, "sessionID")

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}

	latStr := r.FormValue("lat")
	lngStr := r.FormValue("lng")
	tsStr := r.FormValue("timestamp")

	if latStr != "" && lngStr != "" {
		fmt.Printf("Received location - Lat: %s, Lng: %s\n", latStr, lngStr)
	} else {
		fmt.Printf("Warning: No location received (lat: '%s', lng: '%s')\n", latStr, lngStr)
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Image required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	imageBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read image", http.StatusBadRequest)
		return
	}
	if len(imageBytes) == 0 {
		http.Error(w, "Empty image", http.StatusBadRequest)
		return
	}

	dir := "uploads/proof"
	_ = os.MkdirAll(dir, 0755)
	filename := sessionIDStr + "-" + header.Filename
	path := filepath.Join(dir, filename)
	if err := os.WriteFile(path, imageBytes, 0644); err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	relativePath := filepath.Join("proof", filename)
	lat, lng := latStr, lngStr
	captureTime := time.Now()
	if tsStr != "" {
		if t, err := time.Parse(time.RFC3339, tsStr); err == nil {
			captureTime = t
		}
	}

	// Try DB-backed session (sessionID must be UUID)
	sessionIDParsed, err := uuid.Parse(sessionIDStr)
	if err == nil {
		latF, _ := strconv.ParseFloat(latStr, 64)
		lngF, _ := strconv.ParseFloat(lngStr, 64)

		img, score, isDup, validationOK, err := h.proofService.ProcessUpload(r.Context(), sessionIDParsed, latF, lngF, captureTime, imageBytes)
		if err != nil {
			if strings.Contains(err.Error(), "session not found") {
				// Fall through to legacy response below
			} else {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		} else {
			// DB path: emit WS and return score response
			h.emitToSession(sessionIDStr, models.ProofUploadEvent{
				ImagePath: relativePath,
				Latitude:  lat,
				Longitude: lng,
				Timestamp: captureTime,
			})
			resp := models.UploadProofResponse{
				Status:       "uploaded",
				Score:        score,
				IsDuplicate:  isDup,
				ValidationOK: validationOK,
			}
			if img != nil {
				resp.Score = img.MetadataScore
			}
			if isDup {
				resp.Score = 0
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resp)
			return
		}
	}

	// Legacy: no DB session
	h.emitToSession(sessionIDStr, models.ProofUploadEvent{
		ImagePath: relativePath,
		Latitude:  lat,
		Longitude: lng,
		Timestamp: captureTime,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "uploaded",
	})
}

// WebSocket Broadcaster (Internal Function)
func (h *ProofHandler) emitToSession(sessionID string, payload interface{}) {
	wsMutex.Lock()
	defer wsMutex.Unlock()

	for _, conn := range wsClients[sessionID] {
		_ = conn.WriteJSON(payload)
	}
}
