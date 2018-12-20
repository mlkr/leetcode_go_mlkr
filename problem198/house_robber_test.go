package problem198

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	input []int
	ans   int
}{
	{
		[]int{1, 2, 3, 1},
		4,
	},

	{
		[]int{2, 7, 9, 3, 1},
		12,
	},

	{
		[]int{2, 1, 1, 2},
		4,
	},

	{
		[]int{},
		0,
	},

	{
		[]int{0},
		0,
	},

	{
		[]int{2, 4, 8, 9, 9, 3},
		19,
	},
}

func Test_rob(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(rob(q.input), q.ans)
	}
}

func Test_rob2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(rob2(q.input), q.ans)
	}
}

func Benchmark_rob(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			rob(q.input)
		}
	}
}

func Benchmark_rob2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			rob2(q.input)
		}
	}
}
