package api

import (
	"encoding/json"
	"fmt"
	"go-push-notification-server/storage"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var redisStore = storage.NewRedisStore("localhost:6379", "", 0) // Redis 연결
var mysqlStore, _ = storage.NewMySQLStore("your-dsn")           // MySQL 연결

// 알림 상태 조회 API (MySQL에서 알림 상태 조회)
func GetNotificationStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	notificationIDStr := vars["notification_id"]

	// ID를 문자열에서 uint로 변환
	notificationID, err := strconv.ParseUint(notificationIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid notification ID", http.StatusBadRequest)
		return
	}

	// MySQL에서 알림 상태 조회
	notification, err := mysqlStore.GetNotificationByID(uint(notificationID))
	if err != nil {
		http.Error(w, "Notification not found", http.StatusNotFound)
		return
	}

	// 응답 데이터 생성
	response := map[string]string{
		"notification_id": fmt.Sprintf("%d", notification.ID),
		"status":          notification.Status,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// 알림 로그 조회 API (MySQL에서 로그 조회)
func GetNotificationLogs(w http.ResponseWriter, r *http.Request) {
	// MySQL에서 모든 알림 로그 가져오기
	notifications, err := mysqlStore.GetAllNotifications()
	if err != nil {
		http.Error(w, "Failed to retrieve notification logs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notifications)
}
