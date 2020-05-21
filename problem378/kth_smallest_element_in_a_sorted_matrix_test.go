package problem378

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	matrix [][]int
	k      int
	ans    int
}{
	{
		[][]int{
			[]int{1, 4, 4},
			[]int{4, 4, 4},
			[]int{4, 4, 15},
		},
		5,
		4,
	},

	{
		[][]int{
			[]int{1, 3, 3},
			[]int{3, 3, 3},
			[]int{3, 3, 15},
		},
		5,
		3,
	},

	{
		[][]int{
			[]int{1, 5, 9},
			[]int{10, 11, 13},
			[]int{12, 13, 15},
		},
		8,
		13,
	},
}

func Test_kthSmallest(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, kthSmallest(q.matrix, q.k), "输入 %v", q.matrix)
	}
}

func Benchmark_kthSmallest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			kthSmallest(q.matrix, q.k)
		}
	}
}
