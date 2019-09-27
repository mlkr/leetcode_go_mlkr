package problem329

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	matrix [][]int
	ans    int
}{
	{
		[][]int{
			[]int{9, 9, 4},
			[]int{6, 6, 8},
			[]int{2, 1, 1},
		},
		4,
	},

	{
		[][]int{
			[]int{3, 4, 5},
			[]int{3, 2, 6},
			[]int{2, 2, 1},
		},
		4,
	},

	{
		[][]int{},
		0,
	},
}

func Test_longestIncreasingPath(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, longestIncreasingPath(q.matrix), "输入 %v", q.matrix)
	}
}

func Benchmark_longestIncreasingPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			longestIncreasingPath(q.matrix)
		}
	}
}
