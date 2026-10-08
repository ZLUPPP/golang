package main

import "testing"

// Report is deliberately large: 1024 * 8 bytes = 8 KB.
type Report struct {
	Title string
	Values [1024]float64
}

//go:noinline
func (r Report) FirstValue() float64 { return r.Values[0] }

//go:noinline
func (r *Report) FirstPointer() float64 { return r.Values[0] }

type Point struct { X, Y float64 }

//go:noinline
func (p Point) SecondValue() float64 { return p.Y }

//go:noinline
func (p *Point) SecondPointer() float64 { return p.Y }


// BenchmarkFirstValue-12          13150670                93.46 ns/op            0 B/op          0 allocs/op
func BenchmarkFirstValue(b *testing.B) {
	rep := Report{Title: "Big Report"}

	b.ResetTimer()
	for range b.N {
		_ = rep.FirstValue()
	}
}

// BenchmarkFirstPointer-12        882906704                1.398 ns/op           0 B/op          0 allocs/op
func BenchmarkFirstPointer(b *testing.B) {
	rep := Report{Title: "Big Report"}

	b.ResetTimer()
	for range b.N {
		_ = rep.FirstPointer()
	}
}

// BenchmarkSecondValue-12         875861902                1.397 ns/op           0 B/op          0 allocs/op
func BenchmarkSecondValue(b *testing.B) {
	p := Point{X: 1.1, Y: 2.2}

	b.ResetTimer()
	for range b.N {
		_ = p.SecondValue()
	}
}

// BenchmarkSecondPointer-12       1000000000               1.133 ns/op           0 B/op          0 allocs/op
func BenchmarkSecondPointer(b *testing.B) {
	p := Point{X: 1.1, Y: 2.2}

	b.ResetTimer()
	for range b.N {
		_ = p.SecondPointer()
	}
}