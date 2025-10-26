package server

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"net/http"
	"os"
	"strconv"
	"time"

	"server/internal/config"
	"server/internal/database"
	"server/internal/handlers"
	"server/internal/repository"
	"server/internal/services"
)

type Server struct {
	port int
	db   database.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	if port == 0 {
		port = 8080
	}

	// Initialize database
	dbService := database.New()

	// Get underlying sql.DB for repositories
	sqlDB := dbService.GetDB()

	// Initialize repositories
	userRepo := repository.NewUserRepository(sqlDB)
	organizationRepo := repository.NewOrganizationRepository(sqlDB)

	causeRepo := repository.NewCauseRepository(sqlDB)

	// Initialize services
	jwtService := services.NewJWTService()
	authService := services.NewAuthService(userRepo, organizationRepo, jwtService)
	causeService := services.NewCauseService(causeRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService, jwtService)
	causeHandler := handlers.NewCauseHandler(causeService, authService, jwtService)

	// Configure OAuth
	config.ConfigureOAuth()

	server := &Server{
		port: port,
		db:   dbService,
	}

	// Routes will be registered via RegisterRoutes

	// Declare Server config
	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      server.RegisterRoutes(authHandler, causeHandler),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return httpServer
}
