package problem120

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumTotal(t *testing.T) {
	ast := assert.New(t)

	questions := []struct {
		triangle [][]int
		path     int
	}{
		{
			[][]int{
				[]int{-1},
				[]int{2, 3},
				[]int{1, -1, -3},
			},
			-1,
		},

		{
			[][]int{
				[]int{2},
				[]int{3, 4},
				[]int{6, 5, 7},
				[]int{4, 1, 8, 3},
			},
			11,
		},
	}

	for _, q := range questions {
		// ast.Equal(minimumTotal(q.triangle), q.path)
		ast.Equal(minimumTotal2(q.triangle), q.path)
	}
}
