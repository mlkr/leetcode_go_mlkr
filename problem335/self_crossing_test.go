package problem335

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	x   []int
	ans bool
}{
	{
		[]int{2, 2, 4, 4, 5, 2},
		false,
	},

	{
		[]int{2, 2, 4, 4, 4, 2},
		true,
	},

	{
		[]int{2, 2, 4, 4, 2, 1},
		false,
	},

	{
		[]int{2, 2, 4, 4, 2, 2},
		true,
	},

	{
		[]int{1, 1, 2, 1, 1},
		true,
	},

	{
		[]int{2, 1, 1},
		false,
	},

	{
		[]int{2, 1, 1, 2},
		true,
	},

	{
		[]int{1, 2, 3, 4},
		false,
	},

	{
		[]int{1, 1, 1, 1},
		true,
	},
}

func Test_isSelfCrossing(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, isSelfCrossing(q.x))
	}
}

func Benchmark_isSelfCrossing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			isSelfCrossing(q.x)
		}
	}
}
