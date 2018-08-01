package problem61

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

func TestRotateRight(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		question{
			para{
				[]int{1, 2, 3, 4, 5},
				2,
			},
			ans{[]int{4, 5, 1, 2, 3}},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(l2a(rotateRight(a2l(para.one), para.two)), ans.one, "输入 %v", para)
	}
}

func a2l(arr []int) *ListNode {
	len := len(arr)
	if len == 0 {
		return &ListNode{}
	}

	head := &ListNode{
		Val: arr[0],
	}

	temp := head
	for i := 1; i < len; i++ {
		temp.Next = &ListNode{
			Val: arr[i],
		}

		temp = temp.Next
	}

	return head
}

func l2a(l *ListNode) []int {
	res := []int{}

	for l != nil {
		res = append(res, l.Val)
		l = l.Next
	}

	return res
}
