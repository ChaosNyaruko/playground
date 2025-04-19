package main

import "fmt"

func main() {
	m := map[int]int{1: 2, 3: 4}
	for k, v := range m {
		fmt.Printf("%v, %v\n", k, v)
		m[k+v] = k - v
	}
	// 1,2 3:4
}
