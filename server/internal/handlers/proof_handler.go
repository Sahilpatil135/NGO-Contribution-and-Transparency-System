package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"server/internal/models"
	"server/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type ProofHandler struct {
	jwtService services.JWTService
}

func NewProofHandler(jwt services.JWTService) *ProofHandler {
	return &ProofHandler{
		jwtService: jwt,
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
		// Temporarily remove auth requirement for session creation (as requested)
		r.Post("/session", h.CreateProofSession)

		r.Post("/upload/{sessionID}", h.UploadProofImage)
	})

	r.Get("/ws/proof/{sessionID}", h.ProofWebSocket)
}

// CreateProofSession + QRcode
func (h *ProofHandler) CreateProofSession(w http.ResponseWriter, r *http.Request) {

	sessionID := uuid.New().String()
	expiresAt := time.Now().Add(15 * time.Minute)

	resp := models.ProofSessionResponse{
		SessionID: sessionID,
		ExpiresAt: expiresAt,
		QRURL:     sessionID, // frontend generates QR code --> points to http://SERVER_IP/mobile/proof/{sessionID}
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

// Upload Image From Mobile
func (h *ProofHandler) UploadProofImage(w http.ResponseWriter, r *http.Request) {
	sessionID := chi.URLParam(r, "sessionID")

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}

	lat := r.FormValue("lat")
	lng := r.FormValue("lng")

	// Log received location for debugging
	if lat != "" && lng != "" {
		fmt.Printf("Received location - Lat: %s, Lng: %s\n", lat, lng)
	} else {
		fmt.Printf("Warning: No location received (lat: '%s', lng: '%s')\n", lat, lng)
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Image required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	dir := "uploads/proof"
	_ = os.MkdirAll(dir, 0755)

	filename := sessionID + "-" + header.Filename
	path := filepath.Join(dir, filename)

	dst, _ := os.Create(path)
	defer dst.Close()

	_, _ = dst.ReadFrom(file)

	// Notify laptop via websocket
	// Send relative path for frontend access
	relativePath := filepath.Join("proof", filename)
	h.emitToSession(sessionID, models.ProofUploadEvent{
		ImagePath: relativePath,
		Latitude:  lat,
		Longitude: lng,
		Timestamp: time.Now(),
	})

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
