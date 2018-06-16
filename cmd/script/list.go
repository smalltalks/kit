package script

import (
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available scripts under scripts directory",
	Long:  banner.Description("List all available scripts under scripts directory"),
	Run: func(cmd *cobra.Command, args []string) {
		files, err := ioutil.ReadDir(directories.Scripts)
		if err != nil {
			panic(err)
		}

		var scripts []os.FileInfo
		for _, f := range files {
			if !f.IsDir() {
				scripts = append(scripts, f)
			}
		}

		for _, f := range scripts {
			println(f.Name())
		}
	},
}
