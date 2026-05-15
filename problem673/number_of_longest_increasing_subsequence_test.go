package problem673

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  int
}{
	{
		[]int{1, 3, 5, 4, 7},
		2,
	},

	{
		[]int{2, 2, 2, 2, 2},
		5,
	},
}

func Test_findNumberOfLIS(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, findNumberOfLIS(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_findNumberOfLIS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findNumberOfLIS(q.nums)
		}
	}
}
