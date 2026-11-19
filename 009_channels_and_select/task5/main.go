package main

import (
	"fmt"
	"time"
)

func slow() string {
	time.Sleep(200 * time.Millisecond)
	return "data"
}

func fetchWithTimeout(timeout time.Duration) (string, error) {
	result := make(chan string, 1)

	go func() {
		result <- slow()
	}()

	select {
	case res := <-result:
		return res, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timeout")
	}

}

func merge(a, b <-chan int) <-chan int {
	out := make(chan int)
	zeros := 0
	go func() {
		defer close(out)
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					zeros++
					a = nil // a nil channel blocks forever: this case is now off
					continue
				}
				out <- v
			case v, ok := <-b:
				if !ok {
					b = nil
					fmt.Println("useless reads from closed a:", zeros)
					continue
				}
				out <- v
			}
		}
	}()
	return out
}

func main() {
	a := make(chan string, 1000)
	b := make(chan string, 1000)

	for i := 0; i < 1000; i++ {
		a <- "a"
		b <- "b"
	}

	count := make(map[string]int)
	for i := 0; i < 1000; i++ {
		select {
		case v := <-a:
			count[v]++
			continue
		default:
		}

		select {
		case v := <-a:
			count[v]++
		case v := <-b:
			count[v]++
		}
	}

	fmt.Println(count)

	val, err := fetchWithTimeout(200 * time.Millisecond)
	fmt.Println(val, err)

	a1, b1 := make(chan int), make(chan int)

	go func() {
		defer close(a1)
		a1 <- 1
		a1 <- 2
	}()

	go func() {
		defer close(b1)
		b1 <- 100
		time.Sleep(20 * time.Millisecond)
		b1 <- 101
		b1 <- 102
	}()

	for v := range merge(a1, b1) {
		fmt.Println(v)
	}

}
