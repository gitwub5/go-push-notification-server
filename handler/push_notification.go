package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gitwub5/go-push-notification-server/api"
	"github.com/gitwub5/go-push-notification-server/core"
	"github.com/google/uuid"
)

func PushNotificationHandler(w http.ResponseWriter, r *http.Request) {
	var notification core.Notification

	// 요청 바디에서 Notification 데이터 파싱
	err := json.NewDecoder(r.Body).Decode(&notification)
	if err != nil {
		api.SendErrorResponse(w, "Invalid request payload", err.Error())
		return
	}

	// 필수 값 검증
	if notification.Title == "" || notification.Message == "" {
		api.SendErrorResponse(w, "Missing required fields: title, message, or token", "")
		return
	}

	// ID 생성 (UUID 사용)
	notificationID := uuid.New().String()
	notification.ID = notificationID
	notification.Status = "delivered" // 성공 시 상태 업데이트

	// 알림을 JSON으로 직렬화
	notificationData, err := json.Marshal(notification)
	if err != nil {
		api.SendErrorResponse(w, "Failed to serialize notification", err.Error())
		return
	}

	// Redis에 알림 저장 (키: 알림 ID, 값: 알림 JSON 데이터)
	err = redisStore.Client.Set(context.Background(), notificationID, notificationData, 0).Err()
	if err != nil {
		log.Printf("Failed to save notification to Redis: %v\n", err)
		api.SendErrorResponse(w, "Failed to save notification", err.Error())
		return
	}

	// 알림을 notifications 리스트에 추가
	err = redisStore.Client.LPush(context.Background(), "notifications", notificationData).Err()
	if err != nil {
		log.Printf("Failed to push notification to Redis list: %v\n", err)
		api.SendErrorResponse(w, "Failed to log notification", err.Error())
		return
	}

	// 로그에 푸시 알림 전송 내용 출력
	log.Printf("Notification sent and saved to Redis with ID %s: %+v\n", notificationID, notification)

	// TODO: 실제 알림 전송
	// err = notification.Send()
	// if err != nil {
	// 	log.Printf("Failed to send notification: %v\n", err) // 에러 로그 기록
	// 	notification.Status = "failed"                       // 상태 업데이트
	// 	sendErrorResponse(w, "Failed to send notification", err.Error())
	// 	return
	// }

	// 성공 응답 반환
	response := map[string]interface{}{
		"status": "success",
		"data": map[string]string{
			"notification_id": notificationID,
			"title":           notification.Title,
			"message":         notification.Message,
		},
	}
	log.Printf("Response sent to client: %+v\n", response)
	api.SendSuccessResponse(w, "Notification sent!", response)
}
