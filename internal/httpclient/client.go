package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Askr-Omorsablin/rck-go-sdk/sdkerrors"
)

// Aliases for convenience.
var (
	ErrAuthentication = sdkerrors.ErrAuthentication
	ErrAPIKeyRequired = sdkerrors.ErrAPIKeyRequired
)

// Client is an internal HTTP client for communicating with the RCK API.
type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

// New creates a new internal HTTP client.
func New(apiKey, baseURL string, timeout time.Duration, baseClient *http.Client) (*Client, error) {
	if apiKey == "" {
		return nil, ErrAPIKeyRequired
	}

	// Ensure we have a non-nil http.Client
	if baseClient == nil {
		baseClient = &http.Client{}
	}

	// Copy the client to avoid modifying the externally passed instance
	httpClient := *baseClient
	httpClient.Timeout = timeout

	return &Client{
		apiKey:     apiKey,
		baseURL:    baseURL,
		httpClient: &httpClient,
	}, nil
}

// Post sends a POST request to the specified endpoint.
func (c *Client) Post(ctx context.Context, endpoint string, payload interface{}) (map[string]interface{}, error) {
	// Serialize request body
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request payload: %w", err)
	}

	// Create HTTP request
	url := c.baseURL + endpoint
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("topos-api-key", c.apiKey)
	req.Header.Set("User-Agent", "RCK-Go-SDK/1.0.0")

	// Send request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, &sdkerrors.NetworkError{Err: err}
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check HTTP status code
	if resp.StatusCode >= 400 {
		return nil, c.handleErrorResponse(resp.StatusCode, respBody)
	}

	// Parse successful JSON response
	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal successful response: %w", err)
	}

	return result, nil
}

// handleErrorResponse converts HTTP error responses to sdkerrors.
func (c *Client) handleErrorResponse(statusCode int, body []byte) error {
	if statusCode == http.StatusUnauthorized || statusCode == http.StatusForbidden {
		return ErrAuthentication
	}

	apiErr := &sdkerrors.APIError{
		StatusCode: statusCode,
	}

	var errorData map[string]interface{}
	if err := json.Unmarshal(body, &errorData); err == nil {
		apiErr.ResponseData = errorData
		if msg, ok := errorData["error"].(string); ok {
			apiErr.Message = msg
		} else {
			apiErr.Message = "API request failed"
		}
	} else {
		apiErr.Message = "API request failed with unparseable error response"
	}

	return apiErr
}
