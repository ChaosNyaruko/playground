package main

import (
	"fmt"
	"io"
	"os"
)

// let's modify" ittttt
func main() {
	fmt.Println(os.Environ())
	io.Copy(os.Stdout, os.Stdin)
}
