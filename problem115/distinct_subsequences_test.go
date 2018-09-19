package problem115

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	s, t string
}

type ans struct {
	one int
}

type question struct {
	para
	ans
}

func TestNumDistinct(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{
				"aabb",
				"abb",
			},
			ans{
				2,
			},
		},

		question{
			para{
				"rabbbit",
				"rabbit",
			},
			ans{
				3,
			},
		},

		question{
			para{
				"babgbag",
				"bag",
			},
			ans{
				5,
			},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(numDistinct(para.s, para.t), ans.one, "输入 %v para")
		ast.Equal(numDistinct2(para.s, para.t), ans.one, "输入 %v para")
	}
}
