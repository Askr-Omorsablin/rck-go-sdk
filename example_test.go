package rck_test

import (
	"context"
	"testing"

	"github.com/rck/rck-go-sdk"
	"github.com/rck/rck-go-sdk/compute"
)

func TestBasicUsage(t *testing.T) {
	// 跳过测试，因为需要真实的 API key
	t.Skip("Skipping test that requires real API key")

	client, err := rck.NewClient("test-api-key")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.Compute.CustomCompute(context.Background(), compute.CustomComputeParams{
		Text: "Hello world",
		Task: "Analyze this text",
	})

	// 由于使用测试 API key，这里会失败，但至少验证了 API 结构
	if err == nil {
		t.Log("✅ SDK structure is correct")
	}
}
