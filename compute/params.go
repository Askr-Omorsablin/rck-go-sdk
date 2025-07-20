package compute

import "github.com/rck/rck-go-sdk/sdkerrors"

// CustomComputeParams 定义了 CustomCompute 方法的参数。
type CustomComputeParams struct {
	Text         string
	Task         string
	OutputSchema map[string]interface{}
	CustomFields map[string]string
	Resources    []map[string]string
}

// Validate 检查参数是否有效。
func (p *CustomComputeParams) Validate() error {
	if p.Text == "" {
		return &sdkerrors.ValidationError{Field: "Text", Message: "is required"}
	}
	if p.Task == "" {
		return &sdkerrors.ValidationError{Field: "Task", Message: "is required"}
	}
	return nil
}

// AnalyzeParams 定义了 Analyze 方法的参数。
type AnalyzeParams struct {
	Text string
	Task string
	// OutputFormat 必须是预定义格式的名称 (如 "basic_analysis")。
	// 如果需要使用自定义 schema，请改用 CustomCompute 方法。
	OutputFormat string
	CustomFields map[string]string
}

// Validate 检查参数是否有效。
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

// TranslateParams 定义了 Translate 方法的参数。
type TranslateParams struct {
	Text                 string
	TargetLanguage       string
	IncludeCulturalNotes bool
}

// Validate 检查参数是否有效。
func (p *TranslateParams) Validate() error {
	if p.Text == "" {
		return &sdkerrors.ValidationError{Field: "Text", Message: "is required"}
	}
	if p.TargetLanguage == "" {
		return &sdkerrors.ValidationError{Field: "TargetLanguage", Message: "is required"}
	}
	return nil
}

// TODO: 在后续实现 Analyze, Translate 等便捷方法时，在这里添加它们的参数结构体。
// 例如:
// type AnalyzeParams struct { ... }
