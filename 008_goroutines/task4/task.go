package main

import (
	"fmt"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

func busy(d time.Duration) {
	deadline := time.Now().Add(d)

	for time.Now().Before(deadline) {
	}
}

func run(name string, n int, fn func()) {
	var wg sync.WaitGroup

	start := time.Now()

	for i := 0; i < n; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			fn()
		}()
	}

	wg.Wait()

	fmt.Printf(
		"%s: threads=%d, duration=%v\n",
		name,
		pprof.Lookup("threadcreate").Count(),
		time.Since(start),
	)
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			busy(2 * time.Second)
		}()
	}
	wg.Wait()

	runtime.GOMAXPROCS(1)

	go func() {
		for {
		}
	}()

	time.Sleep(10 * time.Millisecond)
	fmt.Println("main is still alive")

	fmt.Println(
		"Start threads:",
		pprof.Lookup("threadcreate").Count(),
	)

	run("time.Sleep", 100, func() {
		time.Sleep(200 * time.Millisecond)
	})

	run("syscall.Nanosleep", 100, func() {
		//ts := syscall.NsecToTimespec(
		//	int64(200 * time.Millisecond),
		//)

		//syscall.NanoSleep(&ts, nil)
	})
}
