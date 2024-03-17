package handlers

import (
	"fmt"
	"os/exec"
	"runtime"
	"ruopen/tui"
)

func Open(filePath string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default:
		cmd = "xdg-open"
	}
	args = append(args, filePath)
	fmt.Println(tui.DefaultStyle.Render(fmt.Sprintf("[OPEN] %s", filePath)))
	return exec.Command(cmd, args...).Start()
}
