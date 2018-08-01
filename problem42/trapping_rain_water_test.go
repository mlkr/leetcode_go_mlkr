package problem42

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para []int
	ans  int
}

func Test_trap(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		question{
			[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			6,
		},

		question{
			[]int{},
			0,
		},
	}

	for _, qs := range questions {
		para, ans := qs.para, qs.ans

		ast.Equal(ans, trap(para), "输入 %v", para)
	}
}
