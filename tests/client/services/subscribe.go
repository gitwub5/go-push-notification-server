package services

import (
	"fmt"
	"log"

	"github.com/gitwub5/go-push-client/api"
)

// Subscribe subscribes the client to the notification server.
func Subscribe(baseURL string, request api.SubscribeRequest) error {
	url := fmt.Sprintf("%s/subscribe", baseURL)
	response, err := api.MakePostRequest(url, request)
	if err != nil {
		return fmt.Errorf("subscription failed: %v", err)
	}
	log.Printf("Subscription successful: %s", response)
	return nil
}

// Unsubscribe unsubscribes the client from the notification server.
func Unsubscribe(baseURL string, request api.SubscribeRequest) error {
	url := fmt.Sprintf("%s/unsubscribe", baseURL)
	response, err := api.MakePostRequest(url, request)
	if err != nil {
		return fmt.Errorf("unsubscription failed: %v", err)
	}
	log.Printf("Unsubscription successful: %s", response)
	return nil
}
