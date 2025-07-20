package sdkerrors

import (
	"errors"
	"fmt"
)

var (
	// ErrAuthentication 表示认证失败，通常是 API Key 无效。
	ErrAuthentication = errors.New("authentication failed, please check API key")
	// ErrAPIKeyRequired 表示在创建客户端时未提供 API Key。
	ErrAPIKeyRequired = errors.New("API key is required")
)

// APIError 表示从 RCK API 返回的一个错误。
type APIError struct {
	Message    string
	StatusCode int
	// ResponseData 包含了 API 响应的原始错误体。
	ResponseData map[string]interface{}
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error: %s (status code: %d)", e.Message, e.StatusCode)
}

// ValidationError 表示一个参数验证错误。
type ValidationError struct {
	Message string
	Field   string
}

func (e *ValidationError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
	}
	return fmt.Sprintf("validation error: %s", e.Message)
}

// NetworkError 表示在与 RCK API 通信时发生的网络问题。
type NetworkError struct {
	Err error
}

func (e *NetworkError) Error() string {
	return fmt.Sprintf("network error: %v", e.Err)
}

func (e *NetworkError) Unwrap() error {
	return e.Err
}
