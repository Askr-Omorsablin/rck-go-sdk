package image

import (
	"context"

	"github.com/Askr-Omorsablin/rck-go-sdk/internal/apimodels"
	"github.com/Askr-Omorsablin/rck-go-sdk/internal/httpclient"
	"github.com/Askr-Omorsablin/rck-go-sdk/sdkerrors"
)

const imageEndpoint = "/sd2is/render"

// Generator provides core methods for accessing RCK image generation functionality.
type Generator struct {
	client *httpclient.Client
}

// NewGenerator creates a new image generator instance.
// This function is called by the main Client and is not typically used directly by users.
func NewGenerator(client *httpclient.Client) *Generator {
	return &Generator{client: client}
}

// Generate generates images based on the provided parameters.
func (g *Generator) Generate(ctx context.Context, params GenerateParams) (*ImageResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	// Build API request payload
	payload := apimodels.APIImageRequest{}
	payload.StartPoint.StartPoint = params.Prompt
	payload.Path.Composition = params.Composition
	payload.Path.Lighting = params.Lighting
	payload.Path.Style = params.Style

	// Send request
	rawResp, err := g.client.Post(ctx, imageEndpoint, payload)
	if err != nil {
		return nil, err
	}

	// Convert map[string]interface{} to strongly typed struct for parsing
	var apiResponse apimodels.APIImageResponse
	if err := decodeMapToStruct(rawResp, &apiResponse); err != nil {
		return nil, &sdkerrors.APIError{Message: "failed to parse API response into ImageResponse"}
	}

	// Convert internal model to public model
	response := &ImageResponse{
		Count:   apiResponse.EndPoint.Count,
		Status:  apiResponse.EndPoint.Status,
		RawData: rawResp,
		Images:  make([]ImageInfo, len(apiResponse.EndPoint.Images)),
	}

	for i, imgData := range apiResponse.EndPoint.Images {
		response.Images[i] = ImageInfo{
			URL:       imgData.URL,
			ImageData: imgData.ImageData,
			Index:     imgData.Index,
			Size:      imgData.Size,
			MimeType:  imgData.MimeType,
		}
	}

	return response, nil
}
