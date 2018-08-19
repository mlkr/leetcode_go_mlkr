package problem91

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
}

type ans struct {
	one int
}

type question struct {
	para
	ans
}

func TestNumdecodeings(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{"12"},
			ans{2},
		},

		question{
			para{"226"},
			ans{3},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(numDecodings(para.one), ans.one, "输入 %v", para)
	}
}
