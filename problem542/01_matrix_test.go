package problem542

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	mat [][]int
	ans [][]int
}{
	{
		[][]int{
			[]int{0},
			[]int{0},
			[]int{0},
			[]int{1},
			[]int{1},
			[]int{1},
		},

		[][]int{
			[]int{0},
			[]int{0},
			[]int{0},
			[]int{1},
			[]int{2},
			[]int{3},
		},
	},

	{
		[][]int{
			[]int{0, 1, 0},
			[]int{0, 1, 0},
			[]int{0, 1, 0},
			[]int{0, 1, 0},
			[]int{0, 1, 0},
		},

		[][]int{
			[]int{0, 1, 0},
			[]int{0, 1, 0},
			[]int{0, 1, 0},
			[]int{0, 1, 0},
			[]int{0, 1, 0},
		},
	},

	{
		[][]int{
			[]int{0, 0, 0},
			[]int{0, 1, 0},
			[]int{0, 0, 0},
		},

		[][]int{
			[]int{0, 0, 0},
			[]int{0, 1, 0},
			[]int{0, 0, 0},
		},
	},

	{
		[][]int{
			[]int{0, 0, 0},
			[]int{0, 1, 0},
			[]int{1, 1, 1},
		},

		[][]int{
			[]int{0, 0, 0},
			[]int{0, 1, 0},
			[]int{1, 2, 1},
		},
	},
}

func Test_updateMatrix(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, updateMatrix(q.mat), "输入 %v", q.mat)
	}
}

func Test_updateMatrix2(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, updateMatrix2(q.mat), "输入 %v", q.mat)
	}
}

func Benchmark_updateMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			updateMatrix(q.mat)
		}
	}
}

func Benchmark_updateMatrix2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			updateMatrix2(q.mat)
		}
	}
}
