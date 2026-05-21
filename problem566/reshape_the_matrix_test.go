package problem566

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	mat  [][]int
	r, c int
	ans  [][]int
}{
	{
		[][]int{
			[]int{1, 2},
			[]int{3, 4},
		},
		1,
		4,
		[][]int{
			[]int{1, 2, 3, 4},
		},
	},

	{
		[][]int{
			[]int{1, 2},
			[]int{3, 4},
		},
		2,
		4,
		[][]int{
			[]int{1, 2},
			[]int{3, 4},
		},
	},
}

func Test_matrixReshape(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, matrixReshape(q.mat, q.r, q.c))
	}
}

func Test_matrixReshape2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, matrixReshape2(q.mat, q.r, q.c))
	}
}

func Benchmark_matrixReshape(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			matrixReshape(q.mat, q.r, q.c)
		}
	}
}

func Benchmark_matrixReshape2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			matrixReshape2(q.mat, q.r, q.c)
		}
	}
}
