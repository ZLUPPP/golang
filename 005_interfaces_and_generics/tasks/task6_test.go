package main

import "testing"

var benchmarkResult int

func makeNumbers() []int {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i
	}
	return nums
}

// BenchmarkSumInts-12                      4651156               268.6 ns/op             0 B/op          0 allocs/op
func BenchmarkSumInts(b *testing.B) {
	nums := makeNumbers()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		benchmarkResult = SumInts(nums)
	}
}

// BenchmarkSumGeneric-12                   4404981               272.1 ns/op             0 B/op          0 allocs/op
func BenchmarkSumGeneric(b *testing.B) {
	nums := makeNumbers()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		benchmarkResult = Sum_(nums)
	}
}

// BenchmarkSumAny-12                       2642337               478.7 ns/op             0 B/op          0 allocs/op
func BenchmarkSumAny(b *testing.B) {
	nums := makeNumbers()
	boxed := ToAny(nums)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		benchmarkResult = SumAny(boxed)
	}
}

// BenchmarkSumAnyWithBoxing-12              109557             10574 ns/op           22336 B/op        745 allocs/op
func BenchmarkSumAnyWithBoxing(b *testing.B) {
	nums := makeNumbers()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		benchmarkResult = SumAny(ToAny(nums))
	}
}
