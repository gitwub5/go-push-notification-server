package api

import (
	"encoding/json"
	"go-push-notification-server/core"
	"log"
	"net/http"
)

// 사용자가 특정 주제를 구독하는 핸들러
func SubscribeHandler(w http.ResponseWriter, r *http.Request) {
	var subscription core.SubscriptionRequest

	// 요청 바디에서 Subscription 데이터 파싱
	err := json.NewDecoder(r.Body).Decode(&subscription)
	if err != nil {
		sendErrorResponse(w, "Invalid request payload", err.Error())
		return
	}

	// 구독 로직 (주제 및 토큰을 저장)
	// 여기서는 단순히 로그로 출력하지만, 실제로는 Redis나 데이터베이스에 저장
	log.Printf("Subscribing token: %s to topic: %s\n", subscription.Token, subscription.Topic)

	// 구독 성공 응답
	sendSuccessResponse(w, "Subscribed to topic successfully!", nil)
}

// 사용자가 특정 주제에서 구독을 취소하는 핸들러
func UnsubscribeHandler(w http.ResponseWriter, r *http.Request) {
	var subscription core.SubscriptionRequest

	// 요청 바디에서 Subscription 데이터 파싱
	err := json.NewDecoder(r.Body).Decode(&subscription)
	if err != nil {
		sendErrorResponse(w, "Invalid request payload", err.Error())
		return
	}

	// 구독 취소 로직 (주제 및 토큰을 삭제)
	// 여기서는 단순히 로그로 출력하지만, 실제로는 Redis나 데이터베이스에서 제거
	log.Printf("Unsubscribing token: %s from topic: %s\n", subscription.Token, subscription.Topic)

	// 구독 취소 성공 응답
	sendSuccessResponse(w, "Unsubscribed from topic successfully!", nil)
}
