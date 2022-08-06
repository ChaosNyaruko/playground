package main

import (
	"log"
	"os"
)

var stdout = log.New(os.Stdout, "info:", log.LstdFlags)
var stderr = log.New(os.Stderr, "error:", log.LstdFlags)

func main() {
	stdout.Println("hello")
	stderr.Println("bye")
}
