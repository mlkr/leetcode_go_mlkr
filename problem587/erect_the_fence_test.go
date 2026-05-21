package problem587

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	trees, ans [][]int
}{
	{
		[][]int{
			[]int{0, 0},
			[]int{0, 1},
			[]int{0, 2},
			[]int{1, 2},
			[]int{2, 2},
			[]int{3, 2},
			[]int{3, 1},
			[]int{3, 0},
			[]int{2, 0},
		},

		[][]int{
			[]int{0, 0},
			[]int{0, 1},
			[]int{0, 2},
			[]int{1, 2},
			[]int{2, 2},
			[]int{3, 2},
			[]int{3, 1},
			[]int{3, 0},
			[]int{2, 0},
		},
	},

	{
		[][]int{
			[]int{1, 1},
			[]int{2, 2},
			[]int{2, 0},
			[]int{2, 4},
			[]int{3, 3},
			[]int{4, 2},
		},
		[][]int{
			[]int{1, 1},
			[]int{2, 0},
			[]int{4, 2},
			[]int{3, 3},
			[]int{2, 4},
		},
	},

	{
		[][]int{
			[]int{1, 2},
			[]int{2, 2},
			[]int{4, 2},
		},
		[][]int{
			[]int{4, 2},
			[]int{2, 2},
			[]int{1, 2},
		},
	},

	{
		[][]int{
			[]int{1, 2},
			[]int{2, 2},
			[]int{4, 2},
			[]int{5, 2},
			[]int{6, 2},
			[]int{7, 2},
		},

		[][]int{
			[]int{4, 2},
			[]int{6, 2},
			[]int{2, 2},
			[]int{5, 2},
			[]int{1, 2},
			[]int{7, 2},
		},
	},
}

func Test_outerTrees(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.ElementsMatch(q.ans, outerTrees(q.trees), "输入 %v", q.trees)
	}
}

func Benchmark_outerTrees(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			outerTrees(q.trees)
		}
	}
}
