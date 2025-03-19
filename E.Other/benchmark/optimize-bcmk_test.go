package main

import "testing"

var sink interface{}

func BenchmarkOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink = expensiveFunction()
	}
}

func expensiveFunction() int {
	return 42
}
