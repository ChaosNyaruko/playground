package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(os.Environ())
	io.Copy(os.Stdout, os.Stdin)
}
