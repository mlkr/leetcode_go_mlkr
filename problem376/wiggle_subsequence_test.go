package problem376

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  int
}{
	{
		[]int{3, 3, 3, 2, 5},
		3,
	},

	{
		[]int{0},
		1,
	},

	{
		[]int{0, 0},
		1,
	},

	{
		[]int{1, 7, 4, 9, 2, 5},
		6,
	},

	{
		[]int{1, 7, 7, 4, 9, 2, 5},
		6,
	},

	{
		[]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8},
		7,
	},

	{
		[]int{1},
		1,
	},

	{
		[]int{1, 2},
		2,
	},
}

func Test_wiggleMaxLength(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, wiggleMaxLength(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_wiggleMaxLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			wiggleMaxLength(q.nums)
		}
	}
}
