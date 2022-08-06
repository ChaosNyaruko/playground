package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type I interface {
	Go() string
}

// bubble sort
func bubble(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// select sort nums using select sort algorithm
func selectSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
}

type server1 struct {
}

func (s *server1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I'm server 1!\n"))
}

type server2 struct {
}

func (s *server2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I'm server 2!\n"))
}

func main() {
	log.Println("start")
	var wg sync.WaitGroup
	go func() {
		time.Sleep(1 * time.Second)
		wg.Add(1)
		defer wg.Done()
		log.Println("lauching server 2")
		log.Fatalf("server2:%v", http.ListenAndServe("localhost:8081", new(server2)))
	}()
	log.Println("lauching server 1")
	log.Fatalf("server1:%v", http.ListenAndServe("localhost:8081", new(server1)))

	wg.Wait()
}
func _main() {
	mi := my{}
	fmt.Println(mi.Go())
	c := make(chan int)
	go func() {
		select {
		case _ = <-c:
			log.Println("goroutine1")
		}
	}()
	go func() {
		select {
		case _ = <-c:
			log.Println("goroutine2")
		}
	}()
	go func() {
		select {
		case _ = <-c:
			log.Println("goroutine3")
		}
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c <- 1
		log.Println("sent")
	}()
	time.Sleep(1 * time.Minute)
}

type my struct{}

func (m my) Go() string {
	return "wft"
}
