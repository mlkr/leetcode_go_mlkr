package problem71

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
}

type ans struct {
	one string
}

type question struct {
	para
	ans
}

func TestSimplifyPath(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{"/."},
			ans{"/"},
		},

		question{
			para{"/home//foo/"},
			ans{"/home/foo"},
		},

		question{
			para{"/home/"},
			ans{"/home"},
		},

		question{
			para{"/a/./b/../../c/"},
			ans{"/c"},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans

		ast.Equal(simplifyPath(para.one), ans.one, "输入 %v", para.one)
	}
}
