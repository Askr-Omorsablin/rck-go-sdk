package compute

import (
	"encoding/json"
	"fmt"
)

// ComputeResponse is the standard response for all text computation API calls.
type ComputeResponse struct {
	// Data contains the unstructured data returned by the API.
	Data map[string]interface{}
}

// Decode decodes the Data portion of the response into the struct pointer you provide.
// This is useful for converting unstructured map[string]interface{} into strongly typed structs.
//
// Usage example:
// var myResult MyTypedStruct
// err := response.Decode(&myResult)
func (r *ComputeResponse) Decode(v interface{}) error {
	if r.Data == nil {
		return fmt.Errorf("response data is nil")
	}

	// 1. Re-encode map[string]interface{} to JSON bytes.
	b, err := json.Marshal(r.Data)
	if err != nil {
		return fmt.Errorf("failed to marshal internal response data: %w", err)
	}

	// 2. Decode JSON bytes into the user-provided target struct v.
	if err := json.Unmarshal(b, v); err != nil {
		return fmt.Errorf("failed to unmarshal response data into target struct: %w", err)
	}

	return nil
}
