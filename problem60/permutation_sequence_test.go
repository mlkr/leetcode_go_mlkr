package problem60

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
	two int
}

type ans struct {
	one string
}

type question struct {
	para
	ans
}

func TestGetPermutation(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{3, 3},
			ans{"213"},
		},
	}

	for _, q := range questions {
		ans, para := q.ans, q.para
		ast.Equal(getPermutation(para.one, para.two), ans.one, "输入 %v", para)
	}
}
