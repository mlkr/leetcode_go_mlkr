package problem368

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  []int
}{
	{
		[]int{1},
		[]int{1},
	},

	{
		[]int{1, 2, 3},
		[]int{1, 2},
	},

	{
		[]int{1, 2, 4, 8},
		[]int{1, 2, 4, 8},
	},

	{
		[]int{1, 2, 4, 5, 8},
		[]int{1, 2, 4, 8},
	},

	{
		[]int{},
		[]int{},
	},
}

func Test_largestDivisibleSubset(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, largestDivisibleSubset(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_largestDivisibleSubset(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			largestDivisibleSubset(q.nums)
		}
	}
}
