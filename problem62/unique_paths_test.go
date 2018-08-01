package problem62

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
	two int
}

type ans struct {
	one int
}

type question struct {
	para
	ans
}

func TestUniquePaths(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		question{
			para{3, 2},
			ans{3},
		},

		question{
			para{7, 3},
			ans{28},
		},

		question{
			para{
				23,
				12,
			},
			ans{
				193536720,
			},
		},

		question{
			para{
				100,
				100,
			},
			ans{
				4631081169483718960,
			},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(uniquePaths(para.one, para.two), ans.one, "输入 %v", para)
	}
}
