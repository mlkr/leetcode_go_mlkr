package problem238

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	input, output []int
}{
	{
		[]int{1, 2, 3, 4},
		[]int{24, 12, 8, 6},
	},

	{
		[]int{0, 0},
		[]int{0, 0},
	},

	{
		[]int{1, 0},
		[]int{0, 1},
	},

	{
		[]int{0, 1},
		[]int{1, 0},
	},

	{
		[]int{0, 4, 0},
		[]int{0, 0, 0},
	},
}

func Test_productExceptSelf(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.output, productExceptSelf(q.input))
	}
}

func Test_productExceptSelf2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.output, productExceptSelf2(q.input))
	}
}

func Benchmark_productExceptSelf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			productExceptSelf(q.input)
		}
	}
}

func Benchmark_productExceptSelf2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			productExceptSelf2(q.input)
		}
	}
}
