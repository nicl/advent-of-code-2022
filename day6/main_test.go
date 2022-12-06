package main

import "testing"

func BenchmarkFindMarker(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FindMarker([]byte(input), 14)
	}
}
