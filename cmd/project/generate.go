package project

import (
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "(Re)Generate project code",
	Long:  banner.Description("(Re)Generate project code"),
	Run: func(cmd *cobra.Command, args []string) {
		println("not implemented")
	},
}
