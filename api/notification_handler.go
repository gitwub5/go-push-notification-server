package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gitwub5/go-push-notification-server/core"
	"github.com/gitwub5/go-push-notification-server/storage"
	"github.com/gorilla/mux"
)

// 전역 변수로 Redis 인스턴스를 선언합니다.
var redisStore *storage.RedisStore

// InitRedisStore는 전역 Redis 인스턴스를 설정하는 함수입니다.
func InitRedisStore(r *storage.RedisStore) {
	redisStore = r
}

// 알림 상태 조회 API (Redis에서 알림 상태 조회)
func GetNotificationStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	notificationIDStr := vars["notification_id"]

	// Redis에서 알림을 ID로 가져오기
	notificationJSON, err := redisStore.Client.Get(context.Background(), notificationIDStr).Result()
	if err != nil {
		http.Error(w, "Notification not found", http.StatusNotFound)
		return
	}

	// JSON 문자열을 Notification 구조체로 역직렬화
	var notification core.Notification
	err = json.Unmarshal([]byte(notificationJSON), &notification)
	if err != nil {
		http.Error(w, "Failed to parse notification data", http.StatusInternalServerError)
		return
	}

	// 응답 데이터 생성
	response := map[string]interface{}{
		"notification_id": notificationIDStr,
		"data":            notification,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// 알림 로그 조회 API (Redis에서 로그 조회)
func GetNotificationLogs(w http.ResponseWriter, r *http.Request) {
	// Redis에서 모든 알림 가져오기
	notifications, err := redisStore.GetAllNotifications(context.Background())
	if err != nil {
		http.Error(w, "Failed to retrieve notification logs", http.StatusInternalServerError)
		return
	}

	// 필요한 필드만 포함하는 구조체 슬라이스 생성
	var filteredLogs []map[string]string
	for _, notificationData := range notifications {
		var notification core.Notification
		err := json.Unmarshal([]byte(notificationData), &notification)
		if err != nil {
			log.Printf("Failed to parse notification data: %v", err)
			continue
		}

		// 필요한 필드만 포함하는 맵 생성
		filteredLog := map[string]string{
			"id":       notification.ID,
			"title":    notification.Title,
			"message":  notification.Message,
			"priority": notification.Priority,
			"status":   notification.Status,
		}
		filteredLogs = append(filteredLogs, filteredLog)
	}

	// JSON 응답 반환
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredLogs)
}
