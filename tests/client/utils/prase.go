package utils

import (
	"strconv"
	"strings"
)

// 유틸리티 함수: tokens 문자열 파싱
func ParseTokens(tokens string) []string {
	return strings.Split(tokens, ",")
}

// 유틸리티 함수: platform 문자열 파싱
func ParsePlatform(platform string) int {
	parsed, err := strconv.Atoi(platform)
	if err != nil {
		return 0 // 기본값 (잘못된 값)
	}
	return parsed
}
