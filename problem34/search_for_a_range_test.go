package problem34

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
	two int
}

type ans struct {
	one []int
}

type question struct {
	p para
	a ans
}

func Test_searchRange(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				[]int{5, 7, 7, 8, 8, 10},
				8,
			},
			a: ans{
				[]int{3, 4},
			},
		},
		question{
			p: para{
				[]int{5, 7, 7, 8, 8, 10},
				6,
			},
			a: ans{
				[]int{-1, -1},
			},
		},
		question{
			p: para{
				[]int{1},
				1,
			},
			a: ans{
				[]int{0, 0},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, searchRange(p.one, p.two), "输入: %v", p)
	}
}
