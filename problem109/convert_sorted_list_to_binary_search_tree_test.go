package problem109

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
}

type ans struct {
	one bool
}

type question struct {
	para
	ans
}

func TestSortedListToBST(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{
				[]int{-10, -3, 0, 5, 9},
			},
			ans{true},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		head := s2l(para.one)
		root := sortedListToBST(head)
		ast.Equal(isBST(root), ans.one, "输入 %v", para)
	}
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

func isBST(root *TreeNode) bool {
	min, max := 0, 0

	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			if min == 0 {
				min = level
			}

			if max == 0 {
				max = level
			}

			if level < min {
				min = level
			}

			if level > max {
				max = level
			}

			return
		}

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	if max-min > 1 {
		return false
	}

	return true
}
