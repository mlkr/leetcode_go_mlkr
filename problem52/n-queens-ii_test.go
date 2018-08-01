package problem52

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

func TestTotalNQueens(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		question{
			para{2},
			ans{0},
		},

		question{
			para{4},
			ans{2},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans

		ast.Equal(totalNQueens(para.one), ans.one, "输入 %v", para.one)
	}
}
