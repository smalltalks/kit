package project

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/smalltalks/kit/common/visual/spinner"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project",
	Long:  "Create a new project",
	Run: func(cmd *cobra.Command, args []string) {
		projectName, err := cmd.Flags().GetString("name")
		if err != nil {
			panic(err)
		}

		spinner.Start(fmt.Sprintf("Creating project %s", projectName))
		<-time.After(5 * time.Second)
		spinner.UpdateMessage("Generating templates")
		<-time.After(5 * time.Second)
		spinner.StopWithMessage("Done generating project")
	},
}

func init() {
	createCmd.Flags().String("name", "", "name of the project (required)")
	createCmd.MarkFlagRequired("name")
}
