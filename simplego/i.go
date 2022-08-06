package main

import (
	"fmt"
	"strconv"
)

type myType struct {
	a int
}

func (m myType) String() string {
	return "hello" + strconv.Itoa(m.a)
}

func main() {
	test(struct{ a int }{a: 2})
	test(myType{a: 3})
	fmt.Println(testi(myType{a: 4}))
}

func test(s struct{ a int }) {
	fmt.Printf("%#v\n", s)
}

func testi(i interface{ String() string }) string {
	return i.String()
}
