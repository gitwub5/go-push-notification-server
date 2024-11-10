package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gitwub5/go-push-client/api"
)

func main() {
	// SERVER_URL 환경 변수 확인 및 기본 URL 설정
	baseURL := os.Getenv("SERVER_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080" // 기본값
	}
	fmt.Printf("Base URL: %s\n", baseURL)

	// 1. Send a test notification
	fmt.Println("1. Sending a test notification...")
	notification := api.Notification{
		Title:    "Hello",
		Message:  "This is a test notification",
		Token:    "example-token",
		Priority: "high", // 우선순위 추가
		Platform: 2,      // 플랫폼 추가 (예: 1 = iOS, 2 = Android)
	}
	notificationID, err := api.SendNotification(baseURL, notification)
	if err != nil {
		fmt.Printf("Failed to send notification: %v\n", err)
		return
	}
	time.Sleep(3 * time.Second) // 3초 대기

	// 2. Subscribe client to notifications
	fmt.Println("2. Subscribing client to notifications...")
	subRequest := api.SubscribeRequest{
		Token: "example-device-token",
		Topic: "primary notification",
	}
	api.Subscribe(baseURL, subRequest)
	time.Sleep(3 * time.Second) // 3초 대기

	// 3. Unsubscribe client from notifications
	fmt.Println("3. Unsubscribing client from notifications...")
	api.Unsubscribe(baseURL, subRequest)
	time.Sleep(3 * time.Second) // 3초 대기

	// 4. Check the status of a notification
	fmt.Println("4. Checking the status of a notification...")
	if notificationID != "" {
		api.CheckNotificationStatus(baseURL, notificationID)
	} else {
		fmt.Println("No notification ID available to check status.")
	}
	time.Sleep(3 * time.Second) // 3초 대기

	// 5. Get all notification logs
	fmt.Println("5. Getting all notification logs...")
	api.GetNotificationLogs(baseURL)
	time.Sleep(3 * time.Second) // 3초 대기

	// 6. Check server health
	fmt.Println("6. Checking server health...")
	api.CheckServerHealth(baseURL)
	time.Sleep(3 * time.Second) // 3초 대기

	// 7. Get Go runtime performance metrics
	fmt.Println("7. Getting Go runtime performance metrics...")
	api.GetGoPerformanceMetrics(baseURL)
	time.Sleep(3 * time.Second) // 3초 대기

	// 8. Get app-level notification statistics
	fmt.Println("8. Getting app-level notification statistics...")
	api.GetNotificationStats(baseURL)
	time.Sleep(3 * time.Second) // 3초 대기

	// 9. Get server configuration
	fmt.Println("9. Getting server configuration...")
	api.GetServerConfig(baseURL)

	// 출력: 모든 요청 완료 후 실행 종료 메시지
	fmt.Println("Execution has completed.")
}
