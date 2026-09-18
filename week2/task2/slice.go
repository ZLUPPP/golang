package main

import "fmt"

func main() {
	src := []int{1, 2, 3, 4}
	dst := make([]int, 2)
	fmt.Println(dst)
	// Copy first two elements
	copy(dst, src)
	fmt.Println("Copied:", dst, "Elements copied:")
	// Modify source
	src[0] = 99
	fmt.Println("Source after mod:", src)
	fmt.Println("Dest after mod:", dst)
}
