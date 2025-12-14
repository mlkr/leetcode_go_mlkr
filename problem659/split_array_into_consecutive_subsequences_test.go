package problem659

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  bool
}{
	{
		[]int{4, 5, 6, 6, 7, 8, 9, 10, 10, 11},
		false,
	},

	{
		[]int{1, 2, 3, 3, 4, 5},
		true,
	},

	{
		[]int{1, 2, 3, 3, 4, 4, 5, 5},
		true,
	},

	{
		[]int{1, 2, 3, 4, 4, 5},
		false,
	},
}

func Test_isPossible(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, isPossible(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_isPossible(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			isPossible(q.nums)
		}
	}
}
