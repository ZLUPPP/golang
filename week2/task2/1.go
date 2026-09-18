package main

import "fmt"

func main() {
	var s1 []float64

	s2 := []float64{1.1, 2.2, 3.3}

	s3 := make([]float64, 3, 5)

	arr := [5]float64{1.1, 2.2, 3.3, 4.4, 5.5}
	s4 := arr[1:4]

	fmt.Println(s1, s2, s3, s4)
}
