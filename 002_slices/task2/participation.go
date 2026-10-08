package main

import "fmt"

func edit(slice []float64) {
	slice[0] = 9.99	
}


func main() {
	arr := [5]float64{1.1, 2.2, 3.3, 4.4, 5.5}
	slice := arr[1:4]

	fmt.Println(arr, slice)

	edit(slice)

	fmt.Println("after edit\n", arr, slice)
}
