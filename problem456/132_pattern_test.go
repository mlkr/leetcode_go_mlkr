package problem456

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  bool
}{
	{
		[]int{1, 2, 3, 4},
		false,
	},

	{
		[]int{3, 1, 4, 2},
		true,
	},

	{
		[]int{-1, 3, 2, 0},
		true,
	},
}

func Test_find132pattern(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, find132pattern(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_find132pattern(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			find132pattern(q.nums)
		}
	}
}
