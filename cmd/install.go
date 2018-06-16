package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install kit dependencies",
	Long:  description("Install kit dependencies"),
	Run: func(cmd *cobra.Command, args []string) {
		installLocalTools, err := cmd.LocalFlags().GetBool("local-tools")
		if err != nil {
			panic(err)
		}

		installCITools, err := cmd.LocalFlags().GetBool("ci-tools")
		if err != nil {
			panic(err)
		}

		if installLocalTools {
			println("Installing local tools...")
			shell(filepath.Join(directories.scripts, "install_local_tools.sh"))
		}

		if installCITools {
			println("Installing ci tools...")
		}
	},
}

func init() {
	installCmd.Flags().SortFlags = false
	installCmd.Flags().Bool("local-tools", false, "install dependencies for local development environment")
	installCmd.Flags().Bool("ci-tools", false, "install dependencies for ci environment")
}
