package problem493

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  int
}{
	{
		[]int{1, 3, 2, 3, 1},
		2,
	},

	{
		[]int{2, 4, 3, 5, 1},
		3,
	},
}

func Test_reversePairs(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, reversePairs(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_reversePairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			reversePairs(q.nums)
		}
	}
}
