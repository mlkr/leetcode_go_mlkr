package problem76

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
	two string
}

type ans struct {
	one string
}

type question struct {
	para
	ans
}

func TestMinWindow(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		question{
			para{"ADOBECODEBANC", "ABC"},
			ans{"BANC"},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(minWindow(para.one, para.two), ans.one, "输入 %v", para)
	}
}
