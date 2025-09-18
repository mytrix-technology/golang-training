package main

import "testing"

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = 1 + 1 // Replace this with your actual code
	}
}
