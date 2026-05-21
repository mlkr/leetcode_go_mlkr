package problem517

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	machines []int
	ans      int
}{
	{
		[]int{4, 0, 0, 4},
		2,
	},

	{
		[]int{1, 0, 5},
		3,
	},

	{
		[]int{9, 1, 8, 8, 9},
		4,
	},

	{
		[]int{5, 2, 5},
		1,
	},

	{
		[]int{0, 3, 0},
		2,
	},

	{
		[]int{0, 2, 0},
		-1,
	},
}

func Test_findMinMoves(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, findMinMoves(q.machines), "输入 %v", q.machines)
	}
}

func Benchmark_findMinMoves(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findMinMoves(q.machines)
		}
	}
}
