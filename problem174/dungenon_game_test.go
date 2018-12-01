package problem174

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para [][]int
	ans  int
}{
	{
		[][]int{
			[]int{-1, 1},
		},
		2,
	},

	{
		[][]int{
			[]int{100},
		},
		1,
	},

	{
		[][]int{
			[]int{-2, -3, 3},
			[]int{-5, -10, 1},
			[]int{10, 30, -5},
		},
		7,
	},
}

func Test_calculateMinimumHP(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(calculateMinimumHP(q.para), q.ans)
	}
}

func Test_calculateMinimumHP2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(calculateMinimumHP2(q.para), q.ans)
	}
}

func Benchmark_calculateMinimumHP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			calculateMinimumHP(q.para)
		}
	}
}
