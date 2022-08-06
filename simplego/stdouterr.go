package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var stdout = log.New(os.Stdout, "info:", log.LstdFlags)
var stderr = log.New(os.Stderr, "error:", log.LstdFlags)

func main() {
	fmt.Print("no auto lf")
	log.Print("auto lf")
	fmt.Print("no auto lf")
	for {
		time.Sleep(1 * time.Second)
		// stdout.Print("\033[31mhello\n\033[0m")
		// stdout.Print("\033[31mhello\033[0m\n")
		fmt.Print("\033[31mhello\n\033[0m")
	}
}
