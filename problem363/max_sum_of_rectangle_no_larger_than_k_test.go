package problem363

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	matrix [][]int
	k      int
	ans    int
}{
	{
		[][]int{
			[]int{1, 0, 1},
			[]int{0, -2, 3},
		},
		2,
		2,
	},

	{
		[][]int{
			[]int{1, 0, 1},
			[]int{0, -2, 3},
		},
		-2,
		-2,
	},

	{
		[][]int{
			[]int{1, 0, 1},
			[]int{0, -2, 3},
		},
		5,
		4,
	},

	{
		[][]int{
			[]int{5, -4, -3, 4},
			[]int{-3, -4, 4, 5},
			[]int{5, 1, 5, -4},
		},
		8,
		8,
	},
}

func Test_maxSumSubmatrix(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, maxSumSubmatrix(q.matrix, q.k), "输入 %v", q.matrix)
	}
}

func Benchmark_maxSumSubmatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			maxSumSubmatrix(q.matrix, q.k)
		}
	}
}
