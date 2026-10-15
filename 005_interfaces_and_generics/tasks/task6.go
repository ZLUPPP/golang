package main

type Number_ interface {
	~int | ~float64
}

func SumInts(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func Sum_[T Number_](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

func SumAny(nums []any) int {
	total := 0
	for _, n := range nums {
		total += n.(int)
	}
	return total
}

func ToAny(nums []int) []any {
	result := make([]any, len(nums))
	for i, n := range nums {
		result[i] = n
	}
	return result
}
