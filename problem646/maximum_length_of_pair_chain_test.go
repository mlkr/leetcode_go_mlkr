package problem646

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	pairs [][]int
	ans   int
}{
	{
		[][]int{
			[]int{1, 2},
			[]int{2, 3},
			[]int{3, 4},
		},
		2,
	},

	{
		[][]int{
			[]int{1, 2},
			[]int{7, 8},
			[]int{4, 5},
		},
		3,
	},
}

func Test_findLongestChain(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, findLongestChain(q.pairs), "输入 %v", q.pairs)
	}
}

func Benchmark_findLongestChain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findLongestChain(q.pairs)
		}
	}
}
