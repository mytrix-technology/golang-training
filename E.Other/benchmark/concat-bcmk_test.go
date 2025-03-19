package main

import (
	"strings"
	"testing"
)

func BenchmarkConcat(b *testing.B) {
	words := []string{"Go", "is", "fast", "and", "awesome"}
	for i := 0; i < b.N; i++ {
		_ = joinStrings(words)
	}
}

func joinStrings(words []string) string {
	return strings.Join(words, " ")
}
