package image

import "github.com/Askr-Omorsablin/rck-go-sdk/sdkerrors"

// GenerateParams defines the parameters for the Generate method.
type GenerateParams struct {
	Prompt      string
	Composition string
	Lighting    string
	Style       string
}

// Validate checks if the parameters are valid.
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
