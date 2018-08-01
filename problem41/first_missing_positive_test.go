package problem41

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para []int
	ans  int
}

func Test_firstMissingPositive(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			[]int{1, 2, 0},
			3,
		},
		question{
			[]int{3, 4, -1, 1},
			2,
		},
		question{
			[]int{7, 8, 9, 11, 12},
			1,
		},
		question{
			[]int{},
			1,
		},
	}

	for _, q := range qs {
		ans, para := q.ans, q.para
		ast.Equal(ans, firstMissingPositive(para), "输入 %v", para)
	}
}
