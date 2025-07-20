package rck

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rck/rck-go-sdk/compute"
	"github.com/rck/rck-go-sdk/image"
	"github.com/rck/rck-go-sdk/internal/httpclient"
)

const (
	defaultBaseURL = "https://relatioe-kernel-zdibtqjzxm.us-west-1.fcapp.run"
	defaultTimeout = 60 * time.Second
)

// Client 是 RCK SDK 的主客户端。
type Client struct {
	// Compute 提供了所有文本计算相关的功能。
	Compute *compute.Kernel
	// Image 提供了所有图像生成相关的功能。
	Image *image.Generator

	internalHTTPClient *httpclient.Client
}

// NewClient 使用 API Key 创建一个新的 RCK 客户端。
// 可以通过传入 Options 来自定义客户端的行为，例如设置超时。
func NewClient(apiKey string, opts ...Option) (*Client, error) {
	if apiKey == "" {
		return nil, httpclient.ErrAPIKeyRequired
	}

	config := &config{
		apiKey:     apiKey,
		baseURL:    defaultBaseURL,
		timeout:    defaultTimeout,
		httpClient: &http.Client{},
	}

	for _, opt := range opts {
		opt(config)
	}

	internalClient, err := httpclient.New(config.apiKey, config.baseURL, config.timeout, config.httpClient)
	if err != nil {
		return nil, err
	}

	client := &Client{
		internalHTTPClient: internalClient,
	}

	client.Compute = compute.NewKernel(internalClient)
	client.Image = image.NewGenerator(internalClient)

	log.Println("RCK Client initialized successfully")
	return client, nil
}

// TestConnection 测试与 API 的连接。
// 它通过执行一个简单的分析请求来验证 API Key 和网络连接。
// 成功时返回 nil，失败时返回一个错误。
func (c *Client) TestConnection(ctx context.Context) error {
	_, err := c.Compute.Analyze(ctx, compute.AnalyzeParams{
		Text:         "test",
		Task:         "simple analysis",
		OutputFormat: "basic_analysis",
	})
	if err != nil {
		return fmt.Errorf("connection test failed: %w", err)
	}
	return nil
}
