package compute

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rck/rck-go-sdk/internal/apimodels"
	"github.com/rck/rck-go-sdk/internal/httpclient"
	"github.com/rck/rck-go-sdk/sdkerrors"
)

const computeEndpoint = "/compute/execute"

// Kernel 提供了访问 RCK 文本计算功能的核心方法。
type Kernel struct {
	client *httpclient.Client
}

// NewKernel 创建一个新的计算内核实例。
// 此函数由主 Client 调用，用户通常不直接使用。
func NewKernel(client *httpclient.Client) *Kernel {
	return &Kernel{client: client}
}

// CustomCompute 执行一个完全自定义的计算任务。
func (k *Kernel) CustomCompute(ctx context.Context, params CustomComputeParams) (*ComputeResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	// 构建 API 请求体
	payload, err := k.buildComputePayload(params.Text, params.Task, params.OutputSchema, params.CustomFields, params.Resources)
	if err != nil {
		return nil, err
	}

	// 发送请求
	rawResp, err := k.client.Post(ctx, computeEndpoint, payload)
	if err != nil {
		return nil, err
	}

	// 提取并返回结果
	endPoint, ok := rawResp["end_point"].(map[string]interface{})
	if !ok {
		return nil, &sdkerrors.APIError{Message: "API response missing 'end_point' field"}
	}

	return &ComputeResponse{Data: endPoint}, nil
}

// Analyze 执行文本分析。
// 这是一个便捷方法，用于常见的分析任务。
func (k *Kernel) Analyze(ctx context.Context, params AnalyzeParams) (*ComputeResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	schema, err := GetPredefinedSchema(params.OutputFormat)
	if err != nil {
		// 包装错误，提供更多上下文
		return nil, &sdkerrors.ValidationError{
			Field:   "OutputFormat",
			Message: fmt.Sprintf("invalid predefined schema name: %v", err),
		}
	}

	payload, err := k.buildComputePayload(params.Text, params.Task, schema, params.CustomFields, nil)
	if err != nil {
		return nil, err
	}

	rawResp, err := k.client.Post(ctx, computeEndpoint, payload)
	if err != nil {
		return nil, err
	}

	endPoint, ok := rawResp["end_point"].(map[string]interface{})
	if !ok {
		return nil, &sdkerrors.APIError{Message: "API response missing 'end_point' field"}
	}

	return &ComputeResponse{Data: endPoint}, nil
}

// Translate 执行文本翻译。
// 这是一个便捷方法，封装了翻译任务的通用逻辑。
func (k *Kernel) Translate(ctx context.Context, params TranslateParams) (*ComputeResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	var task string
	if params.IncludeCulturalNotes {
		task = fmt.Sprintf("Translate text to %s and provide cultural background notes", params.TargetLanguage)
	} else {
		task = fmt.Sprintf("Translate text to %s", params.TargetLanguage)
	}

	schema, err := GetPredefinedSchema("translation")
	if err != nil {
		// 这个错误理论上不应该发生，因为 "translation" 是硬编码的。
		// 但为了健壮性，我们仍然处理它。
		return nil, fmt.Errorf("internal error: failed to get 'translation' schema: %w", err)
	}

	customFields := map[string]string{
		"target_language":        params.TargetLanguage,
		"include_cultural_notes": fmt.Sprintf("%v", params.IncludeCulturalNotes),
	}

	payload, err := k.buildComputePayload(params.Text, task, schema, customFields, nil)
	if err != nil {
		return nil, err
	}

	rawResp, err := k.client.Post(ctx, computeEndpoint, payload)
	if err != nil {
		return nil, err
	}

	endPoint, ok := rawResp["end_point"].(map[string]interface{})
	if !ok {
		return nil, &sdkerrors.APIError{Message: "API response missing 'end_point' field"}
	}

	return &ComputeResponse{Data: endPoint}, nil
}

// buildComputePayload 是一个辅助函数，用于创建发送到 API 的 map。
func (k *Kernel) buildComputePayload(
	text, task string,
	schema map[string]interface{},
	customFields map[string]string,
	resources []map[string]string,
) (*apimodels.APIComputeRequest, error) {

	startPoint := apimodels.APIStartPoint{
		StartPoint: text,
		Resource:   resources,
	}

	path := make(map[string]interface{})
	path["expectPath"] = task

	// 添加 schema
	if schema != nil {
		schemaBytes, err := json.Marshal(schema)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal output schema: %w", err)
		}
		path["endpointClass"] = string(schemaBytes)
	}

	// 添加自定义字段
	for key, value := range customFields {
		path[key] = value
	}

	return &apimodels.APIComputeRequest{
		StartPoint: startPoint,
		Path:       path,
	}, nil
}
