package problem73

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one [][]int
}

type ans struct {
	one [][]int
}

type question struct {
	para
	ans
}

func TestSetZeroes(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{
				[][]int{
					[]int{0, 1, 2, 0},
					[]int{3, 4, 5, 2},
					[]int{1, 3, 1, 5},
				},
			},
			ans{
				[][]int{
					[]int{0, 0, 0, 0},
					[]int{0, 4, 5, 0},
					[]int{0, 3, 1, 0},
				},
			},
		},

		question{
			para{
				[][]int{
					[]int{1, 1, 1},
					[]int{1, 0, 1},
					[]int{1, 1, 1},
				},
			},
			ans{
				[][]int{
					[]int{1, 0, 1},
					[]int{0, 0, 0},
					[]int{1, 0, 1},
				},
			},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		setZeroes(para.one)
		ast.Equal(para.one, ans.one, "输入 %v", para.one)
	}
}
