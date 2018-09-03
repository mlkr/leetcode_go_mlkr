package problem99

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

type para struct {
	pre, in []int
}

type ans struct {
	one bool
}

func TestRecoverTree(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{
				[]int{146, 71, 55, 321, -33, -13, 231, 399},
				[]int{-33, 321, 55, 71, 146, 231, -13, 399},
			},
			ans{true},
		},

		question{
			para{
				[]int{1, 3, 2},
				[]int{3, 2, 1},
			},
			ans{true},
		},

		question{
			para{
				[]int{3, 1, 4, 2},
				[]int{1, 3, 2, 4},
			},
			ans{true},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		root := prein2Tree(para.pre, para.in)
		// recoverTree(root)
		// recoverTree2(root)
		recoverTree3(root)
		ast.Equal(isValidBST(root), ans.one)
	}
}

// 先序序列, 中序序列转二叉树
func prein2Tree(pre, in []int) *TreeNode {
	preLen, inLen := len(pre), len(in)
	if preLen != inLen {
		panic("先序序列, 中序序列 数组长度不一致")
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

	idx := indexOf(pre[0], in)
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

	msg := fmt.Sprintf("%v 不在数组中", n)
	panic(msg)
}

// 验证是否为搜索二叉树
func isValidBST(root *TreeNode) bool {
	// int 最小值, 最大值
	min, max := -1<<63, 1<<63-1

	return recur(min, max, root)
}

func recur(min, max int, root *TreeNode) bool {
	if root == nil {
		return true
	}

	return min < root.Val && root.Val < max &&
		recur(min, root.Val, root.Left) &&
		recur(root.Val, max, root.Right)
}
