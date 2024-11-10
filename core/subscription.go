package core

// SubscriptionRequest는 사용자의 구독 요청을 나타냅니다.
type SubscriptionRequest struct {
	Token    string `json:"token"`    // 디바이스 토큰
	Topic    string `json:"topic"`    // 구독할 주제
	Platform int    `json:"platform"` // 플랫폼 정보 (예: 1 = iOS, 2 = Android)
}
