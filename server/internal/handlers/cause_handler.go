package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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
			protected.Post("/cover/upload", c.UploadCoverImage)
			protected.Post("/products/upload", c.UploadProductImage)
			protected.Delete("/{ID}", c.DeleteCause)
		})

		r.Get("/", c.GetAllCauses)
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

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	req.CreatedAt = time.Now()
	req.CollectedAmount = 0
	req.IsActive = true

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

func (c *CauseHandler) GetAllCauses(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")

	var ctx context.Context

	if limit != "" {
		_, err := strconv.Atoi(limit)
		if err != nil {
			http.Error(w, "Limit Parameter is invalid", http.StatusBadRequest)
			return
		}
		ctx = context.WithValue(r.Context(), "limit", limit)
	} else {
		ctx = r.Context()
	}

	causesResult, err := c.causeService.GetAll(ctx)
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

	if err != nil || cause.Organization.ID != organization.ID {
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

// UploadCoverImage handles secure upload of a single campaign cover image.
// It validates image content type and stores the file under uploads/covers,
// returning a public URL path that can be saved as cover_image_url.
func (c *CauseHandler) UploadCoverImage(w http.ResponseWriter, r *http.Request) {
	// Ensure requester is an authenticated organization (same as CreateCause)
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	org, err := c.authService.GetOrganizationByID(r.Context(), userID)
	if err != nil || org == nil || org.User.Role != string(models.RoleTypeOrganization) {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Cover image required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Read a small buffer to detect content type safely
	sniff := make([]byte, 512)
	n, err := file.Read(sniff)
	if err != nil && !errors.Is(err, io.EOF) {
		http.Error(w, "Failed to read image", http.StatusBadRequest)
		return
	}
	if n == 0 {
		http.Error(w, "Empty image", http.StatusBadRequest)
		return
	}

	contentType := http.DetectContentType(sniff)
	if !strings.HasPrefix(contentType, "image/") {
		http.Error(w, "Only image uploads are allowed", http.StatusBadRequest)
		return
	}

	// Reset reader to include the sniffed bytes
	imageBytes, err := io.ReadAll(io.MultiReader(strings.NewReader(string(sniff[:n])), file))
	if err != nil {
		http.Error(w, "Failed to read full image", http.StatusBadRequest)
		return
	}

	// Derive safe extension
	exts, _ := mime.ExtensionsByType(contentType)
	ext := ""
	if len(exts) > 0 {
		ext = exts[0]
	}
	if ext == "" {
		ext = filepath.Ext(header.Filename)
	}
	if ext == "" {
		ext = ".img"
	}

	dir := "uploads/covers"
	_ = os.MkdirAll(dir, 0755)

	filename := uuid.New().String() + ext
	path := filepath.Join(dir, filename)

	if err := os.WriteFile(path, imageBytes, 0644); err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	publicPath := filepath.ToSlash(filepath.Join("uploads", "covers", filename))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"url": "/" + publicPath,
	})
}

// UploadProductImage handles secure upload of a single product image for a cause.
// It mirrors the cover image upload but stores under uploads/products and returns
// a URL that can be saved in the cause_products.image_url column.
func (c *CauseHandler) UploadProductImage(w http.ResponseWriter, r *http.Request) {
	// Reuse the same authentication as CreateCause
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	org, err := c.authService.GetOrganizationByID(r.Context(), userID)
	if err != nil || org == nil || org.User.Role != string(models.RoleTypeOrganization) {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Product image required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	sniff := make([]byte, 512)
	n, err := file.Read(sniff)
	if err != nil && !errors.Is(err, io.EOF) {
		http.Error(w, "Failed to read image", http.StatusBadRequest)
		return
	}
	if n == 0 {
		http.Error(w, "Empty image", http.StatusBadRequest)
		return
	}

	contentType := http.DetectContentType(sniff)
	if !strings.HasPrefix(contentType, "image/") {
		http.Error(w, "Only image uploads are allowed", http.StatusBadRequest)
		return
	}

	imageBytes, err := io.ReadAll(io.MultiReader(strings.NewReader(string(sniff[:n])), file))
	if err != nil {
		http.Error(w, "Failed to read full image", http.StatusBadRequest)
		return
	}

	exts, _ := mime.ExtensionsByType(contentType)
	ext := ""
	if len(exts) > 0 {
		ext = exts[0]
	}
	if ext == "" {
		ext = filepath.Ext(header.Filename)
	}
	if ext == "" {
		ext = ".img"
	}

	dir := "uploads/products"
	_ = os.MkdirAll(dir, 0755)

	filename := uuid.New().String() + ext
	path := filepath.Join(dir, filename)

	if err := os.WriteFile(path, imageBytes, 0644); err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	publicPath := filepath.ToSlash(filepath.Join("uploads", "products", filename))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"url": "/" + publicPath,
	})
}

