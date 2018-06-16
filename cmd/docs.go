package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// docsCmd represents the docs command
var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Generate kit documentation",
	Long:  description("Generate kit documentation"),
	Run: func(cmd *cobra.Command, args []string) {
		out, err := cmd.LocalFlags().GetString("out")
		if err != nil {
			panic(err)
		}

		fmt.Printf("Generating markdown documentation to %s\n", out)
		doc.GenMarkdownTree(rootCmd, out)
	},
}

func init() {
	docsCmd.Flags().SortFlags = false
	docsCmd.Flags().String("out", directories.docs, "output directory")
}
