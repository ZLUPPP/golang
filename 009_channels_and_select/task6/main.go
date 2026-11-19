package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

type Result struct {
	Word  string
	Upper string
}

func producer(words []string, done <-chan struct{}) <-chan string {
	jobs := make(chan string)

	go func() {
		defer close(jobs)
		for _, word := range words {
			select {
			case jobs <- word:
			case <-done:
				return
			}
		}
	}()

	return jobs
}

func startWorkers(n int, jobs <-chan string, done <-chan struct{}) <-chan Result {
	results := make(chan Result)
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				var word string
				var ok bool

				select {
				case word, ok = <-jobs:
					if !ok {
						return
					}
				case <-done:
					return
				}

				result := Result{Word: word, Upper: strings.ToUpper(word)}

				select {
				case results <- result:
				case <-done:
					return
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

var ErrQueueFull = fmt.Errorf("queue full")

func Submit(queue chan<- string, job string, timeout time.Duration) error {
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	select {
	case queue <- job:
		return nil
	case <-timer.C:
		return ErrQueueFull
	}
}

func main() {
	words := []string{"one", "two", "three", "four", "five", "six", "seven"}

	done := make(chan struct{})
	results := startWorkers(2, producer(words, done), done)
	for result := range results {
		fmt.Println(result)
	}
	fmt.Println("after full run:", runtime.NumGoroutine())

	done = make(chan struct{})
	results = startWorkers(2, producer(words, done), done)
	for i := 0; i < 3; i++ {
		fmt.Println(<-results)
	}
	close(done)

	for range results {
	}
	fmt.Println("after cancel:", runtime.NumGoroutine())

	queue := make(chan string, 1)
	fmt.Println("first submit:", Submit(queue, "first", time.Millisecond))
	fmt.Println("second submit:", Submit(queue, "second", 10*time.Millisecond))
}
