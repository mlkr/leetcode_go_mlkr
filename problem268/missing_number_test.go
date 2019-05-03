package problem268

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  int
}{
	{
		[]int{3, 0, 1},
		2,
	},

	{
		[]int{9, 6, 4, 2, 3, 5, 7, 0, 1},
		8,
	},
}

func Test_missingNumber(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, missingNumber(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_missingNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			missingNumber(q.nums)
		}
	}
}
