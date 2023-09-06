package gohtml

import "testing"

func Benchmark_Speed(b *testing.B) {
	b.Run("DSL", func(b *testing.B) {
		for i := 0; i < b.N; i++ {

		}
	})

	b.Run("template", func(b *testing.B) {
		for i := 0; i < b.N; i++ {

		}
	})
}
