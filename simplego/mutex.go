package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.RWMutex
	mu.Lock()
	fmt.Println("locked")
	defer mu.Unlock()
	mu.RLock()
	fmt.Println("rlock after lock")
	defer mu.RUnlock()
}
