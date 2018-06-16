package script

import (
	"path/filepath"

	"github.com/smalltalks/kit/cmd/common/shell"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run arbitrary script under scripts directory",
	Long:  banner.Description("Run an arbitrary script under scripts directory"),
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		script := filepath.Join(directories.Scripts, args[0])
		shell.Exec(script, args[1:]...)
	},
}
