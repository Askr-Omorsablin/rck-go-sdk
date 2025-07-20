package image

import "github.com/rck/rck-go-sdk/sdkerrors"

// GenerateParams 定义了 Generate 方法的参数。
type GenerateParams struct {
	Prompt      string
	Composition string
	Lighting    string
	Style       string
}

// Validate 检查参数是否有效。
func (p *GenerateParams) Validate() error {
	if p.Prompt == "" {
		return &sdkerrors.ValidationError{Field: "Prompt", Message: "is required"}
	}
	if p.Composition == "" {
		return &sdkerrors.ValidationError{Field: "Composition", Message: "is required"}
	}
	if p.Lighting == "" {
		return &sdkerrors.ValidationError{Field: "Lighting", Message: "is required"}
	}
	if p.Style == "" {
		return &sdkerrors.ValidationError{Field: "Style", Message: "is required"}
	}
	return nil
}
