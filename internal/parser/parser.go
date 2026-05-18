package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"thelaserunicorn/goku/pkg/converter"

	yaml "gopkg.in/yaml.v3"
)

// Parse reads a JSON or YAML file and returns the parsed data as a map.
// It reuses the converter package from Phase 1.
func Parse(filePath string) (map[string]interface{}, error) {
	// Check file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file not found: %s", filePath)
	}

	// Detect format and validate
	if _, err := converter.DetectFormat(filePath); err != nil {
		return nil, err
	}

	// Read and parse file
	data, _, err := converter.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ParseFromData parses raw data (bytes) given a format string ("json" or "yaml")
func ParseFromData(data []byte, format string) (map[string]interface{}, error) {
	config := make(map[string]interface{})

	switch format {
	case "json":
		if err := json.Unmarshal(data, &config); err != nil {
			return nil, fmt.Errorf("failed to parse JSON: %w", err)
		}
	case "yaml":
		if err := yaml.Unmarshal(data, &config); err != nil {
			return nil, fmt.Errorf("failed to parse YAML: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}

	return config, nil
}

// ValidateFileFormat checks if the file has a valid .json or .yaml/.yml extension
func ValidateFileFormat(filePath string) error {
	ext := filepath.Ext(filePath)
	switch ext {
	case ".json", ".yaml", ".yml":
		return nil
	default:
		return fmt.Errorf("unsupported file format: %s. Must be .json, .yaml, or .yml", ext)
	}
}
