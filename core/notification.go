package core

// Notification은 푸시 알림의 데이터 구조를 정의합니다.
type Notification struct {
	Title    string `json:"title"`
	Message  string `json:"message"`
	Token    string `json:"token"`    // 디바이스 토큰
	Priority string `json:"priority"` // 알림 우선순위 (예: "high", "normal")
}

// Send는 알림을 보내는 메소드로, 실제 푸시 알림 서비스를 연결할 수 있습니다.
func (n *Notification) Send() error {
	// 예시: 실제 푸시 서비스(Firebase, APNs 등)와의 연동 로직을 구현합니다.
	// 현재는 단순히 전송을 시뮬레이션합니다.
	return nil
}
