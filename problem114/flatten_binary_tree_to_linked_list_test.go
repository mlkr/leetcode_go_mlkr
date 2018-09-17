package problem114

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	pre, in []int
}

type ans struct {
	pre   []int
	level [][]int
}

type question struct {
	para
	ans
}

func TestFlatten(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{
				[]int{1, 2, 3, 5, 4},
				[]int{5, 3, 2, 4, 1},
			},
			ans{
				[]int{1, 2, 3, 5, 4},
				[][]int{
					[]int{1},
					[]int{2},
					[]int{3},
					[]int{5},
					[]int{4},
				},
			},
		},

		question{
			para{
				[]int{1, 2},
				[]int{2, 1},
			},
			ans{
				[]int{1, 2},
				[][]int{
					[]int{1},
					[]int{2},
				},
			},
		},

		question{
			para{
				[]int{},
				[]int{},
			},
			ans{
				[]int{},
				[][]int{},
			},
		},

		question{
			para{
				[]int{1, 2, 3, 4, 5, 6},
				[]int{3, 2, 4, 1, 5, 6},
			},
			ans{
				[]int{1, 2, 3, 4, 5, 6},
				[][]int{
					[]int{1},
					[]int{2},
					[]int{3},
					[]int{4},
					[]int{5},
					[]int{6},
				},
			},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans

		root := prein2Tree(para.pre, para.in)
		flatten(root)

		pre := preOrder(root)
		level := levelOrder(root)

		ast.Equal(pre, ans.pre, "输入 %v", pre)
		ast.Equal(level, ans.level, "输入 %v", level)
	}
}

func prein2Tree(pre, in []int) *TreeNode {
	preLen, inLen := len(pre), len(in)
	if preLen != inLen {
		panic("pre in 数组不一致!")
	}

	if preLen == 0 {
		return nil
	}

	root := &TreeNode{
		Val: pre[0],
	}

	if preLen == 1 {
		return root
	}

	idx := indexOf(root.Val, in)

	root.Left = prein2Tree(pre[1:idx+1], in[:idx])
	root.Right = prein2Tree(pre[idx+1:], in[idx+1:])

	return root
}

func indexOf(n int, nums []int) int {
	for k, v := range nums {
		if v == n {
			return k
		}
	}

	panic("没有找到值!")
}

func preOrder(root *TreeNode) []int {
	res := []int{}
	roots := []*TreeNode{}

	if root == nil {
		return res
	}

	for root != nil || len(roots) != 0 {
		if root != nil {
			roots = append(roots, root)
			res = append(res, root.Val)

			root = root.Left
		} else {
			root = roots[len(roots)-1]
			roots = roots[:len(roots)-1]

			root = root.Right
		}

	}

	return res
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}

	var getLevelOrder func(root *TreeNode, level int)
	getLevelOrder = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if len(res) == level {
			res = append(res, []int{root.Val})
		} else {
			res[level] = append(res[level], root.Val)
		}

		getLevelOrder(root.Left, level+1)
		getLevelOrder(root.Right, level+1)
	}

	getLevelOrder(root, 0)

	return res
}
