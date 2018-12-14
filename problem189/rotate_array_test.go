package problem189

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	input []int
	k     int
	ans   []int
}{
	{
		[]int{1, 2, 3, 4, 5, 6, 7},
		3,
		[]int{5, 6, 7, 1, 2, 3, 4},
	},

	{
		[]int{-1, -100, 3, 99},
		2,
		[]int{3, 99, -1, -100},
	},
}

func Test_rotate(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		rotate(q.input, q.k)
		ast.Equal(q.input, q.ans)
	}
}

func Test_rotate2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		rotate2(q.input, q.k)
		ast.Equal(q.input, q.ans)
	}
}

func Test_rotate3(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		rotate3(q.input, q.k)
		ast.Equal(q.input, q.ans)
	}
}

func Benchmark_rotate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			rotate(q.input, q.k)
		}
	}
}

func Benchmark_rotate2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			rotate2(q.input, q.k)
		}
	}
}

func Benchmark_rotate3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			rotate3(q.input, q.k)
		}
	}
}
