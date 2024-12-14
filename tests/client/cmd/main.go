package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gitwub5/go-push-client/api"
)

var baseURL string

func main() {
	// SERVER_URL 환경 변수 확인 및 기본 URL 설정
	baseURL = os.Getenv("SERVER_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080" // 기본값
	}
	fmt.Printf("Base URL: %s\n", baseURL)

	// Gin 라우터 설정
	router := gin.Default()

	// 템플릿 경로 설정
	router.LoadHTMLGlob("templates/*")

	// 엔드포인트 등록
	router.GET("/", serveHome)
	router.GET("/sendDirect", renderSendDirect)
	router.POST("/sendDirect", handleSendDirect)
	router.GET("/sendTopic", renderSendTopic)
	router.POST("/sendTopic", handleSendTopic)

	// 서버 실행
	port := ":8081" // 웹 서버 포트
	fmt.Printf("Client server is running on http://localhost%s\n", port)
	router.Run(port)
}

// 홈 화면 렌더링
func serveHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}

// Send Direct Notification 화면 렌더링
func renderSendDirect(c *gin.Context) {
	c.HTML(http.StatusOK, "send_direct.html", nil)
}

// Send Topic Notification 화면 렌더링
func renderSendTopic(c *gin.Context) {
	c.HTML(http.StatusOK, "send_topic.html", nil)
}

// 직접 전송 처리
func handleSendDirect(c *gin.Context) {
	var notification api.Notification

	// 폼 데이터 파싱
	tokens := c.PostForm("tokens")
	platform := c.PostForm("platform")
	title := c.PostForm("title")
	message := c.PostForm("message")

	// Notification 객체 생성
	notification.Tokens = parseTokens(tokens)
	notification.Platform = parsePlatform(platform)
	notification.Title = title
	notification.Message = message

	// 검증
	if len(notification.Tokens) == 0 || notification.Platform == 0 || notification.Title == "" || notification.Message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	// 알림 전송
	response, err := api.SendDirectNotification(baseURL, notification)
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
func handleSendTopic(c *gin.Context) {
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
	response, err := api.SendTopicNotification(baseURL, topic, notification)
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

// 유틸리티 함수: tokens 문자열 파싱
func parseTokens(tokens string) []string {
	return strings.Split(tokens, ",")
}

// 유틸리티 함수: platform 문자열 파싱
func parsePlatform(platform string) int {
	parsed, err := strconv.Atoi(platform)
	if err != nil {
		return 0 // 기본값 (잘못된 값)
	}
	return parsed
}
