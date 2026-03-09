package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
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
		Amount  int    `json:"amount"`
		Receipt string `json:"receipt"`
	}

	dump, err := httputil.DumpRequest(r, true)

	fmt.Printf("%+v\n", string(dump))

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	fmt.Printf("INVALID NO")

	order, err := h.paymentService.CreateOrder(req.Amount, req.Receipt)
	fmt.Printf("PAYMENT1: %+v\n", order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("PAYMENT: %+v\n", order)

	resp := map[string]any{
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
