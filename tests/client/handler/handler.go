package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gitwub5/go-push-client/api"
	"github.com/gitwub5/go-push-client/services"
	"github.com/gitwub5/go-push-client/utils"
)

var baseURL = os.Getenv("SERVER_URL")

// 직접 전송 처리
func HandleSendDirect(c *gin.Context) {
	if baseURL == "" {
		baseURL = "http://localhost:8080" // 기본값
	}

	var notification api.Notification

	// 폼 데이터 파싱
	tokens := c.PostForm("tokens")
	platform := c.PostForm("platform")
	title := c.PostForm("title")
	message := c.PostForm("message")

	// Notification 객체 생성
	notification.Tokens = utils.ParseTokens(tokens)
	notification.Platform = utils.ParsePlatform(platform)
	notification.Title = title
	notification.Message = message

	// 검증
	if len(notification.Tokens) == 0 || notification.Platform == 0 || notification.Title == "" || notification.Message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	// 알림 전송
	response, err := services.SendDirectNotification(baseURL, notification)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to send notification: %v", err)})
		return
	}

	// 성공 응답
	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"response": response,
	})
}

// 토픽 기반 전송 처리
func HandleSendTopic(c *gin.Context) {
	if baseURL == "" {
		baseURL = "http://localhost:8080" // 기본값
	}
	var notification api.Notification

	// 폼 데이터 파싱
	topic := c.PostForm("topic")
	title := c.PostForm("title")
	message := c.PostForm("message")

	// Notification 객체 생성
	notification.Title = title
	notification.Message = message

	// 검증
	if topic == "" || notification.Title == "" || notification.Message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	// 알림 전송
	response, err := services.SendTopicNotification(baseURL, topic, notification)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to send topic notification: %v", err)})
		return
	}

	// 성공 응답
	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"response": response,
	})
}
