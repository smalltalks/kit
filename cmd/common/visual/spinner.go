package visual

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
)

var pkgSpinner *spinner.Spinner

// StartSpinner start the cli-spinner with message
func StartSpinner(message string) {
	pkgSpinner.Suffix = fmt.Sprintf(" %s", message)
	pkgSpinner.Start()
}

// StopSpinner stop the cli-spinner
func StopSpinner() {
	pkgSpinner.Stop()
}

func init() {
	pkgSpinner = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	pkgSpinner.Color("cyan")
}
