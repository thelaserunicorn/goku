package cmd

import (
	"fmt"
	"os"
	"strconv"

	"thelaserunicorn/goku/internal/db"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a resource by ID",
	Long:  `Deletes the resource with the specified ID from resource_table`,
	RunE:  runDelete,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func runDelete(cmd *cobra.Command, args []string) error {
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

	// Delete the resource
	if err := database.Delete(id); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return err
	}

	fmt.Printf("Resource with ID %d deleted successfully\n", id)
	return nil
}
