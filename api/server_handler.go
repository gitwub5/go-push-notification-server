package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	stats_api "github.com/fukata/golang-stats-api-handler"
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Golang 성능 지표 API
func GetGoStats(w http.ResponseWriter, r *http.Request) {
	stats_api.Handler(w, r)
}

func GetAppStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notificationStats)
}

// 서버 설정 파일 조회 API
func GetServerConfig(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("../config/config.yml")
	if err != nil {
		log.Printf("Error reading config file: %v", err)
		http.Error(w, "Could not read config file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/x-yaml")
	w.Write(data)
}
