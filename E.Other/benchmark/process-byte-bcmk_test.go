package main

import "testing"

func BenchmarkProcessBytes(b *testing.B) {
	data := make([]byte, 1024*1024) // 1 MB
	for i := 0; i < b.N; i++ {
		processBytes(data)
	}
}

func processBytes(data []byte) int {
	return len(data)
}
