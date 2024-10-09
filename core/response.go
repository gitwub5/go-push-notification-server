package core

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // 성공 시 반환되는 데이터
}

type APIError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"` // 에러 상세 정보
}
