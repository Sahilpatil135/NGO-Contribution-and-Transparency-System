package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/razorpay/razorpay-go"
)

type PaymentService struct {
	client    *razorpay.Client
	keySecret string
}

func NewPaymentService(keyID, keySecret string) *PaymentService {
	return &PaymentService{
		client:    razorpay.NewClient(keyID, keySecret),
		keySecret: keySecret,
	}
}

func (s *PaymentService) CreateOrder(amount int, receipt string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"amount":          amount * 100, // amount in paise
		"currency":        "INR",
		"receipt":         receipt,
		"payment_capture": 1,
	}
	return s.client.Order.Create(data, nil)
}

func (s *PaymentService) VerifySignature(orderID, paymentID, signature string) bool {
	data := orderID + "|" + paymentID
	h := hmac.New(sha256.New, []byte(s.keySecret))
	h.Write([]byte(data))
	expected := hex.EncodeToString(h.Sum(nil))
	return expected == signature
}
