package util

import "testing"

func BenchmarkSquareRoot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SquareRoot(16)
	}
}

func BenchmarkAnotherSquareRoot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AnotherSquareRoot(16)
	}
}
