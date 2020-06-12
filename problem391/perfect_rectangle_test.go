package problem391

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	rectangles [][]int
	ans        bool
}{
	{
		[][]int{
			[]int{1, 1, 3, 3},
			[]int{3, 1, 4, 2},
			[]int{3, 2, 4, 4},
			[]int{1, 3, 2, 4},
			[]int{2, 3, 3, 4},
		},
		true,
	},

	{
		[][]int{
			[]int{1, 1, 2, 3},
			[]int{1, 3, 2, 4},
			[]int{3, 1, 4, 2},
			[]int{3, 2, 4, 4},
		},
		false,
	},

	{
		[][]int{
			[]int{1, 1, 3, 3},
			[]int{3, 1, 4, 2},
			[]int{1, 3, 2, 4},
			[]int{3, 2, 4, 4},
		},
		false,
	},

	{
		[][]int{
			[]int{1, 1, 3, 3},
			[]int{3, 1, 4, 2},
			[]int{1, 3, 2, 4},
			[]int{2, 2, 4, 4},
		},
		false,
	},

	{
		[][]int{},
		false,
	},

	{
		[][]int{
			[]int{},
		},
		false,
	},

	{
		[][]int{
			[]int{0, 0, 4, 1},
			[]int{7, 0, 8, 2},
			[]int{6, 2, 8, 3},
			[]int{5, 1, 6, 3},
			[]int{4, 0, 5, 1},
			[]int{6, 0, 7, 2},
			[]int{4, 2, 5, 3},
			[]int{2, 1, 4, 3},
			[]int{0, 1, 2, 2},
			[]int{0, 2, 2, 3},
			[]int{4, 1, 5, 2},
			[]int{5, 0, 6, 1},
		},
		true,
	},

	{
		[][]int{
			[]int{0, 0, 1, 1},
			[]int{0, 0, 2, 1},
			[]int{1, 0, 2, 1},
			[]int{0, 2, 2, 3},
		},
		false,
	},
}

func Test_isRectangleCover(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, isRectangleCover(q.rectangles), "输入 %v", q.rectangles)
	}
}

func Benchmark_isRectangleCover(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			isRectangleCover(q.rectangles)
		}
	}
}
