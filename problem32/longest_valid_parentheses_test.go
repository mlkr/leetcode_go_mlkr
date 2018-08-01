package problem32

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
	p para
	a ans
}

func Test_LongestValidParentheses(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				")(()()))((((())))(",
			},
			a: ans{
				8,
			},
		},
		question{
			p: para{
				"()(()",
			},
			a: ans{
				2,
			},
		},
		question{
			p: para{
				"()()",
			},
			a: ans{
				4,
			},
		},
		question{
			p: para{
				"())()()(()",
			},
			a: ans{
				4,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, LongestValidParentheses(p.one), "输入: %v", p)
	}
}
