package main

import (
	"sync"
	"testing"
)

// BenchmarkErrorReturn-12         1000000000               1.104 ns/op           0 B/op          0 allocs/op
func BenchmarkErrorReturn(t *testing.B) {

	for i := 0; i < t.N; i++ {
		checkErr(-1)
	}
}

// BenchmarkPanicRecover-12         6800496               183.5 ns/op             0 B/op          0 allocs/op
func BenchmarkPanicRecover(t *testing.B) {

	for i := 0; i < t.N; i++ {
		safeCheck(-1)
	}
}

// BenchmarkIncNoDefer-12          291820915                3.863 ns/op           0 B/op          0 allocs/op
func BenchmarkIncNoDefer(t *testing.B) {
	c := Counter{mu: sync.Mutex{}, n: 0}

	t.ResetTimer()

	for i := 0; i < t.N; i++ {
		c.IncNoDefer()
	}
}

// BenchmarkIncDefer-12            277659836                4.328 ns/op           0 B/op          0 allocs/op
func BenchmarkIncDefer(t *testing.B) {
	c := Counter{mu: sync.Mutex{}, n: 0}

	t.ResetTimer()

	for i := 0; i < t.N; i++ {
		c.IncDefer()
	}
}
