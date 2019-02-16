package problem216

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	k, n int
	ans  [][]int
}{
	{
		5,
		3,
		[][]int{},
	},
	{
		3,
		7,
		[][]int{
			[]int{1, 2, 4},
		},
	},

	{
		3,
		9,
		[][]int{
			[]int{1, 2, 6},
			[]int{1, 3, 5},
			[]int{2, 3, 4},
		},
	},
}

func Test_combinationSum3(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(combinationSum3(q.k, q.n), q.ans)
	}
}

func Benchmark_combinationSum3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			combinationSum3(q.k, q.n)
		}
	}
}
