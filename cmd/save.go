package cmd

import (
	"fmt"
	"os"

	"thelaserunicorn/goku/internal/db"
	"thelaserunicorn/goku/internal/parser"

	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save a resource from a JSON or YAML file to the database",
	Long:  `Parses the input file (JSON or YAML) and saves the resource to resource_table`,
	RunE:  runSave,
}

func init() {
	rootCmd.AddCommand(saveCmd)
}

func runSave(cmd *cobra.Command, args []string) error {
	// Validate file format
	if err := parser.ValidateFileFormat(inputFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return err
	}

	// Parse the file
	data, err := parser.Parse(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return err
	}

	// Connect to database
	database, err := db.New(db.GetConfig())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return err
	}
	defer database.Close()

	// Initialize schema if needed
	if err := database.InitSchema(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return err
	}

	// Save the resource
	id, err := database.Save(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return err
	}

	fmt.Printf("Resource saved successfully with ID: %d\n", id)
	return nil
}
