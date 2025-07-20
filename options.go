package rck

import (
	"net/http"
	"time"
)

// config 用于 NewClient 的内部配置。
type config struct {
	apiKey     string
	baseURL    string
	timeout    time.Duration
	httpClient *http.Client
}

// Option 是一个函数，用于配置 RCKClient。
type Option func(*config)

// WithTimeout 设置 API 请求的超时时间。
func WithTimeout(timeout time.Duration) Option {
	return func(c *config) {
		c.timeout = timeout
	}
}

// WithHTTPClient 允许提供一个自定义的 http.Client。
// 这对于需要自定义传输层、代理或测试的场景非常有用。
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *config) {
		if httpClient != nil {
			c.httpClient = httpClient
		}
	}
}
