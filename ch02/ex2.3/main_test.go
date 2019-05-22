package main

import (
	"testing"

	"github.com/tesujiro/ProgrammingLanguageGo/ch02/ex2.3/popcount"
)

func BenchmarkMain(b *testing.B) {

	b.ResetTimer()

	b.Run("Loop", func(b *testing.B) {
		benchmark(b, popcount.PopCount2)
	})

	b.Run("OneStep", func(b *testing.B) {
		benchmark(b, popcount.PopCount)
	})

	return
}

func benchmark(b *testing.B, fn func(uint64) int) {
	b.ResetTimer()
	for i := 1; i < b.N; i++ {
		fn(uint64(i))
	}
}
