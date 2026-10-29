package task3

import "testing"

// BenchmarkJoinIDsSlow-12          1678242               724.8 ns/op           240 B/op         19 allocs/op
func BenchmarkJoinIDsSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JoinIDsSlow([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}

// BenchmarkJoinIDsFast-12         12628866                96.47 ns/op           80 B/op          1 allocs/op
func BenchmarkJoinIDsFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JoinIDsFast([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}

func TestSlowAllocations(t *testing.T) {
	allocs := testing.AllocsPerRun(100, func() {
		_ = JoinIDsSlow([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	})
	t.Logf("allocs: %v", allocs)
}

func TestFastAllocations(t *testing.T) {
	allocs := testing.AllocsPerRun(100, func() {
		_ = JoinIDsFast([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	})
	t.Logf("allocs: %v", allocs)
	if allocs > 1 {
		t.Errorf("allocs: %v", allocs)
	}
}
