package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gitwub5/go-push-notification-server/config"
	"github.com/gitwub5/go-push-notification-server/handler"
	"github.com/gitwub5/go-push-notification-server/storage/mysql"
	"github.com/gitwub5/go-push-notification-server/storage/redis"
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
	db, err := mysql.NewMySQLStore(cfg.MySQL.User, cfg.MySQL.Password, cfg.MySQL.Host, cfg.MySQL.Port, cfg.MySQL.Database)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	handler.InitStore(db)

	// Redis 초기화
	redisAddr := fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port)
	redisStore := redis.NewRedisStore(redisAddr, cfg.Redis.Password, 0)
	handler.InitRedisStore(redisStore)

	// 새로운 gorilla/mux 라우터 생성
	r := mux.NewRouter()

	// 푸시 알림 핸들러 설정
	// TODO: topic도 같이 받아서 해당 topic에 구독한 사용자에게 알림을 보내도록 수정
	// TODO: 플로우 변경 -> Redis를 큐 시스템으로 활용 (Pub/Sub)
	/*
		Redis의 Pub/Sub 또는 Stream 기능을 활용하여 비동기 알림 전송을 구현
		•	send 요청이 들어오면 Redis에 알림 이벤트를 게시
		•	별도의 **알림 처리 워커(worker)**가 Redis를 구독하고 이벤트를 처리하여 구독자들에게 알림을 보내는 방식
		1.	send 요청 수신.
		2.	Redis에 알림 이벤트 게시.
		3.	워커가 Redis에서 이벤트를 구독.
		4.	워커가 MySQL 또는 Redis에서 구독자 정보를 조회 후 알림 전송.
	*/
	r.HandleFunc("/send", handler.PushDirectNotificationHandler).Methods("POST")
	r.HandleFunc("/send/{topic}", handler.PushTopicNotificationHandler).Methods("POST")

	// 구독 및 구독 취소 핸들러 설정
	r.HandleFunc("/subscribe", handler.SubscribeHandler).Methods("POST")
	r.HandleFunc("/unsubscribe", handler.UnsubscribeHandler).Methods("POST")

	// 알림 상태 핸들러 설정
	r.HandleFunc("/api/status/{notification_id}", handler.GetNotificationStatus).Methods("GET")
	r.HandleFunc("/api/logs", handler.GetNotificationLogs).Methods("GET")

	// 서버 상태 및 설정 핸들러 설정
	r.HandleFunc("/api/health", handler.HealthCheck).Methods("GET")
	r.HandleFunc("/api/stat/go", handler.GetGoStats).Methods("GET")
	r.HandleFunc("/api/stat/app", handler.GetAppStats).Methods("GET")
	r.HandleFunc("/api/config", handler.GetServerConfig).Methods("GET")

	// 서버 실행 로그
	serverAddr := fmt.Sprintf(":%d", cfg.Server.Port)
	utils.InfoLogger.Printf("Starting server on %s\n", serverAddr)

	// 서버 실행 및 오류 처리
	err = http.ListenAndServe(serverAddr, r)
	if err != nil {
		utils.ErrorLogger.Fatalf("Server failed to start: %v", err)
	}
}
