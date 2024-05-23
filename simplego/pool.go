package main

import (
	"fmt"
	"runtime"
	"sync"
)

type MyObject struct{}

func main() {
	p := &sync.Pool{
		New: func() interface{} { return new(MyObject) },
	}

	for i := 0; i < 100000; i++ {
		obj := p.Get().(*MyObject)
		// Use the object...
		p.Put(obj)
	}
	var r runtime.MemStats
	runtime.ReadMemStats(&r)
	fmt.Printf("GC stats: %#v\n", r)
	select {}
}
