package task7

import "testing"

func TestRender(t *testing.T) {
	ids := make([]int, 200)

	for i := range ids {
		ids[i] = i
	}

	newResult := RenderNew(ids)
	poolResult := RenderPool(ids)

	if newResult != poolResult {
		t.Fatalf(
			"RenderNew=%d, RenderPool=%d",
			newResult,
			poolResult,
		)
	}
}

var result int

// BenchmarkRenderPool-12            382746              3316 ns/op             320 B/op        100 allocs/op
func BenchmarkRenderPool(b *testing.B) {
	ids := make([]int, 200)

	for i := range ids {
		ids[i] = i
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result = RenderPool(ids)
	}
}

// BenchmarkRenderNew-12             350703              3582 ns/op            4416 B/op        101 allocs/op
func BenchmarkRenderNew(b *testing.B) {
	ids := make([]int, 200)

	for i := range ids {
		ids[i] = i
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result = RenderNew(ids)
	}
}
