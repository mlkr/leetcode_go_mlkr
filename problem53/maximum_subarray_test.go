package problem53

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
}

type ans struct {
	one int
}

type question struct {
	para
	ans
}

func TestMaxSubArray(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		question{
			para{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			ans{6},
		},

		question{
			para{[]int{-2, -1, -4, -3}},
			ans{-1},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(ans.one, maxSubArray(para.one), "输入 %v", para.one)
	}
}
