package handlers

import (
	"encoding/json"
	"net/http"
	"server/internal/services"
)

type PaymentHandler struct {
	paymentService *services.PaymentService
	keyID          string
}

func NewPaymentHandler(ps *services.PaymentService, keyID string) *PaymentHandler {
	return &PaymentHandler{paymentService: ps, keyID: keyID}
}

func (h *PaymentHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Amount int    `json:"amount"`
		Receipt string `json:"receipt"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	order, err := h.paymentService.CreateOrder(req.Amount, req.Receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := map[string]interface{}{
		"orderId":  order["id"],
		"key":      h.keyID,
		"amount":   order["amount"],
		"currency": order["currency"],
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *PaymentHandler) VerifyPayment(w http.ResponseWriter, r *http.Request) {
	var req struct {
		OrderID   string `json:"order_id"`
		PaymentID string `json:"payment_id"`
		Signature string `json:"signature"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	valid := h.paymentService.VerifySignature(req.OrderID, req.PaymentID, req.Signature)
	resp := map[string]bool{"success": valid}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
