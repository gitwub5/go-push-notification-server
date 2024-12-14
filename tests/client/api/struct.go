package api

// Notification represents the notification structure.
type Notification struct {
	ID       string   `json:"id"`       // 알림 ID
	Tokens   []string `json:"tokens"`   // 디바이스 토큰 배열
	Platform int      `json:"platform"` // 1 = iOS, 2 = Android
	Title    string   `json:"title"`
	Message  string   `json:"message"`
	Status   string   `json:"status"` // 알림 상태 (pending, sent, failed)
}

// SubscribeRequest represents a subscription request.
type SubscribeRequest struct {
	Token string `json:"token"`
	Topic string `json:"topic"`
}

type NotificationResponse struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	Topic   string `json:"topic"`
}
