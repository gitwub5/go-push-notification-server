package services

import (
	"fmt"
	"log"

	"github.com/gitwub5/go-push-client/api"
)

// CheckNotificationStatus checks the status of a specific notification.
func CheckNotificationStatus(baseURL, notificationID string) ([]byte, error) {
	url := fmt.Sprintf("%s/api/status/%s", baseURL, notificationID)
	return api.MakeGetRequest(url)
}

// GetNotificationLogs retrieves all notification logs.
func GetNotificationLogs(baseURL string) ([]byte, error) {
	url := fmt.Sprintf("%s/api/logs", baseURL)
	return api.MakeGetRequest(url)
}

// CheckServerHealth checks if the server is running and healthy.
func CheckServerHealth(baseURL string) ([]byte, error) {
	url := fmt.Sprintf("%s/api/health", baseURL)
	return api.MakeGetRequest(url)
}

// GetGoPerformanceMetrics retrieves Go runtime performance metrics.
func GetGoPerformanceMetrics(baseURL string) {
	url := fmt.Sprintf("%s/api/stat/go", baseURL)
	response, err := api.MakeGetRequest(url)
	if err != nil {
		log.Fatalf("Failed to get Go performance metrics: %v", err)
	}
	log.Printf("Go performance metrics: %s", response)
}

// GetNotificationStats retrieves app-level notification statistics.
func GetNotificationStats(baseURL string) {
	url := fmt.Sprintf("%s/api/stat/app", baseURL)
	response, err := api.MakeGetRequest(url)
	if err != nil {
		log.Fatalf("Failed to get notification statistics: %v", err)
	}
	log.Printf("Notification statistics: %s", response)
}

// GetServerConfig retrieves the server configuration.
func GetServerConfig(baseURL string) {
	url := fmt.Sprintf("%s/api/config", baseURL)
	response, err := api.MakeGetRequest(url)
	if err != nil {
		log.Fatalf("Failed to get server configuration: %v", err)
	}
	log.Printf("Server configuration: %s", response)
}
