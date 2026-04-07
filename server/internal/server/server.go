package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"server/internal/blockchain"
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

	// Initialize Blockchain ETH Client
	blockchainClient, err := blockchain.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(sqlDB)
	organizationRepo := repository.NewOrganizationRepository(sqlDB)
	causeRepo := repository.NewCauseRepository(sqlDB)
	causeVoteRepo := repository.NewCauseVoteRepository(sqlDB)
	causeReviewRepo := repository.NewCauseReviewRepository(sqlDB)
	donationRepo := repository.NewDonationRepository(sqlDB)
	proofSessionRepo := repository.NewProofSessionRepository(sqlDB)
	proofImageRepo := repository.NewProofImageRepository(sqlDB)
	disbursementRepo := repository.NewDisbursementRepository(sqlDB)

	// Initialize services
	jwtService := services.NewJWTService()
	authService := services.NewAuthService(userRepo, organizationRepo, jwtService)
	causeService := services.NewCauseService(causeRepo)
	causeVoteService := services.NewCauseVoteService(causeVoteRepo)
	causeReviewService := services.NewCauseReviewService(causeReviewRepo)
	proofService := services.NewProofService(proofSessionRepo, proofImageRepo, causeRepo)

	// Initialize blockchain services
	chainService, err := blockchain.NewDonationChainService(
		blockchainClient,
		os.Getenv("CONTRACT_ADDRESS_1"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize milestone tracker service
	trackerService, err := blockchain.NewMilestoneTrackerService(
		blockchainClient,
		os.Getenv("MILESTONE_TRACKER_ADDRESS"),
	)
	if err != nil {
		log.Printf("Warning: Failed to initialize milestone tracker service: %v", err)
		// Continue without tracker if not configured
	}

	donationService := services.NewDonationService(donationRepo, *chainService, trackerService, causeRepo)

	// Start milestone tracker event listener if tracker service is available
	if trackerService != nil {
		eventListener, err := services.NewEscrowEventListener(
			blockchainClient,
			os.Getenv("MILESTONE_TRACKER_ADDRESS"),
			disbursementRepo,
			organizationRepo,
			causeRepo,
		)
		if err != nil {
			log.Printf("Warning: Failed to initialize event listener: %v", err)
		} else {
			// Start listening for events in a goroutine
			go func() {
				if err := eventListener.Start(context.Background()); err != nil {
					log.Printf("Event listener error: %v", err)
				}
			}()
			log.Println("Milestone tracker event listener started")
		}
	}

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService, jwtService)
	ipfsService := services.NewIPFSService()
	causeHandler := handlers.NewCauseHandler(causeService, authService, jwtService, causeVoteService, causeReviewService, ipfsService)
	donationHandler := handlers.NewDonationHandler(donationService, authService, jwtService)
	proofHandler := handlers.NewProofHandler(jwtService, proofService, organizationRepo, causeRepo)
	disbursementHandler := handlers.NewDisbursementHandler(disbursementRepo, organizationRepo, jwtService)

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
		Handler:      server.RegisterRoutes(authHandler, causeHandler, donationHandler, proofHandler, disbursementHandler),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return httpServer
}
