package problem462

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  int
}{
	{
		[]int{1, 2, 3},
		2,
	},

	{
		[]int{1, 1, 2, 3},
		3,
	},

	{
		[]int{8},
		0,
	},

	{
		[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		25,
	},
}

func Test_minMoves22(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, minMoves22(q.nums), "输入 %v", q.nums)
	}
}

func Test_minMoves2(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, minMoves2(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_minMoves2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			minMoves2(q.nums)
		}
	}
}

func Benchmark_minMoves22(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			minMoves22(q.nums)
		}
	}
}
