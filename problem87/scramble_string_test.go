package problem87

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

func TestIsScramble(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{
				"great",
				"rgeat",
			},
			ans{
				true,
			},
		},

		question{
			para{
				"great",
				"rgaet",
			},
			ans{
				true,
			},
		},
		question{
			para{
				"great",
				"aterg",
			},
			ans{
				true,
			},
		},

		question{
			para{
				"abcde",
				"caebd",
			},
			ans{
				false,
			},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(isScramble(para.one, para.two), ans.one, "输入 %v", para)
	}
}
