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
	ID       string   `json:"id"`       // 알림 ID
	Tokens   []string `json:"tokens"`   // 디바이스 토큰 배열
	Platform int      `json:"platform"` // 1 = iOS, 2 = Android
	Title    string   `json:"title"`
	Message  string   `json:"message"`
	Status   string   `json:"status"` // 알림 상태 (pending, sent, failed)
}

type NotificationResponse struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	Topic   string `json:"topic"`
}

// SendNotification sends a notification to the notification server and returns the notification ID.
func SendDirectNotification(baseURL string, notification Notification) (NotificationResponse, error) {
	url := fmt.Sprintf("%s/send", baseURL)
	fmt.Printf("Sending notification to URL: %s\n", url)

	response, err := MakePostRequest(url, notification)
	if err != nil {
		return NotificationResponse{}, fmt.Errorf("failed to send notification (MakePostRequest): %w", err)
	}

	// 서버 응답을 Go 구조체로 파싱
	var result struct {
		Status string               `json:"status"`
		Data   NotificationResponse `json:"data"`
	}
	if err := json.Unmarshal(response, &result); err != nil {
		return NotificationResponse{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return result.Data, nil
}

// SendTopicNotification sends a topic-based notification to the server.
func SendTopicNotification(baseURL string, topic string, notification Notification) (NotificationResponse, error) {
	url := fmt.Sprintf("%s/send/%s", baseURL, topic)
	response, err := MakePostRequest(url, notification)
	if err != nil {
		return NotificationResponse{}, fmt.Errorf("failed to send topic notification: %v", err)
	}

	// 서버 응답을 Go 구조체로 파싱
	var result struct {
		Status string               `json:"status"`
		Data   NotificationResponse `json:"data"`
	}
	if err := json.Unmarshal(response, &result); err != nil {
		return NotificationResponse{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return result.Data, nil
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
func Subscribe(baseURL string, request SubscribeRequest) error {
	url := fmt.Sprintf("%s/subscribe", baseURL)
	response, err := MakePostRequest(url, request)
	if err != nil {
		return fmt.Errorf("subscription failed: %v", err)
	}
	log.Printf("Subscription successful: %s", response)
	return nil
}

// Unsubscribe unsubscribes the client from the notification server.
func Unsubscribe(baseURL string, request SubscribeRequest) error {
	url := fmt.Sprintf("%s/unsubscribe", baseURL)
	response, err := MakePostRequest(url, request)
	if err != nil {
		return fmt.Errorf("unsubscription failed: %v", err)
	}
	log.Printf("Unsubscription successful: %s", response)
	return nil
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
func CheckNotificationStatus(baseURL, notificationID string) ([]byte, error) {
	url := fmt.Sprintf("%s/api/status/%s", baseURL, notificationID)
	return MakeGetRequest(url)
}

// GetNotificationLogs retrieves all notification logs.
func GetNotificationLogs(baseURL string) ([]byte, error) {
	url := fmt.Sprintf("%s/api/logs", baseURL)
	return MakeGetRequest(url)
}

// CheckServerHealth checks if the server is running and healthy.
func CheckServerHealth(baseURL string) ([]byte, error) {
	url := fmt.Sprintf("%s/api/health", baseURL)
	return MakeGetRequest(url)
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
