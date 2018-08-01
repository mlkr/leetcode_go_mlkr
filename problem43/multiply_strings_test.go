package problem43

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

func Test_multiply(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		question{
			para{"2", "3"},
			ans{"6"},
		},

		question{
			para{"123", "456"},
			ans{"56088"},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(multiply(para.one, para.two), ans.one, "输入: %v", para)
	}
}
