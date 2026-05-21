package problem498

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	mat [][]int
	ans []int
}{
	{
		[][]int{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8, 9},
		},

		[]int{1, 2, 4, 7, 5, 3, 6, 8, 9},
	},

	{
		[][]int{
			[]int{1, 2},
			[]int{3, 4},
		},

		[]int{1, 2, 3, 4},
	},

	{
		[][]int{},
		[]int{},
	},

	{
		[][]int{
			[]int{},
		},
		[]int{},
	},
}

func Test_findDiagonalOrder(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, findDiagonalOrder(q.mat), "输入 %v", q.mat)
	}
}

func Benchmark_findDiagonalOrder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findDiagonalOrder(q.mat)
		}
	}
}
