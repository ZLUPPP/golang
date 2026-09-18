package main

import "fmt"

func main() {
	cars := []string{"Ferrari", "Honda", "Ford", "BYD"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))
}
