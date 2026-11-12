package main

import (
	"sync"
	"testing"
)

//go:noinline
func work() {}

// BenchmarkGoroutineStartStop-12           2372120               510.8 ns/op            32 B/op          2 allocs/op
func BenchmarkGoroutineStartStop(t *testing.B) {
	for i := 0; i < t.N; i++ {
		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()
			work()
		}()

		wg.Wait()
	}
}

// BenchmarkFuncCall-12                    904569432                1.318 ns/op           0 B/op          0 allocs/op
func BenchmarkFuncCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		work()
	}
}
