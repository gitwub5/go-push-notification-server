package handler

import (
	"encoding/json"

	"log"
	"net/http"

	"github.com/gitwub5/go-push-notification-server/api"
	"github.com/gitwub5/go-push-notification-server/core"
	"github.com/gitwub5/go-push-notification-server/storage/mysql"
)

// 전역 변수로 데이터베이스 인스턴스를 선언합니다.
var store *mysql.MySQLStore

// InitStore는 전역 데이터베이스 인스턴스를 설정하는 함수입니다.
func InitStore(s *mysql.MySQLStore) {
	store = s
}

// 사용자가 특정 주제를 구독하는 핸들러
func SubscribeHandler(w http.ResponseWriter, r *http.Request) {
	var subscription core.SubscriptionRequest

	// 요청 바디에서 Subscription 데이터 파싱
	err := json.NewDecoder(r.Body).Decode(&subscription)
	if err != nil {
		api.SendErrorResponse(w, "Invalid request payload", err.Error())
		return
	}

	// 구독 정보 MySQL에 추가
	newSubscriber := mysql.Subscriber{
		Token:    subscription.Token,
		Platform: subscription.Platform,
		Topic:    subscription.Topic,
	}
	err = store.AddSubscriber(newSubscriber)
	if err != nil {
		log.Printf("Failed to add subscriber to MySQL: %v", err)
		api.SendErrorResponse(w, "Failed to subscribe to topic", err.Error())
		return
	}

	log.Printf("Subscribing token: %s to topic: %s with platform: %d\n", subscription.Token, subscription.Topic, subscription.Platform)
	api.SendSuccessResponse(w, "Subscribed to topic successfully!", nil)
}

// 사용자가 특정 주제에서 구독을 취소하는 핸들러
func UnsubscribeHandler(w http.ResponseWriter, r *http.Request) {
	var subscription core.SubscriptionRequest

	// 요청 바디에서 Subscription 데이터 파싱
	err := json.NewDecoder(r.Body).Decode(&subscription)
	if err != nil {
		api.SendErrorResponse(w, "Invalid request payload", err.Error())
		return
	}

	// 구독 정보 MySQL에서 삭제
	err = store.DeleteSubscriber(subscription.Token, subscription.Topic, subscription.Platform)
	if err != nil {
		log.Printf("Failed to remove subscriber from MySQL: %v", err)
		api.SendErrorResponse(w, "Failed to unsubscribe from topic", err.Error())
		return
	}

	log.Printf("Unsubscribing token: %s from topic: %s\n", subscription.Token, subscription.Topic)
	api.SendSuccessResponse(w, "Unsubscribed from topic successfully!", nil)
}
