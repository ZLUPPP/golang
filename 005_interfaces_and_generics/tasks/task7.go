package main

import (
	"fmt"
	"strings"
)

type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Push(value T) *List[T] {
	return &List[T]{
		val:  value,
		next: l,
	}
}

func (l *List[T]) Len() int {
	count := 0

	for current := l; current != nil; current = current.next {
		count++
	}

	return count
}

func (l *List[T]) String() string {
	var parts []string

	for current := l; current != nil; current = current.next {
		parts = append(parts, fmt.Sprint(current.val))
	}

	return "[" + strings.Join(parts, ", ") + "]"
}

func main() {
	var numbers *List[int]

	numbers = numbers.Push(10)
	numbers = numbers.Push(20)
	numbers = numbers.Push(30)

	fmt.Println("List:", numbers)
	fmt.Println("Length:", numbers.Len())

	var names *List[string]

	names = names.Push("Alice")
	names = names.Push("Bob")

	fmt.Println("Names:", names)
	fmt.Println("Length:", names.Len())

	var empty *List[int]
	fmt.Println("Empty:", empty.String())
	fmt.Println("Empty length:", empty.Len())
}
