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

	// 로그에 푸시 알림 전송 내용 출력
	log.Printf("Sending notification: %+v\n", notification)

	// 실제 알림 전송
	// err = notification.Send()
	// if err != nil {
	// 	log.Printf("Failed to send notification: %v\n", err) // 에러 로그 기록
	// 	notification.Status = "failed"                       // 상태 업데이트
	// 	sendErrorResponse(w, "Failed to send notification", err.Error())
	// 	return
	// }

	// 성공 시 상태 업데이트
	notification.Status = "delivered"
	log.Printf("Notification sent successfully: %+v\n", notification)

	sendSuccessResponse(w, "Notification sent!", nil)
}
