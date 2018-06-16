package project

import (
	"github.com/smalltalks/kit/cmd/common/meta"
	"github.com/spf13/cobra"
)

var (
	banner = meta.GetBanner()
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage projects",
	Long:  banner.Description("Manage projects"),
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
