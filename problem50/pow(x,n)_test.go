package problem50

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one float64
	two int
}

type ans struct {
	one float64
}

type question struct {
	para
	ans
}

func Test_myPow(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		question{
			para{2, 1},
			ans{2},
		},

		question{
			para{2, 3},
			ans{8},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(ans.one, myPow(para.one, para.two), "输入 %v", para)
	}
}
