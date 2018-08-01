package problem69

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

func TestMySqrt(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		question{
			para{4},
			ans{2},
		},

		question{
			para{8},
			ans{2},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(mySqrt(para.one), ans.one, "输入 %v", para.one)
	}
}
