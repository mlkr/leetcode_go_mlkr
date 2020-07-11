package problem410

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	m    int
	ans  int
}{
	{
		[]int{2, 3, 1, 2, 4, 3},
		5,
		4,
	},

	{
		[]int{7, 2, 5, 10, 8},
		2,
		18,
	},

	{
		[]int{7, 2, 5, 10, 8},
		3,
		14,
	},
}

func Test_splitArray(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, splitArray(q.nums, q.m), "输入 %v", q.nums)
	}
}

func Benchmark_splitArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			splitArray(q.nums, q.m)
		}
	}
}
