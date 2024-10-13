package utils

import (
	"io"
	"log"
	"os"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func InitLogger() {
	// 로그 파일 열기
	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	// 멀티 로거 설정 (파일과 표준 출력에 로그 기록)
	multiWriter := io.MultiWriter(file, os.Stdout)

	// 로거 생성
	InfoLogger = log.New(multiWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(multiWriter, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
