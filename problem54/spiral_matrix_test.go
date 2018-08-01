package problem54

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one [][]int
}

type ans struct {
	one []int
}

type question struct {
	para
	ans
}

func TestSpiralOrder(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		question{
			para{[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}},
			ans{
				[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
			},
		},

		question{
			para{[][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			}},
			ans{
				[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
			},
		},
	}

	for _, q := range questions {
		ans, para := q.ans, q.para
		ast.Equal(ans.one, spiralOrder(para.one), "输入 %v", para.one)
	}
}
