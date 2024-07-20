package problem581

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  int
}{
	// {
	// 	[]int{1, 2, 3, 3, 3},
	// 	0,
	// },

	{
		[]int{2, 3, 3, 2, 4},
		3,
	},

	{
		[]int{1, 3, 2, 2, 2},
		4,
	},

	{
		[]int{2, 6, 4, 8, 10, 9, 15},
		5,
	},

	{
		[]int{1, 2, 3, 4},
		0,
	},

	{
		[]int{1},
		0,
	},
}

func Test_findUnsortedSubarray(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, findUnsortedSubarray(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_findUnsortedSubarray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findUnsortedSubarray(q.nums)
		}
	}
}
