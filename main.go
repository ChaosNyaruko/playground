package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("in shell", os.Getenv("mypath"))
	cmd := exec.Command("zsh", "subprocess.zsh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
