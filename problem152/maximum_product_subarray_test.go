package problem152

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	input  []int
	output int
}{
	{
		[]int{2, 3, -2, 4},
		6,
	},

	{
		[]int{-2, 0, -1},
		0,
	},

	{
		[]int{-2, 0, -1, -3},
		3,
	},

	{
		[]int{-2, 0, 1, 3},
		3,
	},
}

func Test_maxProduct(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(maxProduct(q.input), q.output)
	}
}

func Benchmark_maxProduct(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		for _, q := range questions {
			maxProduct(q.input)
		}
	}
}
