package problem384

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
}{
	{
		[]int{1, 2, 3},
	},

	{
		[]int{8, 9, 3, 4, 8, 3, 1, 6, 7},
	},
}

func Test_Solution(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		s := Constructor(q.nums)
		ast.ElementsMatch(q.nums, s.Shuffle(), "输入 %v", q.nums)
		ast.Equal(q.nums, s.Reset(), "输入 %v", q.nums)
		ast.ElementsMatch(q.nums, s.Shuffle(), "输入 %v", q.nums)
	}
}

func Benchmark_Solution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			s := Constructor(q.nums)
			s.Shuffle()
			s.Reset()
			s.Shuffle()
		}
	}
}
