package main

import (
	"cmp"
	"fmt"
)

type Score int

type Number interface {
	~int | ~float64
}

func Sum[T Number](nums []T) T {
	var sum T
	for _, n := range nums {
		sum += n
	}
	return sum
}

func Average[T Number](nums []T) (float64, bool) {
	var sum T
	if len(nums) == 0 {
		return 0, false
	}
	for _, n := range nums {
		sum += n
	}
	return float64(sum) / float64(len(nums)), true
}

func MaxOf[T cmp.Ordered](items []T) (T, bool) {
	var maximum T
	if len(items) == 0 {
		return maximum, false
	}
	for _, n := range items {
		if n > maximum {
			maximum = n
		}
	}
	return maximum, true
}

func IndexOf[T comparable](items []T, target T) int {
	for i, item := range items {
		if item == target {
			return i
		}
	}
	return -1
}

func Filter[T any](items []T, keep func(T) bool) []T {
	var filtered []T
	for _, item := range items {
		if keep(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

func main() {
	scores := []Score{60, 75, 90, 85, 40}

	// 1. Sum и Average
	fmt.Println("Sum:", Sum(scores))

	avg, ok := Average(scores)
	fmt.Println("Average:", avg, ok)

	avg, ok = Average([]Score{})
	fmt.Println("Average empty:", avg, ok)

	// 2. MaxOf
	maxInt, ok := MaxOf([]int{3, 8, 2, 10})
	fmt.Println("Max int:", maxInt, ok)

	maxString, ok := MaxOf([]string{"apple", "banana", "zebra"})
	fmt.Println("Max string:", maxString, ok)

	maxScore, ok := MaxOf(scores)
	fmt.Println("Max Score:", maxScore, ok)

	maxInt, ok = MaxOf([]int{})
	fmt.Println("Max empty:", maxInt, ok)

	// 3. IndexOf
	fmt.Println("IndexOf 75:", IndexOf(scores, Score(75)))
	fmt.Println("IndexOf 100:", IndexOf(scores, Score(100)))

	// 4. Filter
	passed := Filter(scores, func(s Score) bool {
		return s >= 75
	})
	fmt.Println("Passed scores:", passed)
}
