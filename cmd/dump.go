package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"thelaserunicorn/goku/internal/db"

	"github.com/spf13/cobra"
)

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Print all resources from the database",
	Long:  `Retrieves all resources from resource_table and prints them in JSON format`,
	RunE:  runDump,
}

func init() {
	rootCmd.AddCommand(dumpCmd)
}

func runDump(cmd *cobra.Command, args []string) error {
	// Connect to database
	database, err := db.New(db.GetConfig())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return err
	}
	defer database.Close()

	// Get all resources
	resources, err := database.GetAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return err
	}

	if len(resources) == 0 {
		fmt.Println("No resources found in the database")
		return nil
	}

	// Print resources as JSON
	output, err := json.MarshalIndent(resources, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return err
	}

	fmt.Println(string(output))
	return nil
}
