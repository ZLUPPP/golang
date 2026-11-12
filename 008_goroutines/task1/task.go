package main

import (
	"sync"
	"time"
)

func ioWork() {
	defer wg.Done()
	time.Sleep(time.Second)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go ioWork()
	}

	wg.Wait()
}
