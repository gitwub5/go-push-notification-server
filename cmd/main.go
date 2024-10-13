package main

import (
	"go-push-notification-server/api"
	"go-push-notification-server/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// 로거 초기화
	utils.InitLogger()

	// 새로운 gorilla/mux 라우터 생성
	r := mux.NewRouter()

	// 푸시 알림 핸들러 설정
	r.HandleFunc("/send", api.PushNotificationHandler).Methods("POST")

	// 구독 및 구독 취소 핸들러 설정
	r.HandleFunc("/subscribe", api.SubscribeHandler).Methods("POST")
	r.HandleFunc("/unsubscribe", api.UnsubscribeHandler).Methods("POST")

	// 알림 상태 핸들러 설정
	r.HandleFunc("/api/status/{notification_id}", api.GetNotificationStatus).Methods("GET")
	r.HandleFunc("/api/logs", api.GetNotificationLogs).Methods("GET")

	// 서버 상태 및 설정 핸들러 설정
	r.HandleFunc("/api/health", api.HealthCheck).Methods("GET")
	r.HandleFunc("/api/stat/go", api.GetGoStats).Methods("GET")
	r.HandleFunc("/api/stat/app", api.GetAppStats).Methods("GET")
	r.HandleFunc("/api/config", api.GetServerConfig).Methods("GET")

	// 서버 실행 로그
	utils.InfoLogger.Println("Starting server on :8080")

	// 서버 실행 및 오류 처리
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		utils.ErrorLogger.Fatalf("Server failed to start: %v", err)
	}
}
