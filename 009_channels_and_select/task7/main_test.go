package main

import (
	"fmt"
	"testing"
)

var sizes = []int{0, 1, 100, 10_000}
var benchmarkResult int

func TestSums(t *testing.T) {
	want := N * (N - 1) / 2

	for _, size := range sizes {
		if got := passN(size); got != want {
			t.Errorf("passN(%d) = %d, want %d", size, got, want)
		}
	}

	if got := passBatches(); got != want {
		t.Errorf("passBatches() = %d, want %d", got, want)
	}
}

func BenchmarkPassN(b *testing.B) {
	for _, size := range sizes {
		b.Run(fmt.Sprintf("buf=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				benchmarkResult = passN(size)
			}
		})
	}
}

func BenchmarkPassBatches(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkResult = passBatches()
	}
}

//BenchmarkPassN/buf=0                 997           1180279 ns/op             128 B/op          2 allocs/op
//BenchmarkPassN/buf=0-2               559           2076635 ns/op             136 B/op          2 allocs/op
//BenchmarkPassN/buf=1                1362            864130 ns/op             144 B/op          2 allocs/op
//BenchmarkPassN/buf=1-2               744           1595163 ns/op             148 B/op          2 allocs/op
//BenchmarkPassN/buf=100              4314            276345 ns/op            1040 B/op          2 allocs/op
//BenchmarkPassN/buf=100-2            2281            513091 ns/op            1052 B/op          2 allocs/op
//BenchmarkPassN/buf=10000                    4554            259933 ns/op           81936 B/op          2 allocs/op
//BenchmarkPassN/buf=10000-2                  2862            414906 ns/op           81936 B/op          2 allocs/op
//BenchmarkPassBatches                       42919             28118 ns/op           89728 B/op        102 allocs/op
//BenchmarkPassBatches-2                     24837             48151 ns/op           89728 B/op        102 allocs/op
