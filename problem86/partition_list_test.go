package problem86

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
	para
	ans
}

func TestPartition(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{[]int{},
				0,
			},
			ans{[]int{}},
		},

		question{
			para{[]int{
				1, 4, 3, 2, 5, 2,
			},
				3,
			},
			ans{[]int{
				1, 2, 2, 4, 3, 5,
			}},
		},

		question{
			para{[]int{
				1, 4, 3, 2, 5, 2,
			},
				5,
			},
			ans{[]int{
				1, 4, 3, 2, 2, 5,
			}},
		},

		question{
			para{[]int{
				1, 4, 3, 2, 5, 2,
			},
				6,
			},
			ans{[]int{
				1, 4, 3, 2, 5, 2,
			}},
		},

		question{
			para{[]int{
				1, 4, 3, 2, 5, 2,
			},
				1,
			},
			ans{[]int{
				1, 4, 3, 2, 5, 2,
			}},
		},

		question{
			para{[]int{
				1, 4, 3, 2, 5, 2,
			},
				0,
			},
			ans{[]int{
				1, 4, 3, 2, 5, 2,
			}},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(l2s(partition(s2l(para.one), para.two)), ans.one, "输入 %v", para)
	}
}

func l2s(head *ListNode) []int {
	res := []int{}
	if head == nil {
		return res
	}

	for {
		res = append(res, head.Val)

		if head.Next == nil {
			break
		}
		head = head.Next
	}

	return res
}

func s2l(nums []int) *ListNode {
	head := &ListNode{}
	temp := head
	for _, v := range nums {
		head.Next = &ListNode{
			Val: v,
		}
		head = head.Next

	}

	return temp.Next
}
