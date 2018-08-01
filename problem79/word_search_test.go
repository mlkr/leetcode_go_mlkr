package problem79

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one [][]byte
	two string
}

type ans struct {
	one bool
}

type question struct {
	para
	ans
}

func TestExist(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{
				[][]byte{
					[]byte{'A', 'B', 'C', 'E'},
					[]byte{'S', 'F', 'C', 'S'},
					[]byte{'A', 'D', 'E', 'E'},
				},
				"ABCCED",
			},
			ans{true},
		},

		question{
			para{
				[][]byte{
					[]byte{'A', 'B', 'C', 'E'},
					[]byte{'S', 'F', 'C', 'S'},
					[]byte{'A', 'D', 'E', 'E'},
				},
				"SEE",
			},
			ans{true},
		},

		question{
			para{
				[][]byte{
					[]byte{'A', 'B', 'C', 'E'},
					[]byte{'S', 'F', 'C', 'S'},
					[]byte{'A', 'D', 'E', 'E'},
				},
				"ABCB",
			},
			ans{false},
		},

		question{
			para{
				[][]byte{
					[]byte{'a', 'b'},
					[]byte{'c', 'd'},
				},
				"abcd",
			},
			ans{false},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(exist(para.one, para.two), ans.one, "输入 %v", para)
	}
}
