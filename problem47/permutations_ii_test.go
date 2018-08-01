package problem47

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

func Test_permuteUnique(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		question{
			para{[]int{1, 1, 2}},
			ans{[][]int{
				[]int{1, 1, 2},
				[]int{1, 2, 1},
				[]int{2, 1, 1},
			}},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(permuteUnique(para.one), ans.one, "输入 %v", para)
	}
}
