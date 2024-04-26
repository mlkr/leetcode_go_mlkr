package problem561

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  int
}{
	{
		[]int{1, 4, 3, 2},
		4,
	},

	{
		[]int{6, 2, 6, 5, 1, 2},
		9,
	},
}

func Test_arrayPairSum(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, arrayPairSum(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_arrayPairSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			arrayPairSum(q.nums)
		}
	}
}
