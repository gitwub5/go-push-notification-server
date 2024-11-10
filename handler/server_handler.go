package handler

import (
	"log"
	"net/http"
	"os"

	stats_api "github.com/fukata/golang-stats-api-handler"
	"github.com/gitwub5/go-push-notification-server/api"
)

// 전송 통계 조회 API (임시 데이터 - TODO: DB 연동)
var notificationStats = map[string]int{
	"success": 100,
	"failure": 5,
}

// 헬스체크 API
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status": "healthy",
	}
	api.SendSuccessResponse(w, "Health check successful", response)
}

// Golang 성능 지표 API
func GetGoStats(w http.ResponseWriter, r *http.Request) {
	stats_api.Handler(w, r)
}

// 애플리케이션 통계 API
func GetAppStats(w http.ResponseWriter, r *http.Request) {
	api.SendSuccessResponse(w, "Application statistics retrieved successfully", notificationStats)
}

// 서버 설정 파일 조회 API
func GetServerConfig(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("config/config.yml")
	if err != nil {
		log.Printf("Error reading config file: %v", err)
		api.SendErrorResponse(w, "Could not read config file", err.Error())
		return
	}

	// 성공 응답 반환
	api.SendSuccessResponse(w, "Server configuration retrieved successfully", string(data))
}
