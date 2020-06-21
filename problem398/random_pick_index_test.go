package problem398

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	pick int
	ans  []int
}{
	{
		[]int{1, 2, 3, 3, 3},
		3,
		[]int{2, 3, 4},
	},

	{
		[]int{1, 2, 3, 3, 3},
		1,
		[]int{0},
	},
}

func Test_Solution(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		s := Constructor(q.nums)
		ast.Contains(q.ans, s.Pick(q.pick), "输入 %v %d", q.nums, q.pick)
	}
}

func Benchmark_Solution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			s := Constructor(q.nums)
			s.Pick(q.pick)
		}
	}
}
