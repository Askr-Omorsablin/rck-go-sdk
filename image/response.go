package image

import (
	"encoding/json"
	"fmt"
)

// ImageResponse is the response for image generation API calls.
type ImageResponse struct {
	Images  []ImageInfo
	Count   int
	Status  string
	RawData map[string]interface{}
}

// ImageInfo contains information about a single generated image.
type ImageInfo struct {
	URL       string
	ImageData string // Base64 encoded image data
	Index     int
	Size      int
	MimeType  string
}

// Success checks if the request was successful and at least one image was generated.
func (r *ImageResponse) Success() bool {
	return r.Status == "success" && r.Count > 0
}

// HasData checks if the image info contains usable data (URL or Base64 data).
func (i *ImageInfo) HasData() bool {
	return i.URL != "" || i.ImageData != ""
}

// decodeMapToStruct is a helper function to decode map[string]interface{} into a struct.
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
