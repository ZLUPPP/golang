package main

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Report struct {
	Data [64 * 1024]byte
}

func buildReport() *Report {
	time.Sleep(20 * time.Millisecond)
	return &Report{}
}

func buildReportCtx(ctx context.Context) (*Report, error) {
	select {
	case <-time.After(20 * time.Millisecond):
		return &Report{}, nil

	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func handle(timeout time.Duration) error {
	ch := make(chan *Report, 1)

	go func() { ch <- buildReport() }()

	select {
	case r := <-ch:
		_ = r
		return nil

	case <-time.After(timeout):
		return errors.New("timeout")
	}
}

func handleContext(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		timeout,
	)
	defer cancel()

	ch := make(chan *Report, 1)

	go func() {
		report, err := buildReportCtx(ctx)
		if err != nil {
			return
		}

		ch <- report
	}()

	select {
	case <-ch:
		return nil

	case <-ctx.Done():
		return ctx.Err()
	}
}

func printStats(label string) {
	runtime.GC()

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	fmt.Printf(
		"%s: goroutines=%d heap=%d MB\n",
		label,
		runtime.NumGoroutine(),
		mem.HeapAlloc/1024/1024,
	)
}

func runRound() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			handleContext(5 * time.Millisecond)
		}()
	}

	wg.Wait()

	time.Sleep(50 * time.Millisecond)
}

func main() {
	printStats("Start")

	for i := 1; i <= 3; i++ {
		runRound()

		printStats(fmt.Sprintf("After round %d", i))
	}
}
