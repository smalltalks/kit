package project

import (
	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage projects",
	Long:  "Manage projects",
}

// Command return the project command
func Command() *cobra.Command {
	return projectCmd
}

func init() {
	cobra.EnableCommandSorting = false
	projectCmd.AddCommand(createCmd)
	projectCmd.AddCommand(generateCmd)
}
