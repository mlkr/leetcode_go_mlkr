package problem81

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	nums   []int
	target int
}

type ans struct {
	one bool
}

type question struct {
	para
	ans
}

func TestSearch(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{
				[]int{},
				5,
			},
			ans{false},
		},

		question{
			para{
				[]int{2, 5, 6, 0, 0, 1, 2},
				0,
			},
			ans{true},
		},

		question{
			para{
				[]int{2, 5, 6, 0, 0, 1, 2},
				3,
			},
			ans{false},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(search(para.nums, para.target), ans.one, "输入 %v", para)
	}
}
