package compute

import (
	"encoding/json"
	"fmt"
)

// ComputeResponse 是所有文本计算 API 调用的标准响应。
type ComputeResponse struct {
	// Data 包含了 API 返回的非结构化数据。
	Data map[string]interface{}
}

// Decode 将响应的 Data 部分解码到你提供的结构体指针中。
// 这对于将非结构化的 map[string]interface{} 转换为强类型结构体非常有用。
//
// 使用示例:
// var myResult MyTypedStruct
// err := response.Decode(&myResult)
func (r *ComputeResponse) Decode(v interface{}) error {
	if r.Data == nil {
		return fmt.Errorf("response data is nil")
	}

	// 1. 将 map[string]interface{} 重新编码为 JSON 字节。
	b, err := json.Marshal(r.Data)
	if err != nil {
		return fmt.Errorf("failed to marshal internal response data: %w", err)
	}

	// 2. 将 JSON 字节解码到用户提供的目标结构体 v。
	if err := json.Unmarshal(b, v); err != nil {
		return fmt.Errorf("failed to unmarshal response data into target struct: %w", err)
	}

	return nil
}
