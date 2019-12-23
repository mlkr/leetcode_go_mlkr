package problem347

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	k    int
	ans  []int
}{
	{
		[]int{1, 1, 1, 2, 2, 2, 3, 3, 3},
		0,
		[]int{},
	},

	{
		[]int{1, 1, 1, 2, 2, 2, 3, 3, 3},
		3,
		[]int{1, 3, 2},
	},

	{
		[]int{3, 2, 1, 1, 1},
		2,
		[]int{1, 3},
	},

	{
		[]int{1, 2, 3, 1, 1},
		2,
		[]int{1, 2},
	},

	{
		[]int{1, 1, 1, 2, 3},
		2,
		[]int{1, 2},
	},

	{
		[]int{1},
		1,
		[]int{1},
	},
}

func Test_topKFrequent(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.ElementsMatch(q.ans, topKFrequent(q.nums, q.k), "输入 %v", q.nums)
	}
}

func Benchmark_topKFrequent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			topKFrequent(q.nums, q.k)
		}
	}
}
