package problem97

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one, two, three string
}

type ans struct {
	one bool
}

type question struct {
	para
	ans
}

func TestIsInterleave(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{
				"aabcc",
				"dbbca",
				"aadbbcbcac",
			},
			ans{true},
		},

		question{
			para{
				"aabcc",
				"dbbca",
				"aadbbbaccc",
			},
			ans{false},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(isInterleave(para.one, para.two, para.three), ans.one)
	}
}
