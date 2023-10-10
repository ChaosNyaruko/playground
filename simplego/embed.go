package main

import (
	"embed"
	"fmt"
)

//go:embed  io.go
var f embed.FS

func main() {
	fmt.Println("start")
	data, _ := f.ReadFile("io.go")
	print(string(data))
	fmt.Println("end")
}
