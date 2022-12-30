package problem494

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums   []int
	target int
	ans    int
}{
	// 例1：
	{
		[]int{1, 3, 4},
		0,
		2,
	},

	{
		[]int{1, 1, 1, 1, 1},
		3,
		5,
	},

	{
		[]int{1},
		1,
		1,
	},

	{
		[]int{1},
		2,
		0,
	},

	{
		[]int{1},
		0,
		0,
	},

	{
		[]int{100},
		-200,
		0,
	},
}

func Test_findTargetSumWays(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, findTargetSumWays(q.nums, q.target), "nums %v", q.nums)
	}
}

func Benchmark_findTargetSumWays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findTargetSumWays(q.nums, q.target)
		}
	}
}
