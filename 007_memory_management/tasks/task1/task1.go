package main

import "fmt"

func depth(n int) int {
	var pad [64]byte
	pad[n%len(pad)] = byte(n)
	if n == 0 {
		return 0
	}
	return depth(n-1) + 1
}

func forever(n int) int {
	return forever(n+1) + 1
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	forever(0)
}
