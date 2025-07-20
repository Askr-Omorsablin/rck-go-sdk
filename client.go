package rck

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Askr-Omorsablin/rck-go-sdk/compute"
	"github.com/Askr-Omorsablin/rck-go-sdk/image"
	"github.com/Askr-Omorsablin/rck-go-sdk/internal/httpclient"
)

const (
	defaultBaseURL = "https://relatioe-kernel-zdibtqjzxm.us-west-1.fcapp.run"
	defaultTimeout = 60 * time.Second
)

// Client is the main client for the RCK SDK.
type Client struct {
	// Compute provides all text computation related functionality.
	Compute *compute.Kernel
	// Image provides all image generation related functionality.
	Image *image.Generator

	internalHTTPClient *httpclient.Client
}

// NewClient creates a new RCK client using an API Key.
// Client behavior can be customized by passing Options, such as setting timeout.
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

	return client, nil
}

// TestConnection tests the connection to the API.
// It verifies the API Key and network connection by executing a simple analysis request.
// Returns nil on success, or an error on failure.
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
