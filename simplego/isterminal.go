package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mattn/go-isatty"
	"golang.org/x/sys/unix"
)

func main() {
	// stdout
	_, err := unix.IoctlGetTermios(int(1), unix.TIOCGETA)
	if err != nil {
		fmt.Println(err.Error() + " stdout")
		log.Println("error:" + err.Error() + " stdout")
	}
	_, err = unix.IoctlGetTermios(int(2), unix.TIOCGETA)
	if err != nil {
		fmt.Println(err.Error() + " stderr")
		log.Println("error:" + err.Error() + " stderr")
	}
	if isatty.IsTerminal(os.Stdout.Fd()) {
		fmt.Println("Is Terminal")
	} else if isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		fmt.Println("Is Cygwin/MSYS2 Terminal")
	} else {
		fmt.Println("Is Not Terminal")
	}
}
