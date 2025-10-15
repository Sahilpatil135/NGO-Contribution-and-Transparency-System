package server

import (
    "server/internal/config"
    "server/internal/handlers"
    "server/internal/services"

    "github.com/go-chi/chi/v5"
    "log"
)

// registerPaymentRoutes mounts payment endpoints on the provided chi router.
func (s *Server) registerPaymentRoutes(r chi.Router) {
    rzp := config.LoadRazorpayConfig()
    if rzp.KeyID == "" || rzp.KeySecret == "" {
        // Warn but still register; requests will fail fast with clear error
        log.Println("warning: RAZORPAY_KEY_ID/RAZORPAY_KEY_SECRET not set; payment endpoints may not work")
    }
    ps := services.NewPaymentService(rzp.KeyID, rzp.KeySecret)
    ph := handlers.NewPaymentHandler(ps, rzp.KeyID)

    r.Post("/api/payment/create-order", ph.CreateOrder)
    r.Post("/api/payment/verify", ph.VerifyPayment)
}
