package problem327

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums         []int
	lower, upper int
	ans          int
}{
	{
		[]int{-2, 5, -1},
		-2,
		2,
		3,
	},

	{
		[]int{},
		-2,
		2,
		0,
	},

	{
		[]int{-2, 5, -1},
		-1,
		1,
		1,
	},
}

func Test_countRangeSum(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, countRangeSum(q.nums, q.lower, q.upper), "输入 %v", q.ans)
	}
}

func Benchmark_countRangeSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			countRangeSum(q.nums, q.lower, q.upper)
		}
	}
}
