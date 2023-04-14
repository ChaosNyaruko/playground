package main

import (
	"os"
	"os/exec"
	"runtime"
)

func main() {
	command := "vim tmp.md"
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", command)
	} else {
		// cmd = exec.Command("sh", "-c", command)
		cmd = exec.Command("vim", "tmp.md")
	}
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
