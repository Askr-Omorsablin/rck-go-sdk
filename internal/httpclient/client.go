package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/rck/rck-go-sdk/sdkerrors"
)

// Aliases for convenience.
var (
	ErrAuthentication = sdkerrors.ErrAuthentication
	ErrAPIKeyRequired = sdkerrors.ErrAPIKeyRequired
)

// Client 是一个内部的 HTTP 客户端，用于与 RCK API 进行通信。
type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

// New 创建一个新的内部 HTTP 客户端。
func New(apiKey, baseURL string, timeout time.Duration, baseClient *http.Client) (*Client, error) {
	if apiKey == "" {
		return nil, ErrAPIKeyRequired
	}

	// 确保我们有一个非 nil 的 http.Client
	if baseClient == nil {
		baseClient = &http.Client{}
	}

	// 复制客户端以避免修改外部传入的实例
	httpClient := *baseClient
	httpClient.Timeout = timeout

	return &Client{
		apiKey:     apiKey,
		baseURL:    baseURL,
		httpClient: &httpClient,
	}, nil
}

// Post 发送一个 POST 请求到指定的端点。
func (c *Client) Post(ctx context.Context, endpoint string, payload interface{}) (map[string]interface{}, error) {
	// 序列化请求体
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request payload: %w", err)
	}

	// 创建 HTTP 请求
	url := c.baseURL + endpoint
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("topos-api-key", c.apiKey)
	req.Header.Set("User-Agent", "RCK-Go-SDK/1.0.0")

	// 发送请求
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, &sdkerrors.NetworkError{Err: err}
	}
	defer resp.Body.Close()

	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// 检查 HTTP 状态码
	if resp.StatusCode >= 400 {
		return nil, c.handleErrorResponse(resp.StatusCode, respBody)
	}

	// 解析成功的 JSON 响应
	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal successful response: %w", err)
	}

	return result, nil
}

// handleErrorResponse 将 HTTP 错误响应转换为 sdkerrors。
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
