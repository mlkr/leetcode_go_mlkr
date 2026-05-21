package problem553

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  string
}{
	{
		[]int{1000, 100, 10, 2},
		"1000/(100/10/2)",
	},

	{
		[]int{2, 3, 4},
		"2/(3/4)",
	},

	{
		[]int{2, 3},
		"2/3",
	},

	{
		[]int{2},
		"2",
	},
}

func Test_optimalDivision(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, optimalDivision(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_optimalDivision(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			optimalDivision(q.nums)
		}
	}
}
