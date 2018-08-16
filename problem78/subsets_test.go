package problem78

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
}

type ans struct {
	one [][]int
}

type question struct {
	para
	ans
}

func TestSubsets(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{
				[]int{1, 2, 3},
			},
			ans{
				[][]int{
					[]int{},
					[]int{1},
					[]int{2},
					[]int{1, 2},
					[]int{3},
					[]int{1, 3},
					[]int{2, 3},
					[]int{1, 2, 3},
				},
			},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(subsets2(para.one), ans.one, "输入 %v", para.one)
	}
}
