package problem665

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  bool
}{

	{
		[]int{2, 3, 3, 2, 2},
		false,
	},

	{
		[]int{3, 4, 2, 3},
		false,
	},

	{
		[]int{4, 2, 3},
		true,
	},

	{
		[]int{4, 2, 1},
		false,
	},
}

func Test_checkPossibility(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, checkPossibility(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_checkPossibility(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			checkPossibility(q.nums)
		}
	}
}
