package problem59

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
}

type ans struct {
	one [][]int
}

type question struct {
	para
	ans
}

func TestGenerateMatrix(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		// question{
		// 	para{0},
		// 	ans{[][]int{}},
		// },

		question{
			para{3},
			ans{[][]int{
				[]int{1, 2, 3},
				[]int{8, 9, 4},
				[]int{7, 6, 5},
			}},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(generateMatrix(para.one), ans.one, "输入 %v", para.one)
	}
}
