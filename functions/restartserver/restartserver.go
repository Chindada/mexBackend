package restartserver

import (
	"os/exec"
)

// Restart Restart
func Restart() {
	exec.Command("cmd.exe", "/C", "Restart.bat").Run()
}
