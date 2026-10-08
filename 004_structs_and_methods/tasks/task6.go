package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return (r.Width + r.Height) * 2 }

func (r *Rectangle) Scale(k float64) {
	r.Width *= k
	r.Height *= k
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 { return c.Radius * c.Radius * 3.14 }
func (c *Circle) Perimeter() float64 { return c.Radius * 2 * 3.14 }


func main() {
	// Circle{Radius: 4} | Circle does not implement Shape (method Area has pointer receiver)

	shapes := []Shape{Rectangle{Width: 2, Height: 3}, &Circle{Radius: 4}}

	for i := range shapes {
		fmt.Println(shapes[i].Area())
		fmt.Println(shapes[i].Perimeter())
	}

	r := Rectangle{3, 4}
	f := r.Area // Links with the Rectangle{3, 4} structure (cause value method copy struct)
	r.Width = 10
	fmt.Println(f(), r.Area()) // 12 40

	c := Circle{8}
	f2 := c.Area
	c.Radius = 10
	fmt.Println(f2(), c.Area()) // 314 314

}