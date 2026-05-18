package cmd

import (
	"fmt"
	"os"
	"strconv"

	"thelaserunicorn/goku/internal/db"
	"thelaserunicorn/goku/internal/parser"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing resource from a JSON or YAML file",
	Long:  `Parses the input file (JSON or YAML) and updates the resource with the specified ID`,
	RunE:  runUpdate,
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func runUpdate(cmd *cobra.Command, args []string) error {
	// Parse the ID
	id, err := strconv.ParseInt(inputID, 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid ID '%s'. Must be a number.\n", inputID)
		return err
	}

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

	// Update the resource
	if err := database.Update(id, data); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return err
	}

	fmt.Printf("Resource with ID %d updated successfully\n", id)
	return nil
}
