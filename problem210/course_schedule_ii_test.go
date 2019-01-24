package problem210

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	numCourses    int
	prerequisites [][]int
	ans           []int
}{
	{
		3,
		[][]int{
			[]int{1, 0},
			[]int{1, 2},
			[]int{0, 1},
		},
		[]int{},
	},

	{
		2,
		[][]int{
			[]int{0, 1},
		},
		[]int{1, 0},
	},

	{
		2,
		[][]int{
			[]int{1, 0},
		},
		[]int{0, 1},
	},

	{
		4,
		[][]int{
			[]int{1, 0},
			[]int{2, 0},
			[]int{3, 1},
			[]int{3, 2},
		},
		[]int{0, 1, 2, 3},
	},
}

func Test_findOrder(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, findOrder(q.numCourses, q.prerequisites))
	}
}

func Benchmark_findOrder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findOrder(q.numCourses, q.prerequisites)
		}
	}
}
