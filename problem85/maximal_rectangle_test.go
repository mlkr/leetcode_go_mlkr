package problem85

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one [][]byte
}

type ans struct {
	one int
}

type question struct {
	para
	ans
}

func TestMaximalRectangle(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{[][]byte{
				[]byte{'1', '0', '1', '0', '0'},
				[]byte{'1', '0', '1', '1', '1'},
				[]byte{'1', '1', '1', '1', '1'},
				[]byte{'1', '0', '0', '1', '0'},
			}},
			ans{6},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(maximalRectangle(para.one), ans.one, "输入 %v", para)
	}
}
