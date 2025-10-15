package config

import "os"

type RazorpayConfig struct {
	KeyID     string
	KeySecret string
}

func LoadRazorpayConfig() RazorpayConfig {
	return RazorpayConfig{
		KeyID:     os.Getenv("RAZORPAY_KEY_ID"),
		KeySecret: os.Getenv("RAZORPAY_KEY_SECRET"),
	}
}
