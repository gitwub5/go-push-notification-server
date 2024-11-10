package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gitwub5/go-push-notification-server/api"
	"github.com/gitwub5/go-push-notification-server/config"
	"github.com/gitwub5/go-push-notification-server/storage"
	"github.com/gitwub5/go-push-notification-server/utils"
	"github.com/gorilla/mux"
)

// TODO: 환경 변수에 개발 환경 설정하여 배포 환경일때 설정파일을 로드하도록 수정 (개발 환경에서는 localhost 사용하게 설정)

func main() {
	// 로거 초기화
	utils.InitLogger()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// MySQL 데이터베이스 초기화
	db, err := storage.NewMySQLStore("file::memory:?cache=shared")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	api.InitStore(db)

	// Redis 초기화
	redisAddr := fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port)
	redisStore := storage.NewRedisStore(redisAddr, cfg.Redis.Password, 0)
	api.InitRedisStore(redisStore)

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
	serverAddr := fmt.Sprintf(":%d", cfg.Server.Port)
	utils.InfoLogger.Printf("Starting server on %s\n", serverAddr)

	// 서버 실행 및 오류 처리
	err = http.ListenAndServe(serverAddr, r)
	if err != nil {
		utils.ErrorLogger.Fatalf("Server failed to start: %v", err)
	}
}
