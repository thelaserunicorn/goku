package converter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

type Format string

const (
	JSON Format = "json"
	YAML Format = "yaml"
)

type configData map[string]interface{}

func DetectFormat(pathForFile string) (Format, error) {
	ext := strings.ToLower(filepath.Ext(pathForFile))
	switch ext {
	case ".json":
		return JSON, nil
	case ".yaml", ".yml":
		return YAML, nil
	default:
		return "", fmt.Errorf("unsupported file format: %s", ext)
	}
}

func ReadFile(pathForFile string) (configData, Format, error) {
	format, err := DetectFormat(pathForFile)
	if err != nil {
		return nil, "", err
	}

	data, err := os.ReadFile(pathForFile)
	if err != nil {
		return nil, "", fmt.Errorf("failed to read file: %w", err)
	}

	var config configData
	switch format {
	case JSON:
		err = json.Unmarshal(data, &config)
	case YAML:
		err = yaml.Unmarshal(data, &config)
	}
	if err != nil {
		return nil, "", fmt.Errorf("failed to parse file as %s: %w", format, err)
	}

	return config, format, nil
}

func (c configData) ToYAML() ([]byte, error) {
	out, err := yaml.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to YAML: %w", err)
	}
	return out, nil
}

func (c configData) ToJSON() ([]byte, error) {
	out, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to convert to JSON: %w", err)
	}
	return out, nil
}

func Convert(inputPath string, outputFormat Format) ([]byte, error) {
	config, inputFormat, err := ReadFile(inputPath)
	if err != nil {
		return nil, err
	}

	if inputFormat == outputFormat {
		return nil, fmt.Errorf("input and output formats are the same: %s", inputFormat)
	}
	switch outputFormat {
	case JSON:
		return config.ToJSON()
	case YAML:
		return config.ToYAML()
	default:
		return nil, fmt.Errorf("unsupported output format: %s", outputFormat)
	}
}
