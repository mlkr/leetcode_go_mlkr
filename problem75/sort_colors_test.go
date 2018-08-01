package problem75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
}

type ans struct {
	one []int
}

type question struct {
	para
	ans
}

func TestSortColors(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{
				[]int{1, 2, 0},
			},
			ans{
				[]int{0, 1, 2},
			},
		},

		question{
			para{
				[]int{2, 0, 1},
			},
			ans{
				[]int{0, 1, 2},
			},
		},

		question{
			para{
				[]int{2, 0, 2, 1, 1, 0},
			},
			ans{
				[]int{0, 0, 1, 1, 2, 2},
			},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		var input []int
		copy(input, para.one)
		sortColors(para.one)
		ast.Equal(para.one, ans.one, "输入 %v", input)
	}
}
