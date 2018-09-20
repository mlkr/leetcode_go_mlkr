package problem118

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	ast := assert.New(t)

	questions := []struct {
		numRows int
		ans     [][]int
	}{
		{
			0,
			[][]int{},
		},

		{
			1,
			[][]int{
				[]int{1},
			},
		},

		{
			2,
			[][]int{
				[]int{1},
				[]int{1, 1},
			},
		},

		{
			3,
			[][]int{
				[]int{1},
				[]int{1, 1},
				[]int{1, 2, 1},
			},
		},

		{
			4,
			[][]int{
				[]int{1},
				[]int{1, 1},
				[]int{1, 2, 1},
				[]int{1, 3, 3, 1},
			},
		},

		{
			5,
			[][]int{
				[]int{1},
				[]int{1, 1},
				[]int{1, 2, 1},
				[]int{1, 3, 3, 1},
				[]int{1, 4, 6, 4, 1},
			},
		},
	}

	for _, q := range questions {
		ast.Equal(generate(q.numRows), q.ans, "输入 %v", q.numRows)
	}
}
