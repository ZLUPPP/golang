package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	const N = 1_000_000

	var before, after runtime.MemStats
	var wg sync.WaitGroup

	stop := make(chan struct{})

	runtime.ReadMemStats(&before)
	start := time.Now()

	for i := 0; i < N; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			<-stop
		}()
	}

	runtime.ReadMemStats(&after)
	fmt.Println("Goroutines:", runtime.NumGoroutine()) // 100001
	fmt.Println("Time:", time.Since(start))            // 293.0983ms   /     3.7015942s

	diff := after.StackSys - before.StackSys
	fmt.Printf("Stack per goroutine: %.2f KB\n", // 8.00 KB
		float64(diff)/float64(N)/1024)

	close(stop)
	wg.Wait()
}
