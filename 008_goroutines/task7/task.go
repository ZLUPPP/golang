package main

import (
	"fmt"
	"sync"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func CountSeq(n int) int {
	count := 0

	for i := 2; i <= n; i++ {
		if isPrime(i) {
			count++
		}
	}

	return count
}

func CountPar(n int, parts int) int {
	if parts <= 0 {
		return 0
	}

	var wg sync.WaitGroup

	result := make([]int, parts)
	chunk := (n + parts - 1) / parts

	for i := 0; i < parts; i++ {
		start := i * chunk
		end := start + chunk

		if end > n {
			end = n
		}

		wg.Add(1)

		go func(index, start, end int) {
			defer wg.Done()

			count := 0

			for j := start; j < end; j++ {
				if isPrime(j) {
					count++
				}
			}

			result[index] = count
		}(i, start, end)
	}

	wg.Wait()

	total := 0

	for _, count := range result {
		total += count
	}

	return total
}

func main() {
	limit := 200_000

	fmt.Println("Sequential:", CountSeq(limit))
	fmt.Println("Parallel 2:", CountPar(limit, 2))
	fmt.Println("Parallel 8:", CountPar(limit, 8))
	fmt.Println("Parallel 1000:", CountPar(limit, 1000))
}
