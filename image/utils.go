package image

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// SaveImages saves all images from ImageResponse to local files.
func (g *Generator) SaveImages(imageResponse *ImageResponse, outputDir, baseFilename string) ([]string, error) {
	if !imageResponse.Success() {
		log.Println("Image response is not successful, cannot save")
		return nil, nil
	}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create output directory: %w", err)
	}

	savedFiles := make([]string, 0, len(imageResponse.Images))
	log.Printf("Starting to save %d images...\n", len(imageResponse.Images))

	for i, imageInfo := range imageResponse.Images {
		var imageData []byte
		var err error

		if imageInfo.ImageData != "" {
			imageData, err = base64.StdEncoding.DecodeString(imageInfo.ImageData)
			if err != nil {
				log.Printf("❌ Failed to decode base64 for image %d: %v\n", i+1, err)
				continue
			}
		} else if imageInfo.URL != "" {
			imageData, err = downloadImage(imageInfo.URL)
			if err != nil {
				log.Printf("❌ Failed to download image %d from %s: %v\n", i+1, imageInfo.URL, err)
				continue
			}
		} else {
			log.Printf("Image %d has no available data source, skipping\n", i+1)
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
			log.Printf("❌ Failed to save image %d to %s: %v\n", i+1, filepath, err)
			continue
		}

		savedFiles = append(savedFiles, filepath)
		log.Printf("✅ Image %d saved: %s (%d bytes)\n", i+1, filepath, len(imageData))
	}

	log.Printf("Successfully saved %d images\n", len(savedFiles))
	return savedFiles, nil
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
