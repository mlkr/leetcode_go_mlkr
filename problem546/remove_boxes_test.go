package problem546

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	boxes []int
	ans   int
}{

	{
		[]int{1, 2, 2, 1, 2, 1, 1, 2, 2, 1, 2, 1, 2, 1, 2, 2, 1, 2, 2, 1, 2, 1, 2, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 1, 2, 2, 1, 2, 1, 2, 2, 1},
		554,
	},

	{
		[]int{1, 1, 1, 2, 2, 2, 2, 1, 2, 1, 1, 2, 2, 1, 2, 1, 2, 2, 2, 2, 2, 1, 2, 2, 1, 1, 1, 2, 2, 1, 2, 1, 2, 2, 1, 2, 1, 1, 1, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 1, 1, 1, 1, 2, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 2, 1},
		1818,
	},

	{
		[]int{1, 2, 2, 1, 1, 1, 2, 1, 1, 2, 1, 2, 1, 1, 2, 2, 1, 1, 2, 2, 1, 1, 1, 2, 2, 2, 2, 1, 2, 1, 1, 2, 2, 1, 2, 1, 2, 2, 2, 2, 2, 1, 2, 1, 2, 2, 1, 1, 1, 2, 2, 1, 2, 1, 2, 2, 1, 2, 1, 1, 1, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 1, 1, 1, 1, 2, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 2, 1},
		2758,
	},

	{
		[]int{1, 3, 2, 2, 2, 3, 4, 3, 1},
		23,
	},

	{
		[]int{1, 1, 1},
		9,
	},

	{
		[]int{1},
		1,
	},
}

func Test_removeBoxes(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, removeBoxes(q.boxes), "输入 %v", q.boxes)
	}
}

func Benchmark_removeBoxes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			removeBoxes(q.boxes)
		}
	}
}
