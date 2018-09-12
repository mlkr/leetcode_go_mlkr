package problem108

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ans struct {
	one bool
}

type para struct {
	one []int
}

type question struct {
	para
	ans
}

func TestSortedArrayToBST(t *testing.T) {
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
		root := sortedArrayToBST(para.one)
		ast.Equal(isBST(root), ans.one, "输入 %v", para)
	}
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
