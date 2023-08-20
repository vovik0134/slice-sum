package slice_sum

import (
	"math/rand"
	"testing"
)

func BenchmarkClassicSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ClassicSum(data)
	}
}

func BenchmarkVectorized1Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Vectorized1Sum(data)
	}
}

func BenchmarkVectorized2Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Vectorized2Sum(data)
	}
}

func BenchmarkVectorized4Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Vectorized4Sum(data)
	}
}

func BenchmarkVectorized4Sum2Acc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Vectorized4Sum2Acc(data)
	}
}

func BenchmarkVectorized4Sum2AccNoInline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Vectorized4Sum2AccNoInline(data)
	}
}

func BenchmarkVectorized8Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Vectorized8Sum(data)
	}
}

func BenchmarkVectorized8Sum2Acc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Vectorized8Sum2Acc(data)
	}
}

func BenchmarkVectorized16Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Vectorized16Sum(data)
	}
}

func BenchmarkVectorized16Sum2Acc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Vectorized16Sum2Acc(data)
	}
}

const n = 100_000_000

var data []int64

func init() {
	data = make([]int64, n)
	for i := 0; i < n; i++ {
		data[i] = rand.Int63()
	}
}
