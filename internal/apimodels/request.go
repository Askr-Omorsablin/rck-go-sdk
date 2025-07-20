package apimodels

// APIComputeRequest 定义了与 /compute/execute 端点完全匹配的请求结构。
// 这是 SDK 内部使用的私有结构。
type APIComputeRequest struct {
	StartPoint APIStartPoint          `json:"start_point"`
	Path       map[string]interface{} `json:"path"`
}

// APIStartPoint 定义了请求中的 start_point 部分。
type APIStartPoint struct {
	StartPoint string              `json:"startPoint"`
	Resource   []map[string]string `json:"resource,omitempty"`
}

// APIImageRequest 定义了与 /sd2is/render 端点完全匹配的请求结构。
type APIImageRequest struct {
	StartPoint struct {
		StartPoint string `json:"startPoint"`
	} `json:"start_point"`
	Path struct {
		Composition string `json:"frame_Composition"`
		Lighting    string `json:"lighting"`
		Style       string `json:"style"`
	} `json:"path"`
}
