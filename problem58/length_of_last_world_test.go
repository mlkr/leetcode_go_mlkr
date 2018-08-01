package problem58

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
	para
	ans
}

func TestLengthOfLastWord(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{"Hello World"},
			ans{5},
		},

		question{
			para{""},
			ans{0},
		},

		question{
			para{"   "},
			ans{0},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(lengthOfLastWord(para.one), ans.one, "输入 %v", para.one)
	}
}
