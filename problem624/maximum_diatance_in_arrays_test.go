package problem624

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	arrays [][]int
	ans    int
}{
	{
		[][]int{
			[]int{-2},
			[]int{-3, -2, 1},
		},
		3,
	},

	{
		[][]int{
			[]int{3, 4},
			[]int{1, 5},
		},
		3,
	},

	{
		[][]int{
			[]int{1, 2, 3},
			[]int{4, 5},
			[]int{1, 2, 3},
		},
		4,
	},

	{
		[][]int{
			[]int{1},
			[]int{1},
		},
		0,
	},
}

func Test_maxDistance(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, maxDistance(q.arrays), "输入 %v", q.arrays)
	}
}

func Benchmark_maxDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			maxDistance(q.arrays)
		}
	}
}
