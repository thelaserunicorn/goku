package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"thelaserunicorn/goku/internal/db"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a resource by ID",
	Long:  `Retrieves a single resource by its ID and prints it as JSON`,
	RunE:  runGet,
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func runGet(cmd *cobra.Command, args []string) error {
	// Parse the ID
	id, err := strconv.ParseInt(inputID, 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid ID '%s'. Must be a number.\n", inputID)
		return err
	}

	// Connect to database
	database, err := db.New(db.GetConfig())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return err
	}
	defer database.Close()

	// Get the resource
	resource, err := database.GetByID(id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return err
	}

	// Print as JSON
	output, err := json.MarshalIndent(resource, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return err
	}

	fmt.Println(string(output))
	return nil
}
