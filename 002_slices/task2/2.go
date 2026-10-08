package main

import "fmt"

func main() {
	arr := [5]float64{1.1, 2.2, 3.3, 4.4, 5.5}

	s := arr[1:4]

	fmt.Println(arr, s)

	s[0] = 8.8
	
	fmt.Println(arr, s)

	arr[2] = 9.9

	fmt.Println(arr, s)
z}
