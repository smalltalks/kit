package shell

import (
	"os"
	"os/exec"
)

// Exec run a command on native shell with piped back stdout & stderr
func Exec(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
