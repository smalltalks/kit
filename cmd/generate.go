package cmd

import (
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate project code",
	Long:  description("Generate project code"),
	Run: func(cmd *cobra.Command, args []string) {
		println("not implemented")
	},
}
