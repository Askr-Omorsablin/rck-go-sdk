package compute

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Askr-Omorsablin/rck-go-sdk/internal/apimodels"
	"github.com/Askr-Omorsablin/rck-go-sdk/internal/httpclient"
	"github.com/Askr-Omorsablin/rck-go-sdk/sdkerrors"
)

const computeEndpoint = "/compute/execute"

// Kernel provides core methods for accessing RCK text computation functionality.
type Kernel struct {
	client *httpclient.Client
}

// NewKernel creates a new compute kernel instance.
// This function is called by the main Client and is not typically used directly by users.
func NewKernel(client *httpclient.Client) *Kernel {
	return &Kernel{client: client}
}

// CustomCompute executes a fully customized computation task.
func (k *Kernel) CustomCompute(ctx context.Context, params CustomComputeParams) (*ComputeResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	// Build API request payload
	payload, err := k.buildComputePayload(params.Text, params.Task, params.OutputSchema, params.CustomFields, params.Resources)
	if err != nil {
		return nil, err
	}

	// Send request
	rawResp, err := k.client.Post(ctx, computeEndpoint, payload)
	if err != nil {
		return nil, err
	}

	// Extract and return result
	endPoint, ok := rawResp["end_point"].(map[string]interface{})
	if !ok {
		return nil, &sdkerrors.APIError{Message: "API response missing 'end_point' field"}
	}

	return &ComputeResponse{Data: endPoint}, nil
}

// Analyze performs text analysis.
// This is a convenience method for common analysis tasks.
func (k *Kernel) Analyze(ctx context.Context, params AnalyzeParams) (*ComputeResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	schema, err := GetPredefinedSchema(params.OutputFormat)
	if err != nil {
		// Wrap error with more context
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

// Translate performs text translation.
// This is a convenience method that encapsulates common translation task logic.
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
		// This error should theoretically not occur since "translation" is hardcoded.
		// But for robustness, we still handle it.
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

// buildComputePayload is a helper function to create the map sent to the API.
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

	// Add schema
	if schema != nil {
		schemaBytes, err := json.Marshal(schema)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal output schema: %w", err)
		}
		path["endpointClass"] = string(schemaBytes)
	}

	// Add custom fields
	for key, value := range customFields {
		path[key] = value
	}

	return &apimodels.APIComputeRequest{
		StartPoint: startPoint,
		Path:       path,
	}, nil
}
