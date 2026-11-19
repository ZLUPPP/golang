package main

import (
	"fmt"
	"time"
)

func run(name string, ch chan int) {
	start := time.Now()
	since := func() time.Duration { return time.Since(start).Round(100 * time.Millisecond) }
	done := make(chan struct{})
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
			fmt.Printf("%s: sent %d at %v\n", name, i, since())
		}
		close(done)
	}()
	time.Sleep(300 * time.Millisecond) // the receiver is busy at first
	for i := 0; i < 3; i++ {
		<-ch
		time.Sleep(100 * time.Millisecond) // and slow
	}
	<-done
}

func main() {
	run("unbuffered", make(chan int))
	run("buffered(3)", make(chan int, 10))

	ch := make(chan string, 1)
	ch <- "a"
	ch <- "b"
	fmt.Println(<-ch, <-ch)
}

// 0
//unbuffered: sent 1 at 300ms
//unbuffered: sent 2 at 400ms
//unbuffered: sent 3 at 500ms
//buffered(3): sent 1 at 300ms
//buffered(3): sent 2 at 400ms
//buffered(3): sent 3 at 500ms

// 1
//unbuffered: sent 1 at 300ms
//unbuffered: sent 2 at 400ms
//unbuffered: sent 3 at 500ms
//buffered(3): sent 1 at 0s
//buffered(3): sent 2 at 300ms
//buffered(3): sent 3 at 400ms

// 2
//unbuffered: sent 1 at 300ms
//unbuffered: sent 2 at 400ms
//unbuffered: sent 3 at 500ms
//buffered(3): sent 1 at 0s
//buffered(3): sent 2 at 0s
//buffered(3): sent 3 at 300ms

// 3
//unbuffered: sent 1 at 300ms
//unbuffered: sent 2 at 400ms
//unbuffered: sent 3 at 500ms
//buffered(3): sent 1 at 0s
//buffered(3): sent 2 at 0s
//buffered(3): sent 3 at 0s

// 10
//unbuffered: sent 1 at 300ms
//unbuffered: sent 2 at 400ms
//unbuffered: sent 3 at 500ms
//buffered(3): sent 1 at 0s
//buffered(3): sent 2 at 0s
//buffered(3): sent 3 at 0s
