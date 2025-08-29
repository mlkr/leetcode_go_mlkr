package problem645

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  []int
}{
	{
		[]int{1, 1, 3},
		[]int{1, 2},
	},

	{
		[]int{3, 2, 3, 4, 6, 5},
		[]int{3, 1},
	},

	{
		[]int{1, 2, 2, 4},
		[]int{2, 3},
	},

	{
		[]int{1, 1},
		[]int{1, 2},
	},

	{
		[]int{2, 2},
		[]int{2, 1},
	},

	{
		[]int{3, 2, 2},
		[]int{2, 1},
	},

	{
		[]int{8, 7, 3, 5, 3, 6, 1, 4},
		[]int{3, 2},
	},
}

func Test_findErrorNums(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, findErrorNums(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_findErrorNums(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findErrorNums(q.nums)
		}
	}
}
