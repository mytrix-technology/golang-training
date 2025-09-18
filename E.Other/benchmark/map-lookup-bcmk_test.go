package main

import "testing"

func BenchmarkMapLookup(b *testing.B) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for i := 0; i < b.N; i++ {
		_ = m["a"]
	}
}
