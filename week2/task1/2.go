
package main

import "fmt"

func main() {
	var a [5]float64
	// 2.
	for i := 0; i < 3; i++ {
		a[i] = float64(i) + 1.5
	}
	fmt.Println(a)
}
