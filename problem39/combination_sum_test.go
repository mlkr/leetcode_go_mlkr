package problem39

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

type para struct {
	candidates []int
	target     int
}

type ans struct {
	one [][]int
}

func Test_combinationSum(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			para{[]int{2, 3, 6, 7}, 7},
			ans{[][]int{
				[]int{2, 2, 3},
				[]int{7},
			}},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		ast.Equal(a.one, combinationSum(p.candidates, p.target), "输入:%v", p)
	}
}
