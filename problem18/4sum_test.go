package problem18

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums   []int
	target int
	ans    [][]int
}{
	{
		[]int{-3, -2, -1, 0, 0, 1, 2, 3},
		0,
		[][]int{
			{-3, -2, 2, 3},
			{-3, -1, 1, 3},
			{-3, 0, 0, 3},
			{-3, 0, 1, 2},
			{-2, -1, 0, 3},
			{-2, -1, 1, 2},
			{-2, 0, 0, 2},
			{-1, 0, 0, 1},
		},
	},

	{
		[]int{1, 0, -1, 0, -2, 2},
		0,
		[][]int{
			{-2, -1, 1, 2},
			{-2, 0, 0, 2},
			{-1, 0, 0, 1},
		},
	},

	{
		[]int{},
		0,
		[][]int{},
	},

	{
		[]int{0, 0, 0, 0, 0, 0},
		0,
		[][]int{
			{0, 0, 0, 0},
		},
	},

	{
		[]int{0, 0, 0, 1, 1, 1},
		1,
		[][]int{
			{0, 0, 0, 1},
		},
	},
}

func Test_fourSum(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, fourSum(q.nums, q.target), "输入 %v", q.nums)
	}
}

func Benchmark_fourSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			fourSum(q.nums, q.target)
		}
	}
}
