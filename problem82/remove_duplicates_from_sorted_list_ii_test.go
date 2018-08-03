package problem82

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
}

type ans struct {
	one []int
}

type question struct {
	para
	ans
}

func TestDeleteDuplicates(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{[]int{1, 1}},
			ans{[]int{}},
		},

		question{
			para{[]int{1, 2, 3, 3, 4, 4, 5}},
			ans{[]int{1, 2, 5}},
		},

		question{
			para{[]int{1, 1, 1, 2, 3}},
			ans{[]int{2, 3}},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(l2s(deleteDuplicates(s2l(para.one))), ans.one, "输入 %v", para)
	}
}

// 链表转切片
func l2s(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// 切片转链表
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
