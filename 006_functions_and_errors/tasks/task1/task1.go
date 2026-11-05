package main

import "fmt"

func makeAccumulator() func(int) int {
	total := 0
	return func(i int) int {
		total += i
		return total
	}
}

func memorize(f func(int) int) func(int) int {
	cache := make(map[int]int)

	return func(i int) int {
		value, exists := cache[i]
		if exists {
			return value
		}

		result := f(i)
		cache[i] = result
		return result
	}

}

func main() {
	acc := makeAccumulator()

	fmt.Println("Accumulator:", acc(12))
	fmt.Println("Accumulator:", acc(2))
	fmt.Println("Accumulator:", acc(-10))

	cnt := 0

	slow := func(n int) int {
		cnt++
		return n * n
	}

	mem := memorize(slow)
	fmt.Println("Memorize:", mem(5))
	fmt.Println("Memorize:", mem(5))
	fmt.Println("Memorize:", mem(5))
	fmt.Println("Memorize:", mem(10))
	fmt.Println("Memorize:", mem(10))

	fmt.Println("counter: ", cnt)

	var funcs []func()
	for i := 0; i < 3; i++ {
		funcs = append(funcs, func() { fmt.Print(i, " ") })
	}
	for _, f := range funcs {
		f()
	}
}
