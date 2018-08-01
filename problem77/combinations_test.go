package problem77

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	n int
	k int
}

type ans struct {
	one [][]int
}

type question struct {
	para
	ans
}

func TestCombine(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{4, 2},
			ans{
				[][]int{
					[]int{1, 2},
					[]int{1, 3},
					[]int{1, 4},
					[]int{2, 3},
					[]int{2, 4},
					[]int{3, 4},
				},
			},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(combine(para.n, para.k), ans.one, "输入 %v", para)
	}
}
