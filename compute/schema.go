package compute

import (
	"fmt"
	"maps"
)

// PREDEFINED_SCHEMAS 存储了所有预定义的输出格式。
var PREDEFINED_SCHEMAS = map[string]map[string]interface{}{
	"basic_analysis": {
		"type": "object",
		"properties": map[string]interface{}{
			"emotion":  map[string]string{"type": "string", "description": "情感分析结果"},
			"theme":    map[string]string{"type": "string", "description": "主题分析"},
			"analysis": map[string]string{"type": "string", "description": "详细分析"},
		},
		"required": []string{"emotion", "theme", "analysis"},
	},
	"poem_creation": {
		"type": "object",
		"properties": map[string]interface{}{
			"poem":             map[string]string{"type": "string", "description": "创作的诗歌"},
			"creative_process": map[string]string{"type": "string", "description": "创作过程"},
			"style_notes":      map[string]string{"type": "string", "description": "风格注释"},
		},
		"required": []string{"poem"},
	},
	"scene_description": {
		"type": "object",
		"properties": map[string]interface{}{
			"scene_description": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"main_subjects": map[string]string{"type": "string", "description": "主要对象和空间关系"},
					"lighting":      map[string]string{"type": "string", "description": "光照条件和氛围"},
					"composition":   map[string]string{"type": "string", "description": "画面构图"},
					"style":         map[string]string{"type": "string", "description": "艺术风格"},
				},
				"required": []string{"main_subjects", "lighting", "composition", "style"},
			},
		},
		"required": []string{"scene_description"},
	},
	"translation": {
		"type": "object",
		"properties": map[string]interface{}{
			"translation":       map[string]string{"type": "string", "description": "翻译结果"},
			"original_language": map[string]string{"type": "string", "description": "源语言"},
			"target_language":   map[string]string{"type": "string", "description": "目标语言"},
			"cultural_notes":    map[string]string{"type": "string", "description": "文化背景注释"},
		},
		"required": []string{"translation"},
	},
}

// GetPredefinedSchema 获取一个预定义 Schema 的深拷贝。
func GetPredefinedSchema(schemaName string) (map[string]interface{}, error) {
	schema, ok := PREDEFINED_SCHEMAS[schemaName]
	if !ok {
		return nil, fmt.Errorf("unknown schema name: %s", schemaName)
	}
	// 返回一个副本以防止外部修改原始 map
	return maps.Clone(schema), nil
}
