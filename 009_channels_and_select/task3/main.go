package main

import (
	"fmt"
	"sync"
)

func collect(workers int) <-chan int {
	results := make(chan int)
	var wg sync.WaitGroup
	for w := 1; w <= workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 3; i++ {
				results <- w*10 + i
			}
		}()
	}
	go func() {
		wg.Wait()      // all writers are finished...
		close(results) // ...so now it's safe to close
	}()
	return results
}

func generator(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for _, num := range nums {
			out <- num
		}
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for v := range in {
			out <- v * v
		}
	}()

	return out
}

func filterEven(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for v := range in {
			if v%2 == 0 {
				out <- v
			}
		}
	}()

	return out
}

func Sum(in <-chan int) int {
	sum := 0
	for v := range in {
		sum += v
	}

	return sum
}

func main() {
	//results := make(chan int)
	//var wg sync.WaitGroup
	//for w := 1; w <= 3; w++ {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		for i := 0; i < 3; i++ {
	//			results <- w*10 + i
	//		}
	//		close(results) // "I'm done" — but the others are not!
	//	}()
	//}
	//sum := 0
	//for v := range results {
	//	sum += v
	//}
	//fmt.Println("sum:", sum)
	//wg.Wait()

	ch := collect(3)
	sum := 0
	for v := range ch {
		sum += v
	}

	fmt.Println(sum)

	ch2 := generator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	ch3 := square(ch2)
	ch4 := filterEven(ch3)
	fmt.Println(Sum(ch4))
}
