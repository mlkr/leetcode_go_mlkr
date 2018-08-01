package problem44

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
	two string
}

type ans struct {
	one bool
}

type question struct {
	para
	ans
}

func Test_isMatch(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		question{
			para{"aa", "a"},
			ans{false},
		},

		question{
			para{"aa", "*"},
			ans{true},
		},

		question{
			para{"cb", "?a"},
			ans{false},
		},

		question{
			para{"adceb", "a*b"},
			ans{true},
		},

		question{
			para{"acdcb", "a*c?b"},
			ans{false},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(isMatch(para.one, para.two), ans.one, "输入 %v", para)

	}

}
