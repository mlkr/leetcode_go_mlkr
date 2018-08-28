package problem96

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
}

type ans struct {
	one int
}

type question struct {
	para
	ans
}

func TestNumTrees(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{7},
			ans{429},
		},

		question{
			para{5},
			ans{42},
		},

		question{
			para{19},
			ans{1767263190},
		},

		question{
			para{3},
			ans{5},
		},

		question{
			para{4},
			ans{14},
		},

		question{
			para{1},
			ans{1},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(numTrees(para.one), ans.one, "输入 %v", para)
	}
}
