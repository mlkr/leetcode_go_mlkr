package problem525

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  int
}{
	{
		[]int{0, 1, 1, 0, 1, 1, 1, 0},
		4,
	},

	{
		[]int{0, 0, 1, 0, 0, 0, 1, 1, 0},
		6,
	},

	{
		[]int{0, 1},
		2,
	},

	{
		[]int{0, 1, 0},
		2,
	},
}

func Test_findMaxLength(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, findMaxLength(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_findMaxLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findMaxLength(q.nums)
		}
	}
}
