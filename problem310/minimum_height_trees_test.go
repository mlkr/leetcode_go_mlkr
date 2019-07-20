package problem310

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	n    int
	edgs [][]int
	ans  []int
}{
	{
		4,
		[][]int{
			[]int{1, 0},
			[]int{1, 2},
			[]int{1, 3},
		},
		[]int{1},
	},

	{
		6,
		[][]int{
			[]int{0, 3},
			[]int{1, 3},
			[]int{2, 3},
			[]int{4, 3},
			[]int{5, 4},
		},
		[]int{3, 4},
	},

	{
		2,
		[][]int{
			[]int{0, 1},
		},
		[]int{0, 1},
	},

	{
		1,
		[][]int{
			[]int{1},
		},
		[]int{0},
	},
}

func Test_findMinHeightTrees(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, findMinHeightTrees(q.n, q.edgs))
	}
}

func Benchmark_findMinHeightTrees(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findMinHeightTrees(q.n, q.edgs)
		}
	}
}
