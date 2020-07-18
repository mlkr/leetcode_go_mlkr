package problem417

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	matrix, ans [][]int
}{
	{
		[][]int{},
		[][]int{},
	},

	{
		[][]int{
			{},
		},

		[][]int{},
	},

	{
		[][]int{
			{1, 2, 2, 3, 5},
			{3, 2, 3, 4, 4},
			{2, 4, 5, 3, 1},
			{6, 7, 1, 4, 5},
			{5, 1, 1, 2, 4},
		},

		[][]int{
			{0, 4},
			{1, 3},
			{1, 4},
			{2, 2},
			{3, 0},
			{3, 1},
			{4, 0},
		},
	},
}

func Test_pacificAtlantic(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, pacificAtlantic(q.matrix), "输入 %v", q.matrix)
	}
}

func Benchmark_pacificAtlantic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			pacificAtlantic(q.matrix)
		}
	}
}
