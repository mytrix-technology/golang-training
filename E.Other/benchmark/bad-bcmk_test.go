package main

import "testing"

var result string

func BenchmarkBad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result = "cached value"
	}
}
