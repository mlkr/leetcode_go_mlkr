package problem483

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  bool
}{
	{
		[]int{0},
		true,
	},

	{
		[]int{1, 5, 2},
		false,
	},

	{
		[]int{1, 5, 233, 7},
		true,
	},
}

func Test_PredictTheWinner(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, PredictTheWinner(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_PredictTheWinner(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			PredictTheWinner(q.nums)
		}
	}
}
