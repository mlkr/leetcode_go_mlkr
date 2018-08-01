package problem45

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ans struct {
	one int
}

type para struct {
	one []int
}

type question struct {
	ans
	para
}

func Test_jump(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		question{
			ans{1},
			para{[]int{2, 3, 1}},
		},

		question{
			ans{2},
			para{[]int{2, 3, 1, 1, 4}},
		},
	}

	for _, q := range questions {
		ans, para := q.ans, q.para

		ast.Equal(ans.one, jump(para.one), "输入 %v", para)
	}

}
