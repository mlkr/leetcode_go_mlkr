package problem674

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  int
}{
	{
		[]int{1, 3, 5, 4, 7},
		3,
	},

	{
		[]int{2, 2, 2, 2, 2},
		1,
	},
}

func Test_findLengthOfLCIS(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, findLengthOfLCIS(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_findLengthOfLCIS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findLengthOfLCIS(q.nums)
		}
	}
}
