package problem89

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
}

type ans struct {
	one []int
}

type question struct {
	para
	ans
}

func TestGrayCode(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{2},
			ans{[]int{0, 1, 3, 2}},
		},

		question{
			para{0},
			ans{[]int{0}},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(grayCode(para.one), ans.one, "输入 %v", para)
	}
}
