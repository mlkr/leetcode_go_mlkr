package problem65

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
}

type ans struct {
	one bool
}

type question struct {
	para
	ans
}

func TestIsNumber(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{"0"},
			ans{true},
		},

		question{
			para{" 0.1 "},
			ans{true},
		},

		question{
			para{"abc"},
			ans{false},
		},

		question{
			para{"1 a"},
			ans{false},
		},

		question{
			para{"2e10"},
			ans{true},
		},

		question{
			para{"e"},
			ans{false},
		},

		question{
			para{"."},
			ans{false},
		},

		question{
			para{".1"},
			ans{true},
		},

		question{
			para{" "},
			ans{false},
		},

		question{
			para{"1  "},
			ans{true},
		},

		question{
			para{"  1"},
			ans{true},
		},

		question{
			para{"-1."},
			ans{true},
		},

		question{
			para{"3-2"},
			ans{false},
		},

		question{
			para{"+.8"},
			ans{true},
		},

		question{
			para{" -."},
			ans{false},
		},

		question{
			para{"6e6.5"},
			ans{false},
		},

		question{
			para{" 005047e+6"},
			ans{true},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(isNumber(para.one), ans.one, "输入 %v", para.one)
	}
}
