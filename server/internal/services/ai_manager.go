package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os/exec"
	"path/filepath"
	"time"
	"log"
	// "fmt"
)

var aiURL = "http://127.0.0.1:8002"

func isAIServiceRunning() bool {

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	log.Println("Checking AI service health at", aiURL + "/health")

	resp, err := client.Get(aiURL + "/health")

	if err != nil {
		log.Println("AI service health check failed:", err)
		return false
	}

	defer resp.Body.Close()

	log.Println("AI service health status code:", resp.StatusCode)

	return resp.StatusCode == 200
}

func startAIService() error {

	cmd := exec.Command(
		"venv\\Scripts\\python.exe",
		"-m",
		"uvicorn",
		"app:app",
		"--host",
		"127.0.0.1",
		"--port",
		"8001",
	)

	// The Go API is run from the "server" directory, so the AI service
	// lives in the "./ai_service" subfolder relative to that.
	cmd.Dir = "ai_service"

	err := cmd.Start()

	if err != nil {
		return err
	}

	time.Sleep(6 * time.Second)
	// for i := 0; i < 30; i++ {

	// 	if isAIServiceRunning() {
	// 		return nil
	// 	}
	
	// 	time.Sleep(2 * time.Second)
	// }
	
	// return fmt.Errorf("AI service failed to start")

	return nil
}

func CallAIService(imagePath string, causeText string) (map[string]interface{}, error) {

	// Ensure we always send an absolute path so the Python service
	// can reliably find the image regardless of its working directory.
	if abs, err := filepath.Abs(imagePath); err == nil {
		imagePath = abs
	}
	log.Println("Calling AI service with image:", imagePath, "causeText:", causeText)

	if !isAIServiceRunning() {
		log.Println("AI service not running, starting...")
		err := startAIService()

		if err != nil {
			log.Println("Failed to start AI service:", err)
			return nil, err
		}
		log.Println("AI service started successfully")
	}

	payload := map[string]interface{}{
		"image_path": imagePath,
		"cause_text": causeText,
	}

	jsonData, _ := json.Marshal(payload)
	log.Println("AI request payload:", string(jsonData))


	resp, err := http.Post(
		aiURL+"/analyze",
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		log.Println("Error calling AI /analyze:", err)
		return nil, err
	}

	defer resp.Body.Close()
	log.Println("AI /analyze status code:", resp.StatusCode)

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)
	log.Println("AI response body:", result)

	return result, nil
}

func CallAIReceiptService(receiptPath string, claimedAmount float64) (map[string]interface{}, error) {
	// Ensure we always send an absolute path so the Python service
	// can reliably find the receipt regardless of its working directory.
	if abs, err := filepath.Abs(receiptPath); err == nil {
		receiptPath = abs
	}
	log.Println("Calling AI receipt service with receipt:", receiptPath, "claimedAmount:", claimedAmount)

	if !isAIServiceRunning() {
		log.Println("AI service not running, starting...")
		err := startAIService()
		if err != nil {
			log.Println("Failed to start AI service:", err)
			return nil, err
		}
		log.Println("AI service started successfully")
	}

	payload := map[string]interface{}{
		"receipt_path":   receiptPath,
		"claimed_amount": claimedAmount,
	}

	jsonData, _ := json.Marshal(payload)
	log.Println("AI receipt request payload:", string(jsonData))

	resp, err := http.Post(
		aiURL+"/analyze-receipt",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		log.Println("Error calling AI /analyze-receipt:", err)
		return nil, err
	}
	defer resp.Body.Close()
	log.Println("AI /analyze-receipt status code:", resp.StatusCode)

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Println("AI receipt response body:", result)

	return result, nil
}