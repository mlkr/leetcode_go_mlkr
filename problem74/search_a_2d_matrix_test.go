package problem74

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one [][]int
	two int
}

type ans struct {
	one bool
}

type question struct {
	para
	ans
}

func TestSearchMatrix(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{[][]int{
				[]int{1, 3},
			},
				3,
			},
			ans{true},
		},

		question{
			para{[][]int{
				[]int{1},
			},
				1,
			},
			ans{true},
		},

		question{
			para{[][]int{
				[]int{1},
			},
				0,
			},
			ans{false},
		},

		question{
			para{[][]int{
				[]int{},
			},
				1,
			},
			ans{false},
		},

		question{
			para{[][]int{
				[]int{1, 3, 5, 7},
				[]int{10, 11, 16, 20},
				[]int{23, 30, 34, 50},
			},
				3,
			},
			ans{true},
		},

		question{
			para{[][]int{
				[]int{1, 3, 5, 7},
				[]int{10, 11, 16, 20},
				[]int{23, 30, 34, 50},
			},
				13,
			},
			ans{false},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(searchMatrix(para.one, para.two), ans.one, "输入 %v", para)
	}
}
