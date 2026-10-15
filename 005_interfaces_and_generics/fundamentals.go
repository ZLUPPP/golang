package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

// PART 1

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct{ W, H float64 }

func (r Rect) Area() float64      { return r.W * r.H }
func (r Rect) Perimeter() float64 { return 2 * (r.W + r.H) }

type Circle struct{ R float64 }

func (c *Circle) Area() float64      { return math.Pi * c.R * c.R }
func (c *Circle) Perimeter() float64 { return 2 * math.Pi * c.R }

func printInfo(s Shape) {
	fmt.Printf("%-14T area=%6.2f perimeter=%6.2f\n", s, s.Area(), s.Perimeter())
}

func writeReport(w io.Writer, name string, score int) {
	fmt.Fprintf(w, "%s: %d points\n", name, score)
}

type Logger interface{ Log(msg string) }
type Saver interface{ Save(data string) error }

type LogSaver interface {
	Logger
	Saver
}

// PART 2

type Celsius float64

func (c Celsius) String() string { return fmt.Sprintf("%.1f°C", float64(c)) }

func describe(v any) string {
	switch x := v.(type) {
	case nil:
		return "<nil>"
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("integer: %v", x)
	case string:
		return fmt.Sprintf("string of %d bytes", len(x))
	case []int:
		return fmt.Sprintf("slice of %d ints", len(x))
	case error:
		return "error " + x.Error()
	case fmt.Stringer:
		return "Stringer " + x.String()
	default:
		return fmt.Sprintf("unknown type: %T", x)
	}
}

type Timeout struct{ Sec int }

func (t Timeout) Error() string  { return fmt.Sprintf("timeout after %ds", t.Sec) }
func (t Timeout) String() string { return "Timeout{" + fmt.Sprint(t.Sec) + "}" }

func kind(v any) string {
	switch v.(type) {
	case fmt.Stringer:
		return "Stringer"
	case error:
		return "error"
	case Timeout:
		return "Timeout"
	}
	return "other"
}

// PART 3

type Shape_ interface{ Area() float64 }

type Square struct{ Side float64 }

func (s *Square) Area() float64 {
	if s == nil {
		return 0
	}
	return s.Side * s.Side
}

// PART 4: Generics and constraints

type Number interface {
	~int | ~int64 | ~float64
}

func Sum[T Number](nums []T) T {
	var s T
	for _, n := range nums {
		s += n
	}
	return s
}

func main() {
	// PART 1
	var _ Shape = Rect{}
	var _ Shape = (*Circle)(nil)

	shapes := []Shape{Rect{3, 4}, &Circle{1}}
	for _, s := range shapes {
		printInfo(s)
	}

	var s Shape
	fmt.Printf("%v %T\n", s, s) // <nil> <nil>
	s = Rect{1, 2}
	fmt.Printf("%v %T\n", s, s) // {1 2} main.Rect
	s = &Circle{2}
	fmt.Printf("%v %T\n", s, s) // &{2} *main.Circle

	var x any
	x = 42
	x = "hello"
	x = []int{1, 2}

	fmt.Println(x)

	writeReport(os.Stdout, "Aruzhan", 95) // print to the terminal
	var sb strings.Builder                // remember Week 3
	writeReport(&sb, "Dias", 80)          // write into a string
	fmt.Print(sb.String())

	// PART 2

	var i any = "Go"

	z := i.(string)
	fmt.Println(z)

	v, ok := i.(int)
	fmt.Println(v, ok)

	//f := i.(float64) // panic
	//fmt.Println(f)

	var t any = Celsius(36.66)
	if st, ok := t.(fmt.Stringer); ok {
		fmt.Println(st)
	}

	fmt.Println(describe(15))

	fmt.Println(kind(Timeout{5}))
	var e error = Timeout{3}
	fmt.Println(e) // String()

	// Part 3

	var l Shape_          // nil interface
	fmt.Println(l == nil) // true

	var sq *Square        // nil pointer
	l = sq                // interface holds (*Square, nil)
	fmt.Println(l == nil) // false
	fmt.Println(l.Area()) // 0
	l = nil
	//fmt.Println(l.Area()) // panic

	// Part 4: Generics and constraints


	fmt.Println(Sum([]int{1, 2, 3})) // 6 — T inferred as int
	fmt.Println(Sum([]float64{1.5, 2.5})) // 4 — T inferred as float64
	fmt.Println(Sum[int64]([]int64{10, 20})) // 30 — T written explicitly

}
