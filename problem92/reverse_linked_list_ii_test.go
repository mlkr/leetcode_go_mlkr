package problem92

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
	m   int
	n   int
}

type ans struct {
	one []int
}

type question struct {
	para
	ans
}

func TestReverseBetween(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{
				[]int{3, 5},
				1,
				2,
			},
			ans{
				[]int{5, 3},
			},
		},

		question{
			para{
				[]int{1, 2, 3, 4, 5},
				2,
				4,
			},
			ans{
				[]int{1, 4, 3, 2, 5},
			},
		},

		question{
			para{
				[]int{1, 2, 3, 4, 5},
				2,
				2,
			},
			ans{
				[]int{1, 2, 3, 4, 5},
			},
		},

		question{
			para{
				[]int{1, 2, 3, 4, 5, 6, 7, 8},
				2,
				4,
			},
			ans{
				[]int{1, 4, 3, 2, 5, 6, 7, 8},
			},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(l2s(reverseBetween(s2l(para.one), para.m, para.n)), ans.one, "输入 %v", para)
	}

}

func s2l(num []int) *ListNode {
	numLen := len(num)
	if numLen == 0 {
		return nil
	}

	head := &ListNode{}
	temp := head
	for i := 0; i < numLen; i++ {
		temp.Val = num[i]

		if i != numLen-1 {
			temp.Next = &ListNode{}
			temp = temp.Next

		}
	}

	return head
}

func l2s(head *ListNode) []int {
	res := []int{}
	for {
		if head == nil {
			break
		}

		res = append(res, head.Val)
		head = head.Next
	}

	return res
}
