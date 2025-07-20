
# RCK Go SDK - Programming Guide

An elegant Go SDK that uses RCK (Relational Calculate Kernel) as an intelligent function kernel.

[![Go Version](https://img.shields.io/badge/go-1.18+-blue.svg)](https://golang.org)
[![Go Report Card](https://goreportcard.com/badge/github.com/Askr-Omorsablin/rck-go-sdk)](https://goreportcard.com/report/github.com/Askr-Omorsablin/rck-go-sdk)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

## 🚀 Quick Start

### Installation

```bash
go get github.com/Askr-Omorsablin/rck-go-sdk
```

### Basic Usage

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Askr-Omorsablin/rck-go-sdk"
	"github.com/Askr-Omorsablin/rck-go-sdk/compute"
)

func main() {
	// Initialize client
	client, err := rck.NewClient("your-api-key")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Get predefined schema
	schema, err := compute.GetPredefinedSchema("basic_analysis")
	if err != nil {
		log.Fatalf("Failed to get schema: %v", err)
	}

	// Use RCK as intelligent function kernel
	result, err := client.Compute.CustomCompute(context.Background(), compute.CustomComputeParams{
		Text:         "Spring has arrived, all things are reviving",
		Task:         "Analyze sentiment and generate corresponding poetry",
		OutputSchema: schema,
	})
	if err != nil {
		log.Fatalf("Computation failed: %v", err)
	}

	fmt.Printf("Analysis result: %v\n", result.Data)
}
```

### Expected Output
```
Analysis result: map[analysis:The text expresses a positive sentiment associated with the arrival of spring and the revitalization of nature, evoking feelings of joy and optimism. emotion:Joyful and optimistic theme:Renewal and rebirth]
```

## 📋 Complete Working Example

Here's a comprehensive, tested example demonstrating all major SDK features:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Askr-Omorsablin/rck-go-sdk"
	"github.com/Askr-Omorsablin/rck-go-sdk/compute"
	"github.com/Askr-Omorsablin/rck-go-sdk/image"
)

func main() {
	apiKey := os.Getenv("RCK_API_KEY")
	if apiKey == "" {
		apiKey = "your-api-key-here" // Replace with your API Key
	}

	// Create client with functional options
	client, err := rck.NewClient(apiKey,
		rck.WithTimeout(60*time.Second), // Optional: set timeout
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	runCompleteExample(client)
}

func runCompleteExample(client *rck.Client) {
	fmt.Println("=== RCK Go SDK Complete Example ===\n")
	ctx := context.Background()

	// 1. Test connection
	fmt.Println("1. Testing connection...")
	if err := client.TestConnection(ctx); err != nil {
		fmt.Printf("❌ Connection failed: %v\n", err)
		return
	}
	fmt.Println("✅ Connection successful")

	// 2. Text analysis with predefined schema
	fmt.Println("\n2. Text Analysis Example:")
	poemText := "Moonlight before my bed, looks like frost on the ground. I raise my head to see the moon, lower it to think of home."
	
	analysisResult, err := client.Compute.Analyze(ctx, compute.AnalyzeParams{
		Text:         poemText,
		Task:         "Analyze the emotion and theme of this poetry",
		OutputFormat: "basic_analysis",
	})
	if err == nil {
		fmt.Printf("Original text: %s\n", poemText)
		fmt.Println("Analysis result:")
		for key, value := range analysisResult.Data {
			fmt.Printf("  %s: %v\n", key, value)
		}
	} else {
		fmt.Printf("Text analysis failed: %v\n", err)
	}

	// 3. Translation functionality
	fmt.Println("\n3. Translation Example:")
	translationResult, err := client.Compute.Translate(ctx, compute.TranslateParams{
		Text:                 poemText,
		TargetLanguage:       "French",
		IncludeCulturalNotes: true,
	})
	if err == nil {
		if translation, ok := translationResult.Data["translation"]; ok {
			fmt.Printf("French translation: %s\n", translation)
		}
		if notes, ok := translationResult.Data["cultural_notes"]; ok {
			fmt.Printf("Cultural notes: %.100s...\n", fmt.Sprintf("%v", notes))
		}
	}

	// 4. Custom schema computation with proper error handling
	fmt.Println("\n4. Custom Schema Example:")
	poemSchema, err := compute.GetPredefinedSchema("poem_creation")
	if err != nil {
		fmt.Printf("Failed to get poem schema: %v\n", err)
		return
	}

	poemCreation, err := client.Compute.CustomCompute(ctx, compute.CustomComputeParams{
		Text:         "Spring flowers blooming in the garden",
		Task:         "Create a poem based on this theme",
		OutputSchema: poemSchema,
		CustomFields: map[string]string{
			"style": "modern free verse",
			"mood":  "joyful and peaceful",
		},
	})
	if err == nil {
		if poem, ok := poemCreation.Data["poem"]; ok {
			fmt.Printf("Created poem: %s\n", poem)
		}
		if process, ok := poemCreation.Data["creative_process"]; ok {
			fmt.Printf("Creative process: %.100s...\n", fmt.Sprintf("%v", process))
		}
	} else {
		fmt.Printf("Poem creation failed: %v\n", err)
	}

	// 5. Image generation
	fmt.Println("\n5. Image Generation Example:")
	imageResult, err := client.Image.Generate(ctx, image.GenerateParams{
		Prompt:      "A serene mountain lake reflecting the sky",
		Composition: "Wide panoramic view showcasing natural beauty",
		Lighting:    "Soft morning sunlight creating peaceful atmosphere",
		Style:       "Traditional Chinese landscape painting style with ink wash",
	})
	if err == nil && imageResult.Success() {
		fmt.Printf("✅ Image generation successful: %d images created\n", imageResult.Count)
		
		// Save images (optional)
		savedFiles, saveErrors := client.Image.SaveImages(imageResult, ".", "example_landscape")
		if len(saveErrors) > 0 {
			fmt.Printf("Some images failed to save: %v\n", saveErrors)
		} else {
			fmt.Printf("Images saved: %v\n", savedFiles)
		}
	} else {
		fmt.Printf("❌ Image generation failed: %v\n", err)
	}

	// 6. Advanced workflow: Poem to Image
	fmt.Println("\n6. Advanced Workflow - Poem to Scene to Image:")
	originalPoem := "Empty mountains, no one in sight, but hearing human voices echo. Returning light enters deep forest, again illuminating green moss."

	// Step 1: Convert poem to scene description
	sceneSchema, err := compute.GetPredefinedSchema("scene_description")
	if err != nil {
		fmt.Printf("Failed to get scene schema: %v\n", err)
		return
	}

	sceneResult, err := client.Compute.CustomCompute(ctx, compute.CustomComputeParams{
		Text:         originalPoem,
		Task:         "Analyze the poem content and create detailed visual scene description",
		OutputSchema: sceneSchema,
		CustomFields: map[string]string{"target_art_style": "Traditional Chinese landscape painting"},
	})

	if err == nil {
		if sceneData, ok := sceneResult.Data["scene_description"].(map[string]interface{}); ok {
			fmt.Println("Scene conversion successful:")
			if subjects, ok := sceneData["main_subjects"]; ok {
				fmt.Printf("  Main subjects: %s\n", subjects)
			}

			// Step 2: Generate image from scene description
			workflowImage, err := client.Image.Generate(ctx, image.GenerateParams{
				Prompt:      fmt.Sprintf("%v", sceneData["main_subjects"]),
				Composition: fmt.Sprintf("%v", sceneData["composition"]),
				Lighting:    fmt.Sprintf("%v", sceneData["lighting"]),
				Style:       fmt.Sprintf("%v", sceneData["style"]),
			})
			if err == nil && workflowImage.Success() {
				fmt.Printf("  ✅ Workflow completed: Poem → Scene → Image (%d images)\n", workflowImage.Count)
			} else {
				fmt.Printf("  ❌ Image generation step failed: %v\n", err)
			}
		}
	} else {
		fmt.Printf("Scene description conversion failed: %v\n", err)
	}

	fmt.Println("\n=== Example Complete ===")
}
```

## 🧠 RCK Compute Engine Core Concepts

RCK compute engine is based on two core components: **Start Point** and **Path**, allowing you to encapsulate complex AI logic into simple Go functions.

### Start Point: Define Initial State

`start_point` is the input for computation, containing two parts:

- **startPoint** (string): Core text prompt
- **resource** (array, optional): Additional non-text resources like images

```json
// Pure text input
{
    "start_point": {
        "startPoint": "Moonlight before my bed, looks like frost on the ground"
    }
}

// Multimodal input (text + images)
{
    "start_point": {
        "startPoint": "Please combine the desolation of 'Image One' with the vastness of 'Image Two' to describe a scene.",
        "resource": [
            {"Image One": "https://url.to/image1.png"},
            {"Image Two": "https://url.to/image2.png"}
        ]
    }
}
```

### Path: Apply Constraints and Define Goals

`path` is a declarative constraint on the transformation process:

- **expectPath** (string): Core instruction telling AI "what to do"
- **Custom Fields** (any): Any custom fields as auxiliary constraints

```json
{
    "path": {
        "expectPath": "Analyze the emotional tone of the poetry and create a modern poem with corresponding mood",
        "style": "modern free verse",
        "mood": "tranquil and profound",
        "target_length": "4-6 lines"
    }
}
```

## 🎯 Recommended Usage Pattern: Functional Encapsulation

Use RCK as the intelligent kernel of functions, encapsulating complex AI logic into simple Go functions:

### Example 1: Sentiment Analysis Function

```go
import (
	"context"
	"fmt"

	"github.com/Askr-Omorsablin/rck-go-sdk"
	"github.com/Askr-Omorsablin/rck-go-sdk/compute"
)

// EmotionResult defines the return structure for emotion analysis
type EmotionResult struct {
	Emotion   string   `json:"emotion"`
	Intensity float64  `json:"intensity"`
	Keywords  []string `json:"keywords"`
}

func AnalyzeEmotion(client *rck.Client, ctx context.Context, text string) (*EmotionResult, error) {
	schema := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"emotion":   map[string]string{"type": "string"},
			"intensity": map[string]string{"type": "number"},
			"keywords":  map[string]interface{}{"type": "array", "items": map[string]string{"type": "string"}},
		},
		"required": []string{"emotion", "intensity", "keywords"},
	}

	result, err := client.Compute.CustomCompute(ctx, compute.CustomComputeParams{
		Text:         text,
		Task:         "Analyze the emotional tendency and intensity of the text",
		OutputSchema: schema,
	})
	if err != nil {
		return nil, err
	}

	var emotionResult EmotionResult
	if err := result.Decode(&emotionResult); err != nil {
		return nil, fmt.Errorf("failed to decode result: %w", err)
	}

	return &emotionResult, nil
}

// Usage
// emotion, err := AnalyzeEmotion(client, context.Background(), "Today is sunny and I feel particularly good!")
```

### Example 2: Intelligent Summary Generation Function

```go
func IntelligentSummary(client *rck.Client, ctx context.Context, content string, maxLength int, style string) (string, error) {
	schema, err := compute.GetPredefinedSchema("basic_analysis")
	if err != nil {
		return "", fmt.Errorf("failed to get schema: %w", err)
	}

	result, err := client.Compute.CustomCompute(ctx, compute.CustomComputeParams{
		Text:         content,
		Task:         fmt.Sprintf("Generate a summary within %d words", maxLength),
		OutputSchema: schema,
		CustomFields: map[string]string{
			"style":         style,
			"focus":         "core viewpoints",
			"output_format": "concise and clear",
		},
	})
	if err != nil {
		return "", err
	}

	summary, ok := result.Data["analysis"].(string)
	if !ok {
		return "", fmt.Errorf("failed to extract summary from result")
	}
	return summary, nil
}

// Usage
// summary, err := IntelligentSummary(client, context.Background(), longText, 50, "academic")
```

### Example 3: Multimodal Creation Function

```go
func CreatePoemFromImage(client *rck.Client, ctx context.Context, imageURL, poemStyle string) (map[string]interface{}, error) {
	schema, err := compute.GetPredefinedSchema("poem_creation")
	if err != nil {
		return nil, fmt.Errorf("failed to get poem schema: %w", err)
	}

	result, err := client.Compute.CustomCompute(ctx, compute.CustomComputeParams{
		Text:         "Please create poetry based on the artistic conception of the image",
		Task:         "Observe image content, feel its artistic conception, and create poetry in corresponding style",
		OutputSchema: schema,
		Resources: []map[string]string{
			{"inspiration_image": imageURL},
		},
		CustomFields: map[string]string{
			"style":               poemStyle,
			"cultural_background": "classical literature",
		},
	})
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// Usage
// poemResult, err := CreatePoemFromImage(client, context.Background(), "https://example.com/sunset.jpg", "seven-character regulated verse")
```

## 📊 Available Predefined Schemas

The SDK provides several predefined output schemas for common use cases:

### 1. Basic Analysis (`basic_analysis`)
```go
schema, err := compute.GetPredefinedSchema("basic_analysis")
// Returns: emotion, theme, analysis fields
```

### 2. Poem Creation (`poem_creation`)  
```go
schema, err := compute.GetPredefinedSchema("poem_creation")
// Returns: poem, creative_process, style_notes fields
```

### 3. Scene Description (`scene_description`)
```go
schema, err := compute.GetPredefinedSchema("scene_description")
// Returns: main_subjects, lighting, composition, style fields
```

### 4. Translation (`translation`)
```go
schema, err := compute.GetPredefinedSchema("translation")
// Returns: translation, original_language, target_language, cultural_notes fields
```

## ⚙️ Configuration

Use functional options pattern to customize client behavior:

```go
import (
	"net/http"
	"time"
	
	"github.com/Askr-Omorsablin/rck-go-sdk"
)

// Create a client with custom configuration
client, err := rck.NewClient("your-api-key",
	rck.WithTimeout(90*time.Second),         // Set request timeout
	rck.WithHTTPClient(&http.Client{...}),   // Provide custom HTTP Client
)
```

## ⚠️ Error Handling

The SDK defines specific error types that you can check using `errors.Is` or type assertions:

```go
import (
	"errors"
	"github.com/Askr-Omorsablin/rck-go-sdk/sdkerrors"
)

_, err := client.Compute.Analyze(ctx, ...)
if err != nil {
	var apiErr *sdkerrors.APIError
	if errors.As(err, &apiErr) {
		// This is an API error (e.g., 4xx, 5xx)
		fmt.Printf("API error: %s, status code: %d\n", apiErr.Message, apiErr.StatusCode)
		if apiErr.ResponseData != nil {
			fmt.Printf("Response details: %v\n", apiErr.ResponseData)
		}
	} else if errors.Is(err, sdkerrors.ErrAuthentication) {
		// This is an authentication error
		fmt.Println("Authentication failed, please check your API Key.")
	} else {
		// Other network errors or unknown errors
		fmt.Printf("Unknown error occurred: %v\n", err)
	}
}
```

## 💡 Best Practices

### 1. Always Use Schemas
```go
// ❌ Bad: Missing OutputSchema may cause 400 errors
result, err := client.Compute.CustomCompute(ctx, compute.CustomComputeParams{
	Text: "Hello",
	Task: "Analyze",
	// Missing OutputSchema
})

// ✅ Good: Always provide OutputSchema
schema, _ := compute.GetPredefinedSchema("basic_analysis")
result, err := client.Compute.CustomCompute(ctx, compute.CustomComputeParams{
	Text:         "Hello",
	Task:         "Analyze",
	OutputSchema: schema,
})
```

### 2. Handle Timeouts Properly
```go
// Set appropriate timeout for your use case
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

result, err := client.Compute.CustomCompute(ctx, params)
```

### 3. Test Connection First
```go
// Test connection before making actual requests
if err := client.TestConnection(ctx); err != nil {
	log.Fatalf("Connection test failed: %v", err)
}
```

## 🌟 Unlimited Flexibility in Language and Format

RCK supports extreme flexibility:

### Any Language
```go
// English processing
schema, _ := compute.GetPredefinedSchema("basic_analysis")
result, _ := client.Compute.CustomCompute(ctx, compute.CustomComputeParams{
	Text:         "To be or not to be, that is the question",
	Task:         "Analyze the philosophical connotations of this Shakespeare quote",
	OutputSchema: schema,
})

// Chinese processing
result, _ = client.Compute.CustomCompute(ctx, compute.CustomComputeParams{
	Text:         "春眠不觉晓，处处闻啼鸟",
	Task:         "Translate to English while preserving poetic sentiment",
	OutputSchema: schema,
})

// Multi-language mixing
result, _ = client.Compute.CustomCompute(ctx, compute.CustomComputeParams{
	Text:         "Hello world, Bonjour le monde",
	Task:         "Identify languages and translate uniformly to English",
	OutputSchema: schema,
})
```

### Any Format
```go
// Mathematical formulas
schema, _ := compute.GetPredefinedSchema("basic_analysis")
result, _ := client.Compute.CustomCompute(ctx, compute.CustomComputeParams{
	Text:         "f(x) = x^2 - 4x + 3",
	Task:         "Find the minimum value of the function and describe the graph",
	OutputSchema: schema,
	CustomFields: map[string]string{"custom_code": "def calculate_min(x): return x**2 - 4*x + 3"},
})

// Code analysis
result, _ = client.Compute.CustomCompute(ctx, compute.CustomComputeParams{
	Text: `
	func fibonacci(n int) int {
		if n <= 1 {
			return n
		}
		return fibonacci(n-1) + fibonacci(n-2)
	}`,
	Task:         "Analyze code complexity and provide optimization suggestions",
	OutputSchema: schema,
	CustomFields: map[string]string{"language": "Go"},
})
```

## 🎨 Image Generation Features

In addition to text computation, the SDK provides powerful image generation capabilities:

```go
func GenerateArtwork(client *rck.Client, ctx context.Context, description, artStyle string) error {
	imageResponse, err := client.Image.Generate(ctx, image.GenerateParams{
		Prompt:      description,
		Composition: "centered composition with strong visual impact",
		Lighting:    "dramatic lighting effects",
		Style:       artStyle,
	})
	if err != nil {
		return err
	}

	if imageResponse.Success() {
		// Save images
		savedFiles, saveErrors := client.Image.SaveImages(imageResponse, ".", "artwork")
		if len(saveErrors) > 0 {
			return fmt.Errorf("image saving failed: %v", saveErrors)
		}
		fmt.Printf("Generation successful: %d images saved to %v\n", imageResponse.Count, savedFiles)
	} else {
		fmt.Println("Image generation failed")
	}
	
	return nil
}

// Usage
// err := GenerateArtwork(client, context.Background(), "A lonely traveler walking under the starry sky", "Van Gogh style oil painting")
```

## 🔧 Complete Example: Intelligent Assistant Function

```go
type IntelligentAssistant struct {
	client *rck.Client
}

func NewIntelligentAssistant(apiKey string) (*IntelligentAssistant, error) {
	client, err := rck.NewClient(apiKey)
	if err != nil {
		return nil, err
	}
	return &IntelligentAssistant{client: client}, nil
}

func (ia *IntelligentAssistant) ProcessRequest(ctx context.Context, userInput, context string) (map[string]interface{}, error) {
	// First analyze user intent
	schema, err := compute.GetPredefinedSchema("basic_analysis")
	if err != nil {
		return nil, fmt.Errorf("failed to get schema: %w", err)
	}

	intentAnalysis, err := ia.client.Compute.CustomCompute(ctx, compute.CustomComputeParams{
		Text:         userInput,
		Task:         "Analyze user intent and categorize",
		OutputSchema: schema,
		CustomFields: map[string]string{"context": context},
	})
	if err != nil {
		return nil, err
	}

	intent, _ := intentAnalysis.Data["theme"].(string)
	
	// Execute different processing logic based on intent
	switch intent {
	case "creative_writing":
		return ia.handleCreativeRequest(ctx, userInput)
	case "data_analysis":
		return ia.handleAnalysisRequest(ctx, userInput)
	default:
		return ia.handleGeneralRequest(ctx, userInput)
	}
}

func (ia *IntelligentAssistant) handleCreativeRequest(ctx context.Context, userInput string) (map[string]interface{}, error) {
	schema, _ := compute.GetPredefinedSchema("poem_creation")
	result, err := ia.client.Compute.CustomCompute(ctx, compute.CustomComputeParams{
		Text:         userInput,
		Task:         "Create content based on user needs",
		OutputSchema: schema,
		CustomFields: map[string]string{
			"creativity_level": "high",
			"style":           "engaging",
			"length":          "moderate",
		},
	})
	if err != nil {
		return nil, err
	}
	return result.Data, nil
}

func (ia *IntelligentAssistant) handleAnalysisRequest(ctx context.Context, userInput string) (map[string]interface{}, error) {
	schema, _ := compute.GetPredefinedSchema("basic_analysis")
	result, err := ia.client.Compute.CustomCompute(ctx, compute.CustomComputeParams{
		Text:         userInput,
		Task:         "Conduct in-depth analysis and provide insights",
		OutputSchema: schema,
		CustomFields: map[string]string{
			"analysis_depth":      "detailed",
			"include_suggestions": "yes",
			"format":              "structured",
		},
	})
	if err != nil {
		return nil, err
	}
	return result.Data, nil
}

func (ia *IntelligentAssistant) handleGeneralRequest(ctx context.Context, userInput string) (map[string]interface{}, error) {
	schema, _ := compute.GetPredefinedSchema("basic_analysis")
	result, err := ia.client.Compute.CustomCompute(ctx, compute.CustomComputeParams{
		Text:         userInput,
		Task:         "Provide helpful answers and suggestions",
		OutputSchema: schema,
		CustomFields: map[string]string{
			"tone":         "friendly",
			"detail_level": "moderate",
		},
	})
	if err != nil {
		return nil, err
	}
	return result.Data, nil
}

// Usage
// assistant, err := NewIntelligentAssistant("your-api-key")
// creativeResult, err := assistant.ProcessRequest(ctx, "Help me write a poem about autumn", "creative")
// analysisResult, err := assistant.ProcessRequest(ctx, "Analyze the trends in this sales data", "business")
```

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📞 Contact Support

For questions or assistance, please contact:

📧 **Email**: omorsablin@gmail.com

---

> 💡 **Core Philosophy**: Use RCK as an intelligent function kernel, describing "what to do" declaratively rather than "how to do it". Let AI handle complex logic while you only need to define inputs, constraints, and expected outputs.
