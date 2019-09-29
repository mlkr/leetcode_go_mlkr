package problem330

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	n    int
	ans  int
}{
	{
		[]int{1, 3},
		6,
		1,
	},

	{
		[]int{1, 5, 10},
		20,
		2,
	},

	{
		[]int{1, 2, 2},
		5,
		0,
	},

	{
		[]int{},
		7,
		3,
	},
}

func Test_minPatches(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, minPatches(q.nums, q.n), "输入 %v", q.nums)
	}
}

func Benchmark_minPatches(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			minPatches(q.nums, q.n)
		}
	}
}
