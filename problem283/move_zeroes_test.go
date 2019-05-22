package problem283

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	input, output []int
}{
	{
		[]int{1},
		[]int{1},
	},

	{
		[]int{1, 0},
		[]int{1, 0},
	},

	{
		[]int{1, 0, 1},
		[]int{1, 1, 0},
	},

	{
		[]int{0, 1, 0, 3, 12},
		[]int{1, 3, 12, 0, 0},
	},
}

func Test_moveZeroes(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		moveZeroes(q.input)
		ast.Equal(q.output, q.input)
	}
}

func Test_moveZeroes2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		moveZeroes2(q.input)
		ast.Equal(q.output, q.input)
	}
}

func Benchmark_moveZeroes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			moveZeroes(q.input)
		}
	}
}

func Benchmark_moveZeroes2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			moveZeroes2(q.input)
		}
	}
}
