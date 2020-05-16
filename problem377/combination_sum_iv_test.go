package problem377

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums   []int
	target int
	ans    int
}{
	{
		[]int{2, 1, 3},
		35,
		1132436852,
	},

	{
		[]int{3, 4, 5, 6, 7, 8, 9, 10},
		10,
		9,
	},

	{
		[]int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25},
		10,
		9,
	},

	{
		[]int{1, 2, 3},
		4,
		7,
	},

	{
		[]int{5, 6, 7},
		4,
		0,
	},

	{
		[]int{},
		4,
		0,
	},
}

func Test_combinationSum4(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, combinationSum4(q.nums, q.target), "输入 %v", q.nums)
	}
}

func Benchmark_combinationSum4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			combinationSum4(q.nums, q.target)
		}
	}
}
