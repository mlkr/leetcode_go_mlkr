package problem209

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s    int
	nums []int
	ans  int
}{
	{
		4,
		[]int{1, 4, 4},
		1,
	},

	{
		7,
		[]int{2, 3, 1, 2, 4, 3},
		2,
	},
}

func Test_minSubArrayLen(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, minSubArrayLen(q.s, q.nums))
	}
}

func Benchmark_minSubArrayLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			minSubArrayLen(q.s, q.nums)
		}
	}
}
