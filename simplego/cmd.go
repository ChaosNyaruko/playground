package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	fzf()
	return
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

func fzf() {
	// cmd := exec.Command("fzf", "--preview=\"echo {}\"")
	cmd := exec.Command("fzf", `--preview="echo {}"`)
	cmd.Stderr = os.Stderr
	in, _ := cmd.StdinPipe()
	var i int
	go func() {
		for {
			i += 1
			fmt.Fprintln(in, i)
			time.Sleep(1 * time.Second)
		}
		in.Close()
	}()

	res, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}
