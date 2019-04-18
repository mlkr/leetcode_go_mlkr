package problem241

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	input string
	ans   []int
}{
	{
		"2-1-1",
		[]int{
			0,
			2,
		},
	},

	{
		"2*3-4*5",
		[]int{
			-34,
			-14,
			-10,
			-10,
			10,
		},
	},
}

func Test_diffWaysToCompute(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.ElementsMatch(q.ans, diffWaysToCompute(q.input))
	}
}

func Benchmark_diffWaysToCompute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			diffWaysToCompute(q.input)
		}
	}
}
