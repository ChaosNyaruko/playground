package main

import (
	"fmt"
	"time"
)

func main() {
	ch := GenerateNatural() // 自然数序列: 2, 3, 4, ...
	start := time.Now()
	for i := 0; i < 50_000; i++ {
		prime := <-ch               // 新出现的素数
		ch = PrimeFilter(ch, prime) // 基于新素数构造的过滤器
	}
	fmt.Printf("over, cost: %v\n", time.Since(start))
}

// 管道过滤器: 删除能被素数整除的数
func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

// 返回生成自然数序列的管道: 2, 3, 4, ...
func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}
