package main

import (
	"fmt"
	"runtime"
	"time"
)

type Record struct {
	ID   int
	Data [256]byte
}

var (
	live []*Record
	sum  int
)

//go:noinline
func consume(r *Record) int {
	return r.ID + int(r.Data[0])
}

func main() {
	start := time.Now()

	// Объекты, которые остаются в памяти
	for i := 0; i < 80_000; i++ {
		live = append(live, &Record{ID: i})
	}

	// Временные объекты
	for i := 0; i < 10_000_000; i++ {
		r := &Record{ID: i}
		sum += consume(r)
	}

	elapsed := time.Since(start)

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Time: %v\n", elapsed)
	fmt.Printf("GC cycles: %d\n", m.NumGC)
	fmt.Printf("HeapAlloc: %.2f MiB\n",
		float64(m.HeapAlloc)/1024/1024)
	fmt.Printf("Sys: %.2f MiB\n",
		float64(m.Sys)/1024/1024)
	fmt.Printf("Total GC pause: %.2f ms\n",
		float64(m.PauseTotalNs)/1e6)
	fmt.Printf("Sum: %d\n", sum)

	runtime.KeepAlive(live)
}
