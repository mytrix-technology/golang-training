package main

import "testing"

func BenchmarkSliceLookup(b *testing.B) {
	s := []struct {
		key   string
		value int
	}{{"a", 1}, {"b", 2}, {"c", 3}}
	for i := 0; i < b.N; i++ {
		for _, item := range s {
			if item.key == "a" {
				_ = item.value
				break
			}
		}
	}
}
