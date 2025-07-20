package image

import (
	"context"

	"github.com/rck/rck-go-sdk/internal/apimodels"
	"github.com/rck/rck-go-sdk/internal/httpclient"
	"github.com/rck/rck-go-sdk/sdkerrors"
)

const imageEndpoint = "/sd2is/render"

// Generator 提供了访问 RCK 图像生成功能的核心方法。
type Generator struct {
	client *httpclient.Client
}

// NewGenerator 创建一个新的图像生成器实例。
// 此函数由主 Client 调用，用户通常不直接使用。
func NewGenerator(client *httpclient.Client) *Generator {
	return &Generator{client: client}
}

// Generate 根据提供的参数生成图像。
func (g *Generator) Generate(ctx context.Context, params GenerateParams) (*ImageResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	// 构建 API 请求体
	payload := apimodels.APIImageRequest{}
	payload.StartPoint.StartPoint = params.Prompt
	payload.Path.Composition = params.Composition
	payload.Path.Lighting = params.Lighting
	payload.Path.Style = params.Style

	// 发送请求
	rawResp, err := g.client.Post(ctx, imageEndpoint, payload)
	if err != nil {
		return nil, err
	}

	// 将 map[string]interface{} 转换为强类型结构体以便解析
	var apiResponse apimodels.APIImageResponse
	if err := decodeMapToStruct(rawResp, &apiResponse); err != nil {
		return nil, &sdkerrors.APIError{Message: "failed to parse API response into ImageResponse"}
	}

	// 转换内部模型到公共模型
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
