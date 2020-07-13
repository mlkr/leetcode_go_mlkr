package problem414

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  int
}{
	{
		[]int{1, -2147483648, 2},
		-2147483648,
	},

	{
		[]int{1, 2, 2, 5, 3, 5},
		2,
	},

	{
		[]int{1, 2, -2147483648},
		-2147483648,
	},

	{
		[]int{5, 2, 2},
		5,
	},

	{
		[]int{3, 2, 1},
		1,
	},

	{
		[]int{1, 2},
		2,
	},

	{
		[]int{2, 2, 3, 1},
		1,
	},
}

func Test_thirdMax(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, thirdMax(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_thirdMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			thirdMax(q.nums)
		}
	}
}
