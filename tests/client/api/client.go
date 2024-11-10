package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Notification represents the notification structure.
type Notification struct {
	ID       string `json:"id"` // 알림 ID 추가
	Title    string `json:"title"`
	Message  string `json:"message"`
	Token    string `json:"token"`
	Priority string `json:"priority"` // 우선순위 추가
	Platform int    `json:"platform"` // 플랫폼 추가
}

// SubscribeRequest represents a subscription request.
type SubscribeRequest struct {
	Token string `json:"token"`
	Topic string `json:"topic"`
}

// MakePostRequest makes a POST request to a given URL with the provided data.
func MakePostRequest(url string, data interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return body, nil
}

// Subscribe subscribes the client to the notification server.
func Subscribe(baseURL string, request SubscribeRequest) {
	url := fmt.Sprintf("%s/subscribe", baseURL)
	response, err := MakePostRequest(url, request)
	if err != nil {
		log.Fatalf("Subscription failed: %v", err)
	}
	log.Printf("Subscription successful: %s", response)
}

// Unsubscribe unsubscribes the client from the notification server.
func Unsubscribe(baseURL string, request SubscribeRequest) {
	url := fmt.Sprintf("%s/unsubscribe", baseURL)
	response, err := MakePostRequest(url, request)
	if err != nil {
		log.Fatalf("Unsubscription failed: %v", err)
	}
	log.Printf("Unsubscription successful: %s", response)
}

// SendNotification sends a notification to the notification server and returns the notification ID.
func SendNotification(baseURL string, notification Notification) (string, error) {
	url := fmt.Sprintf("%s/send", baseURL)
	response, err := MakePostRequest(url, notification)
	if err != nil {
		return "", fmt.Errorf("failed to send notification: %w", err)
	}

	// 서버 응답을 Go 구조체로 파싱
	var result struct {
		Status string `json:"status"`
		Data   struct {
			InnerData struct {
				NotificationID string `json:"notification_id"`
				Title          string `json:"title"`
				Message        string `json:"message"`
			} `json:"data"`
			Status string `json:"status"`
		} `json:"data"`
	}
	if err := json.Unmarshal([]byte(response), &result); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	if result.Data.InnerData.NotificationID == "" {
		return "", fmt.Errorf("notification ID not found in response")
	}

	return result.Data.InnerData.NotificationID, nil
}

// MakeGetRequest sends a GET request to the given URL and returns the response body.
func MakeGetRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return body, nil
}

// CheckNotificationStatus checks the status of a specific notification.
func CheckNotificationStatus(baseURL, notificationID string) {
	url := fmt.Sprintf("%s/api/status/%s", baseURL, notificationID)
	response, err := MakeGetRequest(url)
	if err != nil {
		log.Fatalf("Failed to check notification status: %v", err)
	}
	log.Printf("Notification status: %s", response)
}

// GetNotificationLogs retrieves all notification logs.
func GetNotificationLogs(baseURL string) {
	url := fmt.Sprintf("%s/api/logs", baseURL)
	response, err := MakeGetRequest(url)
	if err != nil {
		log.Fatalf("Failed to get notification logs: %v", err)
	}
	log.Printf("Notification logs: %s", response)
}

// CheckServerHealth checks if the server is running and healthy.
func CheckServerHealth(baseURL string) {
	url := fmt.Sprintf("%s/api/health", baseURL)
	response, err := MakeGetRequest(url)
	if err != nil {
		log.Fatalf("Failed to check server health: %v", err)
	}
	log.Printf("Server health: %s", response)
}

// GetGoPerformanceMetrics retrieves Go runtime performance metrics.
func GetGoPerformanceMetrics(baseURL string) {
	url := fmt.Sprintf("%s/api/stat/go", baseURL)
	response, err := MakeGetRequest(url)
	if err != nil {
		log.Fatalf("Failed to get Go performance metrics: %v", err)
	}
	log.Printf("Go performance metrics: %s", response)
}

// GetNotificationStats retrieves app-level notification statistics.
func GetNotificationStats(baseURL string) {
	url := fmt.Sprintf("%s/api/stat/app", baseURL)
	response, err := MakeGetRequest(url)
	if err != nil {
		log.Fatalf("Failed to get notification statistics: %v", err)
	}
	log.Printf("Notification statistics: %s", response)
}

// GetServerConfig retrieves the server configuration.
func GetServerConfig(baseURL string) {
	url := fmt.Sprintf("%s/api/config", baseURL)
	response, err := MakeGetRequest(url)
	if err != nil {
		log.Fatalf("Failed to get server configuration: %v", err)
	}
	log.Printf("Server configuration: %s", response)
}
