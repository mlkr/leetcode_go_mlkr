package problem473

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  bool
}{
	{
		[]int{1, 1, 2, 2, 2},
		true,
	},

	{
		[]int{3, 3, 3, 3, 4},
		false,
	},

	{
		[]int{3, 3, 3, 3, 1},
		false,
	},

	{
		[]int{3, 3, 3},
		false,
	},

	{
		[]int{7, 3, 3, 3},
		false,
	},

	// seen的用处，看这个例子
	{
		[]int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3},
		true,
	},
}

func Test_makesquare(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, makesquare(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_makesquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			makesquare(q.nums)
		}
	}
}
