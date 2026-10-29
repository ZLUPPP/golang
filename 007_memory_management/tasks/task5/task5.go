package main

import (
	"bytes"
	"fmt"
	"runtime"
)

func readFile(i int) []byte {
	data := make([]byte, 1<<20) // pretend we read a 1 MB file
	copy(data, fmt.Sprintf("HEADER-%04d", i))
	return data
}

func keepHeadersSlice(n int) [][]byte {
	var headers [][]byte
	for i := 0; i < n; i++ {
		data := readFile(i)
		headers = append(headers, data[:11]) // 11 bytes... or 1 MB?
	}
	return headers
}

func keepHeadersCopy(n int) [][]byte {
	var headers [][]byte
	for i := 0; i < n; i++ {
		data := readFile(i)
		headers = append(headers, bytes.Clone(data[:11]))
	}
	return headers
}

func measure(name string, fn func(int) [][]byte) {
	runtime.GC()

	var before runtime.MemStats
	runtime.ReadMemStats(&before)

	headers := fn(100)

	runtime.GC()

	var after runtime.MemStats
	runtime.ReadMemStats(&after)

	fmt.Printf(
		"%s: %.2f MiB\n",
		name,
		float64(int64(after.HeapAlloc)-int64(before.HeapAlloc))/(1024*1024),
	)

	runtime.KeepAlive(headers)
}

func main() {
	measure("Subslice", keepHeadersSlice)
	measure("Clone", keepHeadersCopy)
}
