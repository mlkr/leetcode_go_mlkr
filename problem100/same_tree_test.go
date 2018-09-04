package problem100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tree struct {
	para, in []int
}

type para struct {
	one, two tree
}

type ans struct {
	one bool
}

type question struct {
	para
	ans
}

func TestIsSameTree(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{
				tree{
					[]int{10, 5, 15},
					[]int{5, 10, 15},
				},
				tree{
					[]int{10, 5, 15},
					[]int{5, 15, 10},
				},
			},
			ans{false},
		},

		question{
			para{
				tree{
					[]int{},
					[]int{},
				},
				tree{
					[]int{0},
					[]int{0},
				},
			},
			ans{false},
		},

		question{
			para{
				tree{
					[]int{1, 2, 3},
					[]int{2, 1, 3},
				},
				tree{
					[]int{1, 2, 3},
					[]int{2, 1, 3},
				},
			},
			ans{true},
		},

		question{
			para{
				tree{
					[]int{1, 2},
					[]int{2, 1},
				},
				tree{
					[]int{1, 2},
					[]int{1, 2},
				},
			},
			ans{false},
		},

		question{
			para{
				tree{
					[]int{1, 2, 1},
					[]int{2, 1, 1},
				},
				tree{
					[]int{1, 1, 2},
					[]int{1, 1, 2},
				},
			},
			ans{false},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		root1 := prein2Tree(para.one.para, para.one.in)
		root2 := prein2Tree(para.two.para, para.two.in)
		ast.Equal(isSameTree(root1, root2), ans.one, "输入 %v ", para)
	}
}

// 先序中序遍历转二叉树
func prein2Tree(pre, in []int) *TreeNode {
	preLen, inLen := len(pre), len(in)
	if preLen != inLen {
		panic("pre, in 数组长度不一样")
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
		if n == v {
			return k
		}
	}

	panic("数组中没有找到值")
}
