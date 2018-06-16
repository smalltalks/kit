package project

import (
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project",
	Long:  banner.Description("Create a new project"),
	Run: func(cmd *cobra.Command, args []string) {
		println("not implemented")
	},
}
