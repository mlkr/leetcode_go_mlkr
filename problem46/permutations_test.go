package problem46

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

func Test_permute(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		question{
			para{[]int{1, 2, 3}},
			ans{[][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			}},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(ans.one, permute(para.one), "输入 %v", para.one)
	}
}
