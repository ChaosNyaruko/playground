package main

import (
	"log"
	"time"
)

func ten(i int) {
	if i == 0 {
		time.Sleep(10 * time.Second)
	}
	log.Printf("ten end: %d\n", i)
}

func main() {
	t := time.NewTicker(2 * time.Second)
	i := 0
	for range t.C {
		log.Println("tick!", i)
		ten(i)
		i++
	}
}
