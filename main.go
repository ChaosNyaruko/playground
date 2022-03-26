package main

import "fmt"

type I interface {
	Go() string
}

func main() {
	mi := my{}
	fmt.Println(mi.Go())
}

type my struct{}

func (m my) Go() string {
	return "wft"
}
