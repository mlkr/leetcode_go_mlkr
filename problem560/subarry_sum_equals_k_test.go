package problem560

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums   []int
	k, ans int
}{
	{
		[]int{-1, -1, 1},
		0,
		1,
	},

	{
		[]int{1, 1, 1},
		2,
		2,
	},

	{
		[]int{1, 2, 3},
		3,
		2,
	},
}

func Test_subarraySum(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, subarraySum(q.nums, q.k), "输入 %v", q.nums)
	}
}

func Benchmark_subarraySum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			subarraySum(q.nums, q.k)
		}
	}
}
