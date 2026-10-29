package task4

import "testing"

// BenchmarkNewPointValue-12       900871818                1.346 ns/op           0 B/op          0 allocs/op
func BenchmarkNewPointValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPointValue(0, 0)
	}
}

// BenchmarkNewPointPointer-12     100000000               10.31 ns/op           16 B/op          1 allocs/op
func BenchmarkNewPointPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPointPointer(0, 0)
	}
}

// BenchmarkCollectNoCap-12          231967              4876 ns/op           25208 B/op         12 allocs/op
func BenchmarkCollectNoCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CollectNoCap(1000)
	}
}

// BenchmarkCollectWithCap-12       3496435               350.8 ns/op             0 B/op          0 allocs/op
func BenchmarkCollectWithCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CollectWithCap(1000)
	}
}
