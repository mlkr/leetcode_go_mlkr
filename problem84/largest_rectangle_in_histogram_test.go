package problem84

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

func TestLargestRectangleArea(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{
				[]int{2, 1, 5, 6, 2, 3},
			},
			ans{10},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(largestRectangleArea(para.one), ans.one, "输入 %v", para)
	}
}
