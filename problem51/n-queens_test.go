package problem51

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
}

type ans struct {
	one [][]string
}

type question struct {
	para
	ans
}

func TestSolveNQueens(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		question{
			para{4},
			ans{[][]string{
				[]string{".Q..", "...Q", "Q...", "..Q."},
				[]string{"..Q.", "Q...", "...Q", ".Q.."},
			}},
		},
	}

	for _, q := range questions {
		ans, para := q.ans, q.para
		ast.Equal(ans.one, solveNQueens(para.one), "输入 %v", para)
	}
}
