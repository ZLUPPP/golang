package main

const N = 10_000

func passN(size int) int {
	ch := make(chan int, size)

	go func() {
		for i := 0; i < N; i++ {
			ch <- i
		}
		close(ch)
	}()

	sum := 0

	for v := range ch {
		sum += v
	}

	return sum
}

func passBatches() int {
	ch := make(chan []int)

	go func() {
		defer close(ch)
		for start := 0; start < N; start += 100 {
			batch := make([]int, 0, 100)
			for i := start; i < start+100 && i < N; i++ {
				batch = append(batch, i)
			}
			ch <- batch
		}
	}()

	sum := 0
	for batch := range ch {
		for _, v := range batch {
			sum += v
		}
	}
	return sum
}
