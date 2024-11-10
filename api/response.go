package api

import (
	"encoding/json"

	"net/http"

	"github.com/gitwub5/go-push-notification-server/core"
)

// 성공 응답을 생성하는 함수
func SendSuccessResponse(w http.ResponseWriter, message string, data interface{}) {
	response := core.APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// 에러 응답을 생성하는 함수
func SendErrorResponse(w http.ResponseWriter, message string, err string) {
	response := core.APIError{
		Status:  "error",
		Message: message,
		Error:   err,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(response)
}
