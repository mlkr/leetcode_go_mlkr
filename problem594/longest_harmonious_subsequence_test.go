package problem594

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  int
}{
	{
		[]int{3, 2, 2, 3, 2, 1, 3, 3, 3, -2, 0, 3, 2, 1, 0, 3, 1, 0, 1, 3, 0, 3, 3},
		14,
	},

	{
		[]int{1, 4, 1, 3, 1, -14, 1, -13},
		2,
	},

	{
		[]int{1, 2},
		2,
	},

	{
		[]int{1, 2, 2, 3, 4, 5, 1, 1, 1, 1},
		7,
	},

	{
		[]int{1, 3, 5, 7, 9, 11, 13, 15, 17},
		0,
	},

	{
		[]int{1, 2, 2, 1},
		4,
	},

	{
		[]int{1, 3, 2, 2, 5, 2, 3, 7},
		5,
	},

	{
		[]int{1, 2, 3, 4},
		2,
	},

	{
		[]int{1, 1, 1, 1},
		0,
	},
}

func Test_findLHS(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, findLHS(q.nums), "输入 %v", q.nums)
	}
}

func Test_findLHS2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, findLHS2(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_findLHS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findLHS(q.nums)
		}

	}
}

func Benchmark_findLHS2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findLHS2(q.nums)
		}

	}
}
