package problem119

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRow(t *testing.T) {
	ast := assert.New(t)

	questions := []struct {
		rowIndex int
		ans      []int
	}{
		{
			3,
			[]int{
				1, 3, 3, 1,
			},
		},
	}

	for _, q := range questions {
		// ast.Equal(getRow(q.rowIndex), q.ans)
		ast.Equal(getRow2(q.rowIndex), q.ans)
	}
}
