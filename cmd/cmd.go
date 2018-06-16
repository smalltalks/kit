package cmd

import (
	"os"

	kitdirectories "github.com/smalltalks/kit/cmd/common/directories"
	"github.com/smalltalks/kit/cmd/common/meta"
	"github.com/smalltalks/kit/cmd/project"
	"github.com/smalltalks/kit/cmd/script"
	"github.com/spf13/cobra"
)

var (
	version     = meta.GetVersion()
	banner      = meta.GetBanner()
	directories = kitdirectories.Get()
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kit",
	Short: "SmallTalks development tool-kit",
	Long:  banner.Description("SmallTalks development tool-kit"),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.EnableCommandSorting = false
	rootCmd.AddCommand(project.Command())
	rootCmd.AddCommand(script.Command())
}
