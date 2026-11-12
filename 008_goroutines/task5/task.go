package main

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"time"
)

func slowQuery() int {
	time.Sleep(50 * time.Millisecond)
	return 42
}

func fetchLeaky(timeout time.Duration) (int, error) {
	ch := make(chan int, 1)

	go func() {
		ch <- slowQuery()
	}()

	select {
	case v := <-ch:
		return v, nil

	case <-time.After(timeout):
		return 0, errors.New("timeout")
	}
}

func startTickerLeaky(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return

			case <-time.After(10 * time.Millisecond):
				// do sth
			}
		}
	}()
}

func main() {
	fmt.Println("Before:", runtime.NumGoroutine())

	for i := 0; i < 100; i++ {
		fetchLeaky(5 * time.Millisecond)

	}
	time.Sleep(100 * time.Millisecond)

	fmt.Println("After fetchLeaky:", runtime.NumGoroutine())

	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 100; i++ {
		startTickerLeaky(ctx)

	}

	fmt.Println("Before cancel:", runtime.NumGoroutine())

	cancel()

	time.Sleep(100 * time.Millisecond)

	fmt.Println("After startTickerLeaky:", runtime.NumGoroutine())

	//pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
	//pprof.Lookup("goroutine").WriteTo(os.Stdout, 2)

}
