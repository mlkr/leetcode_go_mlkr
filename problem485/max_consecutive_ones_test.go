package problem485

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  int
}{
	{
		[]int{1, 1, 0, 1, 1, 1},
		3,
	},

	{
		[]int{1, 0, 1, 1, 0, 1},
		2,
	},
}

func Test_findMaxConsecutiveOnes(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, findMaxConsecutiveOnes(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_findMaxConsecutiveOnes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findMaxConsecutiveOnes(q.nums)
		}
	}
}
