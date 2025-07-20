package image

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// SaveImages saves all images from ImageResponse to local files.
// It returns a list of successfully saved file paths and a slice of errors
// encountered during saving individual files.
func (g *Generator) SaveImages(imageResponse *ImageResponse, outputDir, baseFilename string) (savedFiles []string, saveErrors []error) {
	if !imageResponse.Success() {
		return nil, []error{fmt.Errorf("image response is not successful, cannot save")}
	}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return nil, []error{fmt.Errorf("failed to create output directory: %w", err)}
	}

	savedFiles = make([]string, 0, len(imageResponse.Images))
	saveErrors = make([]error, 0)

	for i, imageInfo := range imageResponse.Images {
		var imageData []byte
		var err error

		if imageInfo.ImageData != "" {
			imageData, err = base64.StdEncoding.DecodeString(imageInfo.ImageData)
			if err != nil {
				saveErrors = append(saveErrors, fmt.Errorf("failed to decode base64 for image %d: %w", i+1, err))
				continue
			}
		} else if imageInfo.URL != "" {
			imageData, err = downloadImage(imageInfo.URL)
			if err != nil {
				saveErrors = append(saveErrors, fmt.Errorf("failed to download image %d from %s: %w", i+1, imageInfo.URL, err))
				continue
			}
		} else {
			saveErrors = append(saveErrors, fmt.Errorf("image %d has no available data source, skipping", i+1))
			continue
		}

		ext := getFileExtension(imageInfo.MimeType)
		var filename string
		if len(imageResponse.Images) == 1 {
			filename = fmt.Sprintf("%s.%s", baseFilename, ext)
		} else {
			filename = fmt.Sprintf("%s_%d.%s", baseFilename, i+1, ext)
		}

		filepath := filepath.Join(outputDir, filename)

		err = os.WriteFile(filepath, imageData, 0644)
		if err != nil {
			saveErrors = append(saveErrors, fmt.Errorf("failed to save image %d to %s: %w", i+1, filepath, err))
			continue
		}

		savedFiles = append(savedFiles, filepath)
	}

	return savedFiles, saveErrors
}

func downloadImage(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)
}

func getFileExtension(mimeType string) string {
	mimeToExt := map[string]string{
		"image/png":  "png",
		"image/jpeg": "jpg",
		"image/jpg":  "jpg",
		"image/webp": "webp",
		"image/gif":  "gif",
		"image/bmp":  "bmp",
	}

	if ext, ok := mimeToExt[strings.ToLower(mimeType)]; ok {
		return ext
	}
	return "png" // Default to png
}
