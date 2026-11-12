package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i <= n/i; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func countPrimes(from, to int) int {
	count := 0

	for i := from; i <= to; i++ {
		if isPrime(i) {
			count++
		}
	}

	return count
}

func cpuWork(workers int) (int, time.Duration) {
	const N = 6_000_000

	start := time.Now()

	results := make([]int, workers)

	var wg sync.WaitGroup

	chunk := N / workers

	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			results[w] = countPrimes(w*chunk, (w+1)*chunk)
		}()
	}
	wg.Wait()

	total := 0
	for _, result := range results {
		total += result
	}

	return total, time.Since(start)
}

// I/O-bound: each "request" just waits 100 ms
func ioWork(workers int) time.Duration {
	start := time.Now()
	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond) // like waiting for a database
		}()
	}
	wg.Wait()
	return time.Since(start)
}

func main() {
	fmt.Println("Logical CPUs:", runtime.NumCPU())
	fmt.Println("Used threads:", runtime.GOMAXPROCS(0))

	_, duration := cpuWork(1)

	fmt.Println("cpuWork:", duration)
	fmt.Println("ioWork:", ioWork(2*runtime.NumCPU()))

	// WORKERS         CPU, default ( 12 )     CPU, GOMAXPROCS=1	IO, default ( 12 )		IO, GOMAXPROCS=1
	// 1               1.3169951s			   1.3124794s			100.6221ms				100.7584ms
	// 2			   825.7736ms			   1.3124929s			100.0028ms				101.0339ms
	// 4			   457.4484ms			   1.3081534s			100.7535ms				101.4413ms
	// NumCPU		   236.5803ms			   1.3086183s			101.7377ms				101.0764ms
	// 2 * NumCPU	   249.6991ms			   1.3158021s			100.692ms				101.0764ms
}
