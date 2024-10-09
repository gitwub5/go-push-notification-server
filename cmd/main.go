package main

import (
	"go-push-notification-server/api"
	"log"
	"net/http"
)

func main() {
	// 푸시 알림 핸들러 설정
	http.HandleFunc("/send", api.PushNotificationHandler)

	// 구독 및 구독 취소 핸들러 설정
	http.HandleFunc("/subscribe", api.SubscribeHandler)
	http.HandleFunc("/unsubscribe", api.UnsubscribeHandler)

	// 서버 실행
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
