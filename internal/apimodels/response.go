package apimodels

// APIComputeResponse 定义了与 /compute/execute 端点响应完全匹配的结构。
// 这是 SDK 内部使用的私有结构。
type APIComputeResponse struct {
	EndPoint map[string]interface{} `json:"end_point"`
}

// APIImageResponse 定义了与 /sd2is/render 端点响应完全匹配的结构。
type APIImageResponse struct {
	EndPoint APIImageEndPoint `json:"end_point"`
}

// APIImageEndPoint 包含图像生成的详细结果。
type APIImageEndPoint struct {
	Images []APIImageInfo `json:"images"`
	Count  int            `json:"count"`
	Status string         `json:"status"`
}

// APIImageInfo 包含单个图像的信息。
type APIImageInfo struct {
	URL       string `json:"url"`
	ImageData string `json:"imageData"` // base64 encoded
	Index     int    `json:"index"`
	Size      int    `json:"size"`
	MimeType  string `json:"mimeType"`
}
