package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/smalltalks/kit/cli/project"
	"github.com/smalltalks/kit/common/meta"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kit",
	Short: "Smalltalks development tool-kit",
	Long:  "Smalltalks development tool-kit",
}

func init() {
	cobra.EnableCommandSorting = false
	rootCmd.AddCommand(project.Command())
}

func main() {
	fmt.Println(color.BlueString(meta.GetBanner()))
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
