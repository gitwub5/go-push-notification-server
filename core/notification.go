package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Notification은 푸시 알림의 데이터 구조를 정의합니다.
type Notification struct {
	ID       string   `json:"id"`       // 알림 ID -> TODO: 사용할 것인가?
	Tokens   []string `json:"tokens"`   // 디바이스 토큰 배열
	Platform int      `json:"platform"` // 1 = iOS, 2 = Android
	Title    string   `json:"title"`
	Message  string   `json:"message"`
	Status   string   `json:"status"` // 알림 상태 (pending, sent, failed)
}

// TODO: send 하는거 service폴더 만들어서 거기로 넘겨야함
// 방식 1: 직접 토큰과 플랫폼을 지정하여 푸시 알림을 전송 (sendDirect)
func (n *Notification) SendDirect() error {
	if n.Platform == 1 {
		return n.sendToAPNs()
	} else if n.Platform == 2 {
		return n.sendToFirebase()
	}
	return fmt.Errorf("unsupported platform")
}

// 방식 2: 구독자 DB에서 해당 topic에 구독한 사용자들의 토큰을 조회하고 플랫폼에 따라 알림을 전송 (sendToSubscribers)
// func (n *Notification) SendToSubscribers(topic string) error {
// 	// TODO: DB 연결하여, Topic을 구독한 사용자 조회 (repository 패턴 사용)
// 	subscribers, err := store.GetSubscribersByTopic(topic)
// 	if err != nil {
// 		return fmt.Errorf("failed to get subscribers for topic '%s': %v", topic, err)
// 	}

// 	if len(subscribers) == 0 {
// 		log.Printf("No subscribers found for topic: %s", topic)
// 		return nil
// 	}

// 	// 구독자별로 전송
// 	for _, subscriber := range subscribers {
// 		tempNotification := *n
// 		tempNotification.Tokens = []string{subscriber.Token} // 단일 토큰 처리
// 		tempNotification.Platform = subscriber.Platform

// 		if err := tempNotification.SendDirect(); err != nil {
// 			log.Printf("Failed to send notification to %s (platform %d): %v", subscriber.Token, subscriber.Platform, err)
// 		} else {
// 			log.Printf("Notification sent successfully to %s (platform %d)", subscriber.Token, subscriber.Platform)
// 		}
// 	}

// 	return nil
// }

// TODO: APNs(Apple Push Notification Service)를 사용하여 iOS 푸시 알림 전송
func (n *Notification) sendToAPNs() error {
	for _, token := range n.Tokens {
		apnsURL := "https://api.push.apple.com/3/device/" + token
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
			log.Printf("Failed to marshal payload for token %s: %v", token, err)
			continue
		}

		req, err := http.NewRequest("POST", apnsURL, bytes.NewBuffer(payloadBytes))
		if err != nil {
			log.Printf("Failed to create request for token %s: %v", token, err)
			continue
		}

		req.Header.Set("Authorization", fmt.Sprintf("bearer %s", apnsAuthToken))
		req.Header.Set("apns-topic", "com.example.app") // APNs 주제 설정
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Failed to send notification to APNs for token %s: %v", token, err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Printf("Failed to send notification to APNs for token %s: %v", token, resp.Status)
		} else {
			log.Printf("Notification sent successfully to APNs for token %s", token)
		}
	}

	return nil
}

// TODO: FCM(Firebase Cloud Messaging)을 사용하여 Android 푸시 알림 전송
func (n *Notification) sendToFirebase() error {
	fcmURL := "https://fcm.googleapis.com/fcm/send"
	fcmServerKey := "your-firebase-server-key"

	for _, token := range n.Tokens {
		payload := map[string]interface{}{
			"to": token,
			"notification": map[string]string{
				"title": n.Title,
				"body":  n.Message,
			},
		}
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			log.Printf("Failed to marshal payload for token %s: %v", token, err)
			continue
		}

		req, err := http.NewRequest("POST", fcmURL, bytes.NewBuffer(payloadBytes))
		if err != nil {
			log.Printf("Failed to create request for token %s: %v", token, err)
			continue
		}

		req.Header.Set("Authorization", fmt.Sprintf("key=%s", fcmServerKey))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Failed to send notification to Firebase for token %s: %v", token, err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Printf("Failed to send notification to Firebase for token %s: %v", token, resp.Status)
		} else {
			log.Printf("Notification sent successfully to Firebase for token %s", token)
		}
	}

	return nil
}
