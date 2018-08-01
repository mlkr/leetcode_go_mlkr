package problem66

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
}

type ans struct {
	one []int
}

type question struct {
	para
	ans
}

func TestPlusOne(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{[]int{1, 2, 3}},
			ans{[]int{1, 2, 4}},
		},

		question{
			para{[]int{4, 3, 2, 1}},
			ans{[]int{4, 3, 2, 2}},
		},

		question{
			para{[]int{8, 9, 9, 9}},
			ans{[]int{9, 0, 0, 0}},
		},

		question{
			para{[]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}},
			ans{[]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 7}},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(plusOne(para.one), ans.one, "输入 %v", para.one)
	}
}
