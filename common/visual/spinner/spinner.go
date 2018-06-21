package spinner

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
)

var pkgSpinner *spinner.Spinner

// Start start the cli-spinner with provided message
func Start(message string) {
	pkgSpinner.Suffix = fmt.Sprintf(" %s", message)
	pkgSpinner.Start()
}

// UpdateMessage update the message of the running cli-spinner
func UpdateMessage(message string) {
	Start(message)
}

// Stop stop the cli-spinner
func Stop() {
	pkgSpinner.Stop()
}

// StopWithMessage stop the cli-spinner and print the provided message
func StopWithMessage(message string) {
	pkgSpinner.Stop()
	fmt.Println(message)
}

func init() {
	pkgSpinner = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	pkgSpinner.Color("cyan")
	pkgSpinner.Stop()
}
