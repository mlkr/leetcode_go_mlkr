package problem80

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
}

type ans struct {
	one int
}

type question struct {
	para
	ans
}

func TestRemoveDuplicates(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{[]int{1, 1, 1, 2, 2, 3}},
			ans{5},
		},

		question{
			para{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}},
			ans{7},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(removeDuplicates(para.one), ans.one, "输入 %v", para)
	}
}
