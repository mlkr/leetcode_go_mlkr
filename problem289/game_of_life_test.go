package problem289

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	board [][]int
	ans   [][]int
}{
	{
		[][]int{
			[]int{0, 1, 0},
			[]int{0, 0, 1},
			[]int{1, 1, 1},
			[]int{0, 0, 0},
		},

		[][]int{
			[]int{0, 0, 0},
			[]int{1, 0, 1},
			[]int{0, 1, 1},
			[]int{0, 1, 0},
		},
	},

	{
		[][]int{},

		[][]int{},
	},

	{
		[][]int{
			[]int{},
		},

		[][]int{
			[]int{},
		},
	},
}

func Test_gameOfLife(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		gameOfLife(q.board)
		ast.Equal(q.ans, q.board)
	}
}

func Benchmark_gameOfLife(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			gameOfLife(q.board)
		}
	}
}
