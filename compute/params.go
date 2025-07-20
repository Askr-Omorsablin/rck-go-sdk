package compute

import "github.com/Askr-Omorsablin/rck-go-sdk/sdkerrors"

// CustomComputeParams defines the parameters for the CustomCompute method.
type CustomComputeParams struct {
	Text         string
	Task         string
	OutputSchema map[string]interface{}
	CustomFields map[string]string
	Resources    []map[string]string
}

// Validate checks if the parameters are valid.
func (p *CustomComputeParams) Validate() error {
	if p.Text == "" {
		return &sdkerrors.ValidationError{Field: "Text", Message: "is required"}
	}
	if p.Task == "" {
		return &sdkerrors.ValidationError{Field: "Task", Message: "is required"}
	}
	return nil
}

// AnalyzeParams defines the parameters for the Analyze method.
type AnalyzeParams struct {
	Text string
	Task string
	// OutputFormat must be a predefined format name (e.g., "basic_analysis").
	// If you need to use a custom schema, please use the CustomCompute method instead.
	OutputFormat string
	CustomFields map[string]string
}

// Validate checks if the parameters are valid.
func (p *AnalyzeParams) Validate() error {
	if p.Text == "" {
		return &sdkerrors.ValidationError{Field: "Text", Message: "is required"}
	}
	if p.Task == "" {
		return &sdkerrors.ValidationError{Field: "Task", Message: "is required"}
	}
	if p.OutputFormat == "" {
		return &sdkerrors.ValidationError{Field: "OutputFormat", Message: "is required"}
	}
	return nil
}

// TranslateParams defines the parameters for the Translate method.
type TranslateParams struct {
	Text                 string
	TargetLanguage       string
	IncludeCulturalNotes bool
}

// Validate checks if the parameters are valid.
func (p *TranslateParams) Validate() error {
	if p.Text == "" {
		return &sdkerrors.ValidationError{Field: "Text", Message: "is required"}
	}
	if p.TargetLanguage == "" {
		return &sdkerrors.ValidationError{Field: "TargetLanguage", Message: "is required"}
	}
	return nil
}

// TODO: Add parameter structures for other convenience methods like Analyze, Translate when implementing them.
// For example:
// type AnalyzeParams struct { ... }
