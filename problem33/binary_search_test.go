package problem33

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one   []int
	two   int
	three int
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

func Test_binary_search(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				[]int{1, 2, 3, 4, 5, 6},
				-1,
				0,
			},
			a: ans{
				-1,
			},
		},
		question{
			p: para{
				[]int{1, 2, 3, 4, 5, 6},
				1,
				0,
			},
			a: ans{
				0,
			},
		},
		question{
			p: para{
				[]int{1, 2, 3, 4, 5, 6},
				2,
				0,
			},
			a: ans{
				1,
			},
		},
		question{
			p: para{
				[]int{1, 2, 3, 4, 5, 6},
				3,
				0,
			},
			a: ans{
				2,
			},
		},
		question{
			p: para{
				[]int{1, 2, 3, 4, 5, 6},
				5,
				0,
			},
			a: ans{
				4,
			},
		},
		question{
			p: para{
				[]int{1, 2, 3, 4, 5, 6},
				6,
				0,
			},
			a: ans{
				5,
			},
		},
		question{
			p: para{
				[]int{1, 2, 3, 4, 5, 6},
				99,
				0,
			},
			a: ans{
				-1,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, binary_search(p.one, p.two, p.three), "输入: %v", p)
	}
}
