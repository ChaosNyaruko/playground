package main

import (
	"log"
	"testing"
)

func foo(n int) int {
	if n == 1 {
		return 1
	}
	return n * foo(n-1)
}

func BenchmarkFoo(b *testing.B) {
	b.Logf("starting benching foo..., b.N=%d", b.N)
	for i := 0; i < 1*b.N; i++ {
		foo(5)
	}
}

func BenchmarkFooNoN(b *testing.B) {
	defer log.Printf("b.N = %d", b.N)
	foo(5)
}

func BenchmarkFoo2(b *testing.B) {
	defer log.Printf("b.N = %d", b.N)
	for i := 0; i < 2*b.N; i++ {
		foo(5)
	}
}
