package problem454

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	a, b, c, d []int
	ans        int
}{
	{
		[]int{0, 1, -1},
		[]int{-1, 1, 0},
		[]int{0, 0, 1},
		[]int{-1, 1, 1},
		17,
	},

	{
		[]int{-1, -1},
		[]int{-1, 1},
		[]int{-1, 1},
		[]int{1, -1},
		6,
	},

	{
		[]int{1, 2},
		[]int{-2, -1},
		[]int{-1, 2},
		[]int{0, 2},
		2,
	},

	{
		[]int{0, 0},
		[]int{0, 0},
		[]int{0, 0},
		[]int{0, 0},
		16,
	},

	{
		[]int{-100, 0},
		[]int{1, 2},
		[]int{2, 3},
		[]int{4, 5},
		0,
	},

	{
		[]int{1, 2},
		[]int{2, 3},
		[]int{4, 5},
		[]int{6, 7},
		0,
	},

	{
		[]int{},
		[]int{},
		[]int{},
		[]int{},
		0,
	},
}

func Test_fourSumCount(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, fourSumCount(q.a, q.b, q.c, q.d), "输入 %v", q.a)
	}
}

func Benchmark_fourSumCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			fourSumCount(q.a, q.b, q.c, q.d)
		}
	}
}
