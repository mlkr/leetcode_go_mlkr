package problem88

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	nums1 []int
	m     int
	nums2 []int
	n     int
}

func (p *para) copy() para {
	res := para{
		nums1: make([]int, len(p.nums1)),
		nums2: make([]int, len(p.nums2)),
	}
	res.m = p.m
	res.n = p.n
	copy(res.nums1, p.nums1)
	copy(res.nums2, p.nums2)
	return res
}

type ans struct {
	one []int
}

type question struct {
	para
	ans
}

func TestMerge(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{
				[]int{1, 2, 4, 5, 6, 0},
				5,
				[]int{3},
				1,
			},
			ans{
				[]int{1, 2, 3, 4, 5, 6},
			},
		},

		question{
			para{
				[]int{1, 2, 3, 0, 0, 0},
				3,
				[]int{2, 5, 6},
				3,
			},
			ans{
				[]int{1, 2, 2, 3, 5, 6},
			},
		},

		question{
			para{
				[]int{2, 5, 6, 0, 0, 0},
				3,
				[]int{1, 2, 3},
				3,
			},
			ans{
				[]int{1, 2, 2, 3, 5, 6},
			},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		c := para.copy()
		// merge(para.nums1, para.m, para.nums2, para.n)
		merge2(para.nums1, para.m, para.nums2, para.n)
		ast.Equal(para.nums1, ans.one, "输入 %v", c)
	}
}
