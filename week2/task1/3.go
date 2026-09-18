package main

import "fmt"

func main() {
	// 3.
	a := [2][2]int{{1, 2}, {3, 4}}

	for i, row := range a {
		for j := range row {
			fmt.Println(a[i][j])
		}
	}
}
