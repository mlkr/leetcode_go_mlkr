package problem279

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	input, output int
}{
	{
		25,
		1,
	},

	{
		12,
		3,
	},

	{
		13,
		2,
	},

	{
		15,
		2,
	},
}

func Test_numSquares(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.output, numSquares(q.input))
	}
}

func Test_numSquares2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.output, numSquares2(q.input))
	}
}

func Benchmark_numSquares(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			numSquares(q.input)
		}
	}
}

func Benchmark_numSquares2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			numSquares2(q.input)
		}
	}
}
