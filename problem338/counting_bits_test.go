package problem338

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	input int
	ans   []int
}{
	{
		2,
		[]int{0, 1, 1},
	},

	{
		5,
		[]int{0, 1, 1, 2, 1, 2},
	},
}

func Test_countBits(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, countBits(q.input))
	}
}

func Benchmark_countBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			countBits(q.input)
		}
	}
}
