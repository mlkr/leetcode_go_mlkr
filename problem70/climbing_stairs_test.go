package problem70

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
}

type ans struct {
	one int
}

type question struct {
	para
	ans
}

func TestClimbStairs(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{1},
			ans{1},
		},

		question{
			para{2},
			ans{2},
		},

		question{
			para{3},
			ans{3},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(climbStairs(para.one), ans.one, "输入 %v", para.one)
	}
}
