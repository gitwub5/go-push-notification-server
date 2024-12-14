package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gitwub5/go-push-client/handler"
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

	// 템플릿 경로 설정 (환경 변수 사용)
	templatePath := os.Getenv("TEMPLATE_PATH")
	if templatePath == "" {
		templatePath = "templates/*" // 로컬 경로 기본값
	}
	router.LoadHTMLGlob(templatePath)
	// 엔드포인트 등록
	router.GET("/", serveHome)
	router.GET("/sendDirect", renderSendDirect)
	router.POST("/sendDirect", handler.HandleSendDirect)
	router.GET("/sendTopic", renderSendTopic)
	router.POST("/sendTopic", handler.HandleSendTopic)

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
