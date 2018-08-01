package problem64

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one [][]int
}

type ans struct {
	one int
}

type question struct {
	para
	ans
}

func TestMinPathSum(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		question{
			para{
				[][]int{
					[]int{1, 3, 1},
					[]int{1, 5, 1},
					[]int{4, 2, 1},
				},
			},
			ans{7},
		},

		question{
			para{
				[][]int{
					[]int{1, 2},
					[]int{5, 6},
					[]int{1, 1},
				},
			},
			ans{8},
		},

		question{
			para{
				[][]int{
					[]int{0, 1},
					[]int{1, 0},
				},
			},
			ans{1},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(minPathSum(para.one), ans.one, "输入 %v", para.one)
	}
}
