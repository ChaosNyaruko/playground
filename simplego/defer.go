package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Println("deferred", i)
		}()
	}
}
