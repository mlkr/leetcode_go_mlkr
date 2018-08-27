package problem95

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
}

type ans struct {
	// 中序序列
	in []int
}

type question struct {
	para
	ans
}

func TestGenerateTrees(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{3},
			ans{
				[]int{1, 2, 3},
			},
		},

		question{
			para{4},
			ans{
				[]int{1, 2, 3, 4},
			},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans

		nodes := generateTrees(para.one)
		for _, node := range nodes {
			ast.Equal(inorderTraversal(node), ans.in, "输入 %v")
		}
	}

	ast.Nil(generateTrees(0), "generateTrees(0) == nil")

}

// 得二叉树的中序序列
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	roots := []*TreeNode{}
	res := []int{}
	for len(roots) != 0 || root != nil {

		if root != nil {
			roots = append(roots, root)

			for root.Left != nil {
				root = root.Left
				roots = append(roots, root)
			}
		}

		node := roots[len(roots)-1:]
		roots = roots[:len(roots)-1]
		res = append(res, node[0].Val)

		root = node[0].Right
	}

	return res
}
