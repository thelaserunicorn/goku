package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"thelaserunicorn/goku/pkg/converter"

	"github.com/spf13/cobra"
)

var inputFile string
var outputFormat string

var rootCmd = &cobra.Command{
	Use:          "goku",
	Short:        "Goku is a CLI tool for converting between JSON and YAML formats.",
	RunE:         runConvert,
	SilenceUsage: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Path to the input file (JSON or YAML)")
	rootCmd.Flags().StringVarP(&outputFormat, "output", "o", "", "Desired output format (json or yaml)")
	rootCmd.MarkFlagRequired("input")
	rootCmd.MarkFlagRequired("output")
}

func runConvert(cmd *cobra.Command, args []string) error {
	outputFormat = strings.ToLower(outputFormat)
	if outputFormat != "json" && outputFormat != "yaml" {
		return fmt.Errorf("invalid output format: %s. Must be 'json' or 'yaml'", outputFormat)
	}

	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		return fmt.Errorf("input file not found: %s", inputFile)
	}

	result, err := converter.Convert(inputFile, converter.Format(outputFormat))
	if err != nil {
		return err
	}

	inputExt := filepath.Ext(inputFile)
	baseName := strings.TrimSuffix(filepath.Base(inputFile), inputExt)
	outputExt := "." + outputFormat
	outputPath := filepath.Join(filepath.Dir(inputFile), baseName+outputExt)
	if err := os.WriteFile(outputPath, result, 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	fmt.Println(string(result))
	return nil
}
