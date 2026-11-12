package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func say(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, i)
	}
}
func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		say("goroutine")
	}()

	wg.Wait()

	say("main")

	var counter atomic.Int64
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Add(1)
		}()
	}
	wg.Wait()
	fmt.Println("counter:", counter.Load())

	counter_list := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter_list[i] = 1
		}()
	}
	wg.Wait()
	sum := 0
	for _, v := range counter_list {
		sum += v
	}
	fmt.Println("counter_list:", sum)

	wg.Add(1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recover:", r)
			}
		}()
		defer wg.Done()

		panic("╰（‵□′）╯")
	}()

	wg.Wait()

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("worker", i, "done")
		}()
	}
	wg.Wait() // 3. block until the counter is 0
	fmt.Println("all workers finished")
}
