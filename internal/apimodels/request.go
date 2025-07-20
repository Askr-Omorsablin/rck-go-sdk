package apimodels

// APIComputeRequest defines the request structure that exactly matches the /compute/execute endpoint.
// This is a private structure used internally by the SDK.
type APIComputeRequest struct {
	StartPoint APIStartPoint          `json:"start_point"`
	Path       map[string]interface{} `json:"path"`
}

// APIStartPoint defines the start_point portion of the request.
type APIStartPoint struct {
	StartPoint string              `json:"startPoint"`
	Resource   []map[string]string `json:"resource,omitempty"`
}

// APIImageRequest defines the request structure that exactly matches the /sd2is/render endpoint.
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
