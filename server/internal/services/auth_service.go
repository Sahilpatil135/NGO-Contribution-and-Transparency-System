package services

import (
	"context"
	"fmt"
	"time"

	"server/internal/models"
	"server/internal/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterUser(ctx context.Context, req *models.CreateUserRequest) (*models.AuthResponse, error)
	RegisterOrganization(ctx context.Context, req *models.CreateOrganizationRequest) (*models.AuthResponse, error)
	Login(ctx context.Context, req *models.LoginRequest) (*models.AuthResponse, error)
	GetUserByID(ctx context.Context, userID uuid.UUID) (*models.User, error)
	CreateOrUpdateOAuthUser(ctx context.Context, provider, providerID, name, email, avatarURL string) (*models.AuthResponse, error)
}

type authService struct {
	userRepo         repository.UserRepository
	organizationRepo repository.OrganizationRepository
	jwtService       JWTService
}

func NewAuthService(userRepo repository.UserRepository, organizationRepo repository.OrganizationRepository, jwtService JWTService) AuthService {
	return &authService{
		userRepo:         userRepo,
		organizationRepo: organizationRepo,
		jwtService:       jwtService,
	}
}

func (a *authService) RegisterUser(ctx context.Context, req *models.CreateUserRequest) (*models.AuthResponse, error) {
	// Check if user already exists
	existingUser, err := a.userRepo.GetByEmail(ctx, req.Email)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("user with email %s already exists", req.Email)
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Create user
	user := &models.User{
		ID:           uuid.New(),
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: stringPtr(string(hashedPassword)),
		Provider:     "email",
		IsActive:     true,
		IsVerified:   false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Save user to database
	if err := a.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Generate JWT token
	token, err := a.jwtService.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &models.AuthResponse{
		User:  *user,
		Token: token,
	}, nil
}

func (a *authService) RegisterOrganization(ctx context.Context, req *models.CreateOrganizationRequest) (*models.AuthResponse, error) {
	// Check if user already exists
	existingUser, err := a.userRepo.GetByEmail(ctx, req.Email)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("user with email %s already exists", req.Email)
	}

	// Check if organization already exists
	existingOrganization, err := a.organizationRepo.GetByEmail(ctx, req.Email)
	if err == nil && existingOrganization != nil {
		return nil, fmt.Errorf("organization with email %s already exists", req.Email)
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Create user
	user := &models.User{
		ID:           uuid.New(),
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: stringPtr(string(hashedPassword)),
		Provider:     "email",
		IsActive:     true,
		IsVerified:   false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Role:         "organization",
	}

	fmt.Println(user.Role)

	// Create user
	organization := &models.Organization{
		ID:                 uuid.New(),
		UserID:             user.ID,
		User:               user,
		OrganizationName:   req.OrganizationName,
		OrganizationType:   req.OrganizationType,
		RegistrationNumber: req.RegistrationNumber,
		About:              req.About,
		WebsiteUrl:         req.WebsiteUrl,
		Address:            req.Address,
		IsApproved:         false,
	}

	// Save user to database
	if err := a.organizationRepo.Create(ctx, organization); err != nil {
		return nil, fmt.Errorf("failed to create organization: %w", err)
	}

	// Generate JWT token
	token, err := a.jwtService.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &models.AuthResponse{
		User:  *user,
		Token: token,
	}, nil
}

func (a *authService) Login(ctx context.Context, req *models.LoginRequest) (*models.AuthResponse, error) {
	// Get user by email
	user, err := a.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("Email does not exist")
	}

	// Check if user has password (email provider)
	if user.PasswordHash == nil {
		return nil, fmt.Errorf("Invalid credentials")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(*user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, fmt.Errorf("Invalid credentials")
	}

	// Generate JWT token
	token, err := a.jwtService.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate token: %w", err)
	}

	return &models.AuthResponse{
		User:  *user,
		Token: token,
	}, nil
}

func (a *authService) GetUserByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	return a.userRepo.GetByID(ctx, userID)
}

func (a *authService) CreateOrUpdateOAuthUser(ctx context.Context, provider, providerID, name, email, avatarURL string) (*models.AuthResponse, error) {
	// Check if user exists by provider ID
	user, err := a.userRepo.GetByProviderID(ctx, provider, providerID)
	if err == nil && user != nil {
		// User exists, update if needed
		updated := false
		if user.Name != name {
			user.Name = name
			updated = true
		}
		if user.Email != email {
			user.Email = email
			updated = true
		}
		if (user.AvatarURL == nil && avatarURL != "") || (user.AvatarURL != nil && *user.AvatarURL != avatarURL) {
			user.AvatarURL = stringPtr(avatarURL)
			updated = true
		}

		if updated {
			if err := a.userRepo.Update(ctx, user); err != nil {
				return nil, fmt.Errorf("failed to update user: %w", err)
			}
		}

		// Generate JWT token
		token, err := a.jwtService.GenerateToken(user.ID, user.Email)
		if err != nil {
			return nil, fmt.Errorf("failed to generate token: %w", err)
		}

		return &models.AuthResponse{
			User:  *user,
			Token: token,
		}, nil
	}

	// Check if user exists by email (different provider)
	existingUser, err := a.userRepo.GetByEmail(ctx, email)
	if err == nil && existingUser != nil {
		// User exists with different provider, update to include OAuth info
		existingUser.Provider = provider
		existingUser.ProviderID = stringPtr(providerID)
		existingUser.AvatarURL = stringPtr(avatarURL)
		existingUser.UpdatedAt = time.Now()

		if err := a.userRepo.Update(ctx, existingUser); err != nil {
			return nil, fmt.Errorf("failed to update user: %w", err)
		}

		// Generate JWT token
		token, err := a.jwtService.GenerateToken(existingUser.ID, existingUser.Email)
		if err != nil {
			return nil, fmt.Errorf("failed to generate token: %w", err)
		}

		return &models.AuthResponse{
			User:  *existingUser,
			Token: token,
		}, nil
	}

	// Create new user
	user = &models.User{
		ID:         uuid.New(),
		Name:       name,
		Email:      email,
		Provider:   provider,
		ProviderID: stringPtr(providerID),
		AvatarURL:  stringPtr(avatarURL),
		IsActive:   true,
		IsVerified: true, // OAuth users are considered verified
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Role:       "user",
	}

	// Save user to database
	if err := a.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Generate JWT token
	token, err := a.jwtService.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &models.AuthResponse{
		User:  *user,
		Token: token,
	}, nil
}

// Helper function to create string pointer
func stringPtr(s string) *string {
	return &s
}
