package compute

import (
	"fmt"
	"maps"
)

// PREDEFINED_SCHEMAS stores all predefined output formats.
var PREDEFINED_SCHEMAS = map[string]map[string]interface{}{
	"basic_analysis": {
		"type": "object",
		"properties": map[string]interface{}{
			"emotion":  map[string]string{"type": "string", "description": "Emotion analysis result"},
			"theme":    map[string]string{"type": "string", "description": "Theme analysis"},
			"analysis": map[string]string{"type": "string", "description": "Detailed analysis"},
		},
		"required": []string{"emotion", "theme", "analysis"},
	},
	"poem_creation": {
		"type": "object",
		"properties": map[string]interface{}{
			"poem":             map[string]string{"type": "string", "description": "Created poem"},
			"creative_process": map[string]string{"type": "string", "description": "Creative process"},
			"style_notes":      map[string]string{"type": "string", "description": "Style notes"},
		},
		"required": []string{"poem"},
	},
	"scene_description": {
		"type": "object",
		"properties": map[string]interface{}{
			"scene_description": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"main_subjects": map[string]string{"type": "string", "description": "Main objects and spatial relationships"},
					"lighting":      map[string]string{"type": "string", "description": "Lighting conditions and atmosphere"},
					"composition":   map[string]string{"type": "string", "description": "Picture composition"},
					"style":         map[string]string{"type": "string", "description": "Artistic style"},
				},
				"required": []string{"main_subjects", "lighting", "composition", "style"},
			},
		},
		"required": []string{"scene_description"},
	},
	"translation": {
		"type": "object",
		"properties": map[string]interface{}{
			"translation":       map[string]string{"type": "string", "description": "Translation result"},
			"original_language": map[string]string{"type": "string", "description": "Source language"},
			"target_language":   map[string]string{"type": "string", "description": "Target language"},
			"cultural_notes":    map[string]string{"type": "string", "description": "Cultural background notes"},
		},
		"required": []string{"translation"},
	},
}

// GetPredefinedSchema gets a deep copy of a predefined schema.
func GetPredefinedSchema(schemaName string) (map[string]interface{}, error) {
	schema, ok := PREDEFINED_SCHEMAS[schemaName]
	if !ok {
		return nil, fmt.Errorf("unknown schema name: %s", schemaName)
	}
	// Return a copy to prevent external modification of the original map
	return maps.Clone(schema), nil
}
