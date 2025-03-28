package problem628

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  int
}{
	{
		[]int{1, 2, 3},
		6,
	},

	{
		[]int{1, 2, 3, 4},
		24,
	},

	{
		[]int{-1, -2, -3},
		-6,
	},

	{
		[]int{-3, -2, -1, 1, 2},
		12,
	},

	{
		[]int{-3, -2, -1, 1, 2, 3, 4},
		24,
	},

	{
		[]int{-4, -3, -2, -1},
		-6,
	},
}

func Test_maximumProduct(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, maximumProduct(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_maximumProduct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			maximumProduct(q.nums)
		}
	}
}
