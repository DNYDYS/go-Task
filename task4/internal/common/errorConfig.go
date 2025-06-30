package common

import "net/http"

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

// 预定义常见错误
var (
	ErrRecordNotFound = &AppError{
		Code:    http.StatusNotFound,
		Message: "记录未找到",
	}
	ErrInvalidInput = &AppError{
		Code:    http.StatusBadRequest,
		Message: "无效的输入参数",
	}
	ErrUnauthorized = &AppError{
		Code:    http.StatusUnauthorized,
		Message: "未授权",
	}
	ErrInternalServer = &AppError{
		Code:    http.StatusInternalServerError,
		Message: "服务器内部错误",
	}
)

func NewAppError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}
