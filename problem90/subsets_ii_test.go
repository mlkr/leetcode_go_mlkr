package problem90

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
}

type ans struct {
	one [][]int
}

type question struct {
	para
	ans
}

func TestSubsetsWithDup(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{[]int{1, 2, 2}},
			ans{[][]int{
				[]int{},
				[]int{1},
				[]int{1, 2},
				[]int{1, 2, 2},
				[]int{2},
				[]int{2, 2},
			}},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		// ast.Equal(subsetsWithDup(para.one), ans.one, "输入 %v", para)
		ast.Equal(subsetsWithDup2(para.one), ans.one, "输入 %v", para)
	}
}
