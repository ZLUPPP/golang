package main

import "testing"

func TestCountPar(t *testing.T) {
	limit := 200_000

	expected := CountSeq(limit)

	for _, parts := range []int{2, 8, 1000, 200000} {
		actual := CountPar(limit, parts)

		if actual != expected {
			t.Errorf(
				"parts=%d: expected=%d, got=%d",
				parts, expected, actual,
			)
		}
	}
}

// BenchmarkCountSeq                    100          11378409 ns/op               0 B/op          0 allocs/op
// BenchmarkCountSeq-2                  100          11319511 ns/op               0 B/op          0 allocs/op
// BenchmarkCountSeq-4                  100          11306087 ns/op               0 B/op          0 allocs/op
// BenchmarkCountSeq-8                  100          11305793 ns/op               0 B/op          0 allocs/op
func BenchmarkCountSeq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountSeq(200_000)
	}
}

// BenchmarkCountPar2                   100          11254668 ns/op             225 B/op          6 allocs/op
// BenchmarkCountPar2-2                 171           7018594 ns/op             224 B/op          6 allocs/op
// BenchmarkCountPar2-4                 169           7027651 ns/op             241 B/op          6 allocs/op
// BenchmarkCountPar2-8                 169           7033873 ns/op             249 B/op          6 allocs/op
func BenchmarkCountPar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountPar(200_000, 2)
	}
}

// BenchmarkCountPar8                   100          11250506 ns/op             848 B/op         18 allocs/op
// BenchmarkCountPar8-2                 190           6234711 ns/op             848 B/op         18 allocs/op
// BenchmarkCountPar8-4                 331           3542509 ns/op             848 B/op         18 allocs/op
// BenchmarkCountPar8-8                 397           3108372 ns/op             945 B/op         18 allocs/op
func BenchmarkCountPar8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountPar(200_000, 8)
	}
}

// BenchmarkCountPar200000               32         133798259 ns/op        21937945 B/op     402181 allocs/op
// BenchmarkCountPar200000-2             31          45258316 ns/op        20805666 B/op     400002 allocs/op
// BenchmarkCountPar200000-4             36          35523556 ns/op        20805669 B/op     400002 allocs/op
// BenchmarkCountPar200000-8             36          35771622 ns/op        20805830 B/op     400002 allocs/op
func BenchmarkCountPar200000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountPar(200_000, 200_000)
	}
}
