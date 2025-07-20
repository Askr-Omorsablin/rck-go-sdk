package image

import (
	"encoding/json"
	"fmt"
)

// ImageResponse 是图像生成 API 调用的响应。
type ImageResponse struct {
	Images  []ImageInfo
	Count   int
	Status  string
	RawData map[string]interface{}
}

// ImageInfo 包含单个生成图像的信息。
type ImageInfo struct {
	URL       string
	ImageData string // Base64 编码的图像数据
	Index     int
	Size      int
	MimeType  string
}

// Success 检查请求是否成功且至少生成了一张图片。
func (r *ImageResponse) Success() bool {
	return r.Status == "success" && r.Count > 0
}

// HasData 检查图像信息是否包含可用的数据（URL 或 Base64 数据）。
func (i *ImageInfo) HasData() bool {
	return i.URL != "" || i.ImageData != ""
}

// decodeMapToStruct 是一个辅助函数，用于将 map[string]interface{} 解码到结构体。
func decodeMapToStruct(data map[string]interface{}, v interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal map: %w", err)
	}
	if err := json.Unmarshal(b, v); err != nil {
		return fmt.Errorf("failed to unmarshal into struct: %w", err)
	}
	return nil
}
