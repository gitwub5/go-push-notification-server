package services

import (
	"encoding/json"
	"fmt"

	"github.com/gitwub5/go-push-client/api"
)

// SendNotification sends a notification to the notification server and returns the notification ID.
func SendDirectNotification(baseURL string, notification api.Notification) (api.NotificationResponse, error) {
	url := fmt.Sprintf("%s/send", baseURL)
	fmt.Printf("Sending notification to URL: %s\n", url)

	response, err := api.MakePostRequest(url, notification)
	if err != nil {
		return api.NotificationResponse{}, fmt.Errorf("failed to send notification (MakePostRequest): %w", err)
	}

	// 서버 응답을 Go 구조체로 파싱
	var result struct {
		Status string                   `json:"status"`
		Data   api.NotificationResponse `json:"data"`
	}
	if err := json.Unmarshal(response, &result); err != nil {
		return api.NotificationResponse{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return result.Data, nil
}

// SendTopicNotification sends a topic-based notification to the server.
func SendTopicNotification(baseURL string, topic string, notification api.Notification) (api.NotificationResponse, error) {
	url := fmt.Sprintf("%s/send/%s", baseURL, topic)
	response, err := api.MakePostRequest(url, notification)
	if err != nil {
		return api.NotificationResponse{}, fmt.Errorf("failed to send topic notification: %v", err)
	}

	// 서버 응답을 Go 구조체로 파싱
	var result struct {
		Status string                   `json:"status"`
		Data   api.NotificationResponse `json:"data"`
	}
	if err := json.Unmarshal(response, &result); err != nil {
		return api.NotificationResponse{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return result.Data, nil
}
