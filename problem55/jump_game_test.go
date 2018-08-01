package problem55

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
}

type ans struct {
	one bool
}

type question struct {
	para
	ans
}

func TestCanJump(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{[]int{1, 0, 0, 1, 1, 2, 2, 0, 2, 2}},
			ans{false},
		},

		question{
			para{[]int{1, 1, 2, 2, 0, 1, 1}},
			ans{true},
		},

		question{
			para{[]int{1, 1, 0, 1}},
			ans{false},
		},

		question{
			para{[]int{2, 3, 1, 1, 4}},
			ans{true},
		},

		question{
			para{[]int{3, 2, 1, 0, 4}},
			ans{false},
		},

		question{
			para{[]int{1}},
			ans{true},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(ans.one, canJump(para.one), "输入 %v", para.one)
	}
}
