package problem33

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para2 struct {
	one []int
	two int
}

type question2 struct {
	p para2
	a ans
}

func Test_Search(t *testing.T) {
	ast := assert.New(t)

	qs := []question2{
		question2{
			p: para2{
				[]int{6, 1, 2, 3, 4, 5},
				-1,
			},
			a: ans{
				-1,
			},
		},
		question2{
			p: para2{
				[]int{6, 1, 2, 3, 4, 5},
				99,
			},
			a: ans{
				-1,
			},
		},
		question2{
			p: para2{
				[]int{6, 1, 2, 3, 4, 5},
				6,
			},
			a: ans{
				0,
			},
		},
		question2{
			p: para2{
				[]int{4, 5, 6, 1, 2, 3},
				5,
			},
			a: ans{
				1,
			},
		},
		question2{
			p: para2{
				[]int{4, 5, 6, 1, 2, 3},
				2,
			},
			a: ans{
				4,
			},
		},
		question2{
			p: para2{
				[]int{3, 5, 1},
				1,
			},
			a: ans{
				2,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, search(p.one, p.two), "输入: %v", p)
	}
}
