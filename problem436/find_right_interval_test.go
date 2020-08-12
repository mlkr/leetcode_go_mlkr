package problem436

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct{
	intervals [][]int
	ans []int
}{
	{
		[][]int{
			{1,12},
			{2,9},
			{3,10},
			{13,14},
			{15,16},
			{16,17},
		},

		[]int{3,3,3,4,5,-1},
	},

	{
		[][]int{
			{1,2},
		},

		[]int{
			-1,
		},
	},

	{
		[][]int{
			{3,4},
			{2,3},
			{1,2},
		},

		[]int{
			-1,
			0,
			1,
		},
	},
}

func Test_findRightInterval(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, findRightInterval(q.intervals), "输入 %v", q.intervals)
	}
}

func Benchmark_findRightInterval(b *testing.B) {
	for i:=0;i<b.N;i++{
		for _, q := range questions {
			findRightInterval(q.intervals)
		}
	}
}