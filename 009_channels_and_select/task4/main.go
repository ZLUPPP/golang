package main

import (
	"runtime"

	"golang.org/x/tour/tree"
)
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var visit func(*tree.Tree)

	visit = func(node *tree.Tree) {
		if node == nil {
			return
		}
		visit(node.Left)
		ch <- node.Value
		visit(node.Right)
	}

	visit(t)

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	var values1, values2 []int

	for v := range ch1 {
		values1 = append(values1, v)
	}
	for v := range ch2 {
		values2 = append(values2, v)
	}

	if len(values1) != len(values2) {
		return false
	}

	for i := range values1 {
		if values1[i] != values2[i] {
			return false
		}
	}

	return true

}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))

	before := runtime.NumGoroutine()

	fmt.Println(Same(tree.New(1), tree.New(2)))

	after := runtime.NumGoroutine()
	fmt.Printf("goroutines: before=%d, after=%d\n", before, after)

}
