package problem452

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	point [][]int
	ans   int
}{

	{
		[][]int{
			{10, 16},
			{2, 8},
			{1, 6},
			{7, 12},
		},
		2,
	},

	// 这个例子说明，为什么不能用start 从小到大排序，要用end
	{
		[][]int{
			{0, 9},
			{1, 6},
			{7, 8},
		},
		2,
	},

	{
		[][]int{{3, 9}, {7, 12}, {3, 8}, {6, 8}, {9, 10}, {2, 9}, {0, 9}, {3, 9}, {0, 6}, {2, 8}},
		2,
	},

	{
		[][]int{
			{1, 2},
			{3, 4},
			{5, 6},
			{7, 8},
		},
		4,
	},

	{
		[][]int{
			{1, 2},
		},
		1,
	},

	{
		[][]int{
			{2, 3},
			{2, 3},
		},
		1,
	},

	{
		[][]int{},
		0,
	},
}

func Test_findMinArrowShots(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, findMinArrowShots(q.point), "输入 %v", q.point)
	}
}

func Benchmark_findMinArrowShots(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findMinArrowShots(q.point)
		}
	}
}
