package problem598

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	m, n int
	ops  [][]int
	ans  int
}{

	{
		3,
		3,
		[][]int{
			[]int{2, 2},
			[]int{3, 3},
		},
		4,
	},

	{
		3,
		3,
		[][]int{},
		9,
	},

	{
		3,
		3,
		[][]int{
			[]int{2, 2},
			[]int{3, 3},
			[]int{3, 3},
			[]int{3, 3},
			[]int{2, 2},
			[]int{3, 3},
			[]int{3, 3},
			[]int{3, 3},
			[]int{2, 2},
			[]int{3, 3},
			[]int{3, 3},
			[]int{3, 3},
		},
		4,
	},
}

func Test_maxCount(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, maxCount(q.m, q.n, q.ops))
	}
}

func Benchmark_maxCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			maxCount(q.m, q.n, q.ops)
		}
	}
}
