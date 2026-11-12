package main

import (
	"runtime"
	"testing"
	"time"
)

func TestFetchNoLeak(t *testing.T) {
	before := runtime.NumGoroutine()

	for i := 0; i < 100; i++ {
		fetchLeaky(5 * time.Millisecond)
	}

	time.Sleep(100 * time.Millisecond)

	after := runtime.NumGoroutine()

	if after > before+2 {
		t.Errorf(
			"goroutine leak: before=%d, after=%d",
			before, after,
		)
	}
}
