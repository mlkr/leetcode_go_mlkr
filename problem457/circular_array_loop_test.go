package problem457

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  bool
}{
	{
		[]int{1, -1, 2, 4, 4},
		true,
	},

	{
		[]int{2, -1, 1, 2, 2},
		true,
	},

	{
		[]int{-1, 2},
		false,
	},

	{
		[]int{-2, 1, -1, -2, -2},
		false,
	},
}

func Test_circularArrayLoop(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, circularArrayLoop(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_circularArrayLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			circularArrayLoop(q.nums)
		}
	}
}
