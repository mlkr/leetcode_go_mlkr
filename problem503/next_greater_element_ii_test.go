package problem503

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums, ans []int
}{
	{
		[]int{1, 2, 3, 4, 5},
		[]int{2, 3, 4, 5, -1},
	},

	{
		[]int{1, 2, 1},
		[]int{2, -1, 2},
	},

	{
		[]int{1, 2, 3, 4, 3},
		[]int{2, 3, 4, -1, 4},
	},
}

func Test_nextGreaterElements(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, nextGreaterElements(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_nextGreaterElements(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			nextGreaterElements(q.nums)
		}
	}
}
