package problem35

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
	two int
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

func Test_searchInsert(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				[]int{1, 3, 5, 6},
				5,
			},
			a: ans{
				2,
			},
		},
		question{
			p: para{
				[]int{1, 3, 5, 6},
				2,
			},
			a: ans{
				1,
			},
		},
		question{
			p: para{
				[]int{1, 3, 5, 6},
				7,
			},
			a: ans{
				4,
			},
		},
		question{
			p: para{
				[]int{1, 3, 5, 6},
				0,
			},
			a: ans{
				0,
			},
		},
		question{
			p: para{
				[]int{1, 3},
				2,
			},
			a: ans{
				1,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, searchInsert(p.one, p.two), "输入: %v", p)
	}
}
