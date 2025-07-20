package apimodels

// APIComputeResponse defines the structure that exactly matches the /compute/execute endpoint response.
// This is a private structure used internally by the SDK.
type APIComputeResponse struct {
	EndPoint map[string]interface{} `json:"end_point"`
}

// APIImageResponse defines the structure that exactly matches the /sd2is/render endpoint response.
type APIImageResponse struct {
	EndPoint APIImageEndPoint `json:"end_point"`
}

// APIImageEndPoint contains the detailed results of image generation.
type APIImageEndPoint struct {
	Images []APIImageInfo `json:"images"`
	Count  int            `json:"count"`
	Status string         `json:"status"`
}

// APIImageInfo contains information about a single image.
type APIImageInfo struct {
	URL       string `json:"url"`
	ImageData string `json:"imageData"` // base64 encoded
	Index     int    `json:"index"`
	Size      int    `json:"size"`
	MimeType  string `json:"mimeType"`
}
