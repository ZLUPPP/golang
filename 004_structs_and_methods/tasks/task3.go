package main

import "fmt"

type Counter struct{ n int }

func (c Counter) Get() int { return c.n }
func (c *Counter) Inc() { c.n++ }
func NewCounter() Counter { return Counter{} }



func main() {
	var c Counter

	c.Inc() // Compiles ✅
	fmt.Println(c.n)
	(&c).Inc() // Compiles ✅
	fmt.Println(c.n)
	fmt.Println(Counter{}.Get()) // Compiles ✅
	// Counter{}.Inc() // Compile error ❌ | Do not have fixed space in memory 
	// NewCounter().Inc() // Compile error ❌ | Do not have fixed space in memory

	m := map[string]Counter{"a": {}}
	m["a"].Get() // Compiles ✅
	// m["a"].Inc() // Compile error ❌ | unaddressable element

	s := []Counter{{}}
	s[0].Inc() // Compiles ✅ | addressable in array

	pm := map[string]*Counter{"a": {}}
	pm["a"].Inc() // Compiles ✅ | We store address, not structure
}