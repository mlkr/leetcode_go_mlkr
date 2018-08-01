package problem67

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
	two string
}

type ans struct {
	one string
}

type question struct {
	para
	ans
}

func TestAddBinary(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{"11", "1"},
			ans{"100"},
		},

		question{
			para{"1010", "1011"},
			ans{"10101"},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(addBinary(para.one, para.two), ans.one, "输入 %v", para)
	}
}
