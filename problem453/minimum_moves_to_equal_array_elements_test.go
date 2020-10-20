package problem453

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
		3,
	},

	{
		[]int{3, 2, 1},
		3,
	},
}

func Test_minMoves(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, minMoves(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_minMoves(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			minMoves(q.nums)
		}
	}
}
