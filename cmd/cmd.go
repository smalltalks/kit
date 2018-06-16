package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var version = "v0.0.1-alpha"
var banner = color.BlueString(`
   ___  __    ___  _________
  |\  \|\  \ |\  \|\___   ___\
  \ \  \/  /|\ \  \|___ \  \_|
   \ \   ___  \ \  \   \ \  \
    \ \  \\ \  \ \  \   \ \  \
     \ \__\\ \__\ \__\   \ \__\
      \|__| \|__|\|__|    \|__|  %s
`, version)

func description(s string) string {
	return fmt.Sprintf("%s\n%s", banner, s)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:  "kit",
	Long: description("Development tool-kit"),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(runtimeCmd)
	rootCmd.AddCommand(versionCmd)
}
