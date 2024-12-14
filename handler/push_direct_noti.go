package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gitwub5/go-push-notification-server/api"
	"github.com/gitwub5/go-push-notification-server/core"
)

func PushDirectNotificationHandler(w http.ResponseWriter, r *http.Request) {
	var notification core.Notification

	// 요청 바디에서 Notification 데이터 파싱
	err := json.NewDecoder(r.Body).Decode(&notification)
	if err != nil {
		api.SendErrorResponse(w, "Invalid request payload", err.Error())
		return
	}

	// 필수 값 검증
	if len(notification.Tokens) == 0 || notification.Platform == 0 {
		api.SendErrorResponse(w, "Missing required fields: tokens and platform are mandatory", "")
		return
	}

	if notification.Title == "" || notification.Message == "" {
		api.SendErrorResponse(w, "Missing required fields: title and message are mandatory", "")
		return
	}

	// TODO: 실제 알림 전송
	// err = notification.SendDirect()
	// if err != nil {
	// 	log.Printf("Failed to send notification: %v\n", err)
	// 	api.SendErrorResponse(w, "Failed to send notification", err.Error())
	// 	return
	// }

	// 성공 응답 반환
	response := map[string]interface{}{
		"title":   notification.Title,
		"message": notification.Message,
	}
	log.Printf("Response sent to client: %+v\n", response)
	api.SendSuccessResponse(w, "Notification sent successfully!", response)
}
