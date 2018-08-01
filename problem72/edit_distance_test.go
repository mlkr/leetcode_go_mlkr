package problem72

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	word1 string
	word2 string
}

type ans struct {
	one int
}

type question struct {
	para
	ans
}

func TestMinDistance(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{"horse", "ros"},
			ans{3},
		},

		question{
			para{"intention", "execution"},
			ans{5},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(minDistance(para.word1, para.word2), ans.one, "输入 %v", para)
	}
}
