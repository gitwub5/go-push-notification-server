package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// TODO: FCM(Firebase Cloud Messaging)을 사용하여 Android 푸시 알림 전송
// TODO: APNs(Apple Push Notification Service)를 사용하여 iOS 푸시 알림 전송

// Notification은 푸시 알림의 데이터 구조를 정의합니다.
type Notification struct {
	ID       string `json:"id"` // 고유 ID (예: UUID)
	Title    string `json:"title"`
	Message  string `json:"message"`
	Token    string `json:"token"`    // 디바이스 토큰 (알람을 받을 디바이스)
	Priority string `json:"priority"` // 알림 우선순위 (예: "high", "normal")
	Platform int    `json:"platform"` // 플랫폼 (1 = iOS, 2 = Android)
	Status   string `json:"status"`   // 알림 상태 (예: "pending", "delivered", "failed")
}

// Send는 알림을 보내는 메소드로, 실제 푸시 알림 서비스를 연결할 수 있습니다.
func (n *Notification) Send() error {
	if n.Platform == 1 {
		// iOS(APNs)로 푸시 알림 전송
		err := n.sendToAPNs()
		if err != nil {
			n.Status = "failed"
			return err
		}
		n.Status = "delivered"
	} else if n.Platform == 2 {
		// Android(Firebase)로 푸시 알림 전송
		err := n.sendToFirebase()
		if err != nil {
			n.Status = "failed"
			return err
		}
		n.Status = "delivered"
	} else {
		n.Status = "failed"
		return fmt.Errorf("unsupported platform")
	}
	return nil
}

// APNs를 통한 iOS 푸시 알림 전송
func (n *Notification) sendToAPNs() error {
	apnsURL := "https://api.push.apple.com/3/device/" + n.Token
	apnsAuthToken := "your-apns-auth-token"

	payload := map[string]interface{}{
		"aps": map[string]interface{}{
			"alert": map[string]string{
				"title": n.Title,
				"body":  n.Message,
			},
			"sound": "default",
		},
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", apnsURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("bearer %s", apnsAuthToken))
	req.Header.Set("apns-topic", "com.example.app") // APNs 주제 설정
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send notification to APNs: %v", resp.Status)
	}

	return nil
}

// Firebase를 통한 Android 푸시 알림 전송
func (n *Notification) sendToFirebase() error {
	fcmURL := "https://fcm.googleapis.com/fcm/send"
	fcmServerKey := "your-firebase-server-key"

	payload := map[string]interface{}{
		"to": n.Token,
		"notification": map[string]string{
			"title": n.Title,
			"body":  n.Message,
		},
		"priority": n.Priority,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fcmURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("key=%s", fcmServerKey))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send notification to Firebase: %v", resp.Status)
	}

	return nil
}
