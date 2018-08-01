package problem40

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans [][]int
}

type para struct {
	candidates []int
	target     int
}

func Test_combinationSum2(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			para{[]int{2, 5, 2, 1, 2}, 5},
			[][]int{
				{1, 2, 2},
				{5},
			},
		},
	}

	for _, q := range qs {
		para, ans := q.para, q.ans
		ast.Equal(ans, combinationSum2(para.candidates, para.target), "输入: %v", ans)
	}
}
