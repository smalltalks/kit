package script

import (
	kitdirectories "github.com/smalltalks/kit/cmd/common/directories"
	"github.com/smalltalks/kit/cmd/common/meta"
	"github.com/spf13/cobra"
)

var (
	banner      = meta.GetBanner()
	directories = kitdirectories.Get()
)

// scriptCmd represents the script command
var scriptCmd = &cobra.Command{
	Use:   "script",
	Short: "Manage scripts",
	Long:  banner.Description("Manage scripts"),
}

// Command return the script command
func Command() *cobra.Command {
	return scriptCmd
}

func init() {
	cobra.EnableCommandSorting = false
	scriptCmd.AddCommand(runCmd)
	scriptCmd.AddCommand(listCmd)
}
