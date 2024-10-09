package api

import (
	"encoding/json"
	"go-push-notification-server/core"
	"log"
	"net/http"
)

func PushNotificationHandler(w http.ResponseWriter, r *http.Request) {
	var notification core.Notification

	// 요청 바디에서 Notification 데이터 파싱
	err := json.NewDecoder(r.Body).Decode(&notification)
	if err != nil {
		sendErrorResponse(w, "Invalid request payload", err.Error())
		return
	}

	// 푸시 알림 전송 로직 (현재는 단순 로그 출력)
	log.Printf("Sending notification: %+v\n", notification)

	// 실제 알림 전송 (현재는 시뮬레이션, 나중에 FCM/APNs 연동)
	err = notification.Send()
	if err != nil {
		sendErrorResponse(w, "Failed to send notification", err.Error())
		return
	}

	sendSuccessResponse(w, "Notification sent!", nil)
}
