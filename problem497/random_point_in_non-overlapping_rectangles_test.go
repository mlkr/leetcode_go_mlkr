package problem497

import (
	"testing"
)

var questions = []struct {
	rects [][]int
}{
	{
		[][]int{
			[]int{-2, -2, 1, 1},
			[]int{2, 2, 4, 6},
			[]int{},
			[]int{},
			[]int{},
			[]int{},
			[]int{},
		},
	},

	{
		[][]int{},
	},
}

func Test_Solution(t *testing.T) {
	for _, q := range questions {
		s := Constructor(q.rects)
		s.Pick()
		s.Pick()
		s.Pick()
		s.Pick()
	}
}

func Benchmark_Solution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			s := Constructor(q.rects)
			s.Pick()
			s.Pick()
			s.Pick()
			s.Pick()
		}
	}
}
