package problem207

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	numCourses    int
	prerequisites [][]int
	ans           bool
}{
	{
		2,
		[][]int{
			[]int{1, 0},
		},
		true,
	},

	{
		2,
		[][]int{
			[]int{1, 0},
			[]int{0, 1},
		},
		false,
	},
}

func Test_canFinish(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(canFinish(q.numCourses, q.prerequisites), q.ans)
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			canFinish(q.numCourses, q.prerequisites)
		}
	}
}
