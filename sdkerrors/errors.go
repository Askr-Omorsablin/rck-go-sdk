package sdkerrors

import (
	"errors"
	"fmt"
)

var (
	// ErrAuthentication indicates authentication failure, usually due to invalid API Key.
	ErrAuthentication = errors.New("authentication failed, please check API key")
	// ErrAPIKeyRequired indicates that no API Key was provided when creating the client.
	ErrAPIKeyRequired = errors.New("API key is required")
)

// APIError represents an error returned from the RCK API.
type APIError struct {
	Message    string
	StatusCode int
	// ResponseData contains the raw error body from the API response.
	ResponseData map[string]interface{}
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error: %s (status code: %d)", e.Message, e.StatusCode)
}

// ValidationError represents a parameter validation error.
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

// NetworkError represents a network issue when communicating with the RCK API.
type NetworkError struct {
	Err error
}

func (e *NetworkError) Error() string {
	return fmt.Sprintf("network error: %v", e.Err)
}

func (e *NetworkError) Unwrap() error {
	return e.Err
}
