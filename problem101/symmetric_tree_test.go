package problem101

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	pre, in []int
}

type ans struct {
	one bool
}

type question struct {
	para
	ans
}

func TestIsSymmetric(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{
				[]int{2, 3, 4, 5, 3, 4},
				[]int{4, 3, 5, 2, 3, 4},
			},
			ans{false},
		},

		question{
			para{
				[]int{1, 2, 3, 3, 2},
				[]int{3, 2, 1, 2, 3},
			},
			ans{false},
		},

		question{
			para{
				[]int{1, 2, 3, 4, 2, 4, 3},
				[]int{3, 2, 4, 1, 4, 2, 3},
			},
			ans{true},
		},

		question{
			para{
				[]int{1, 2, 3, 2, 3},
				[]int{2, 3, 1, 2, 3},
			},
			ans{false},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(isSymmetric(prein2Tree(para.pre, para.in)), ans.one, "输入 %v", para)
	}
}

// 先,中序遍历转二叉树
func prein2Tree(pre, in []int) *TreeNode {
	preLen, inLen := len(pre), len(in)
	if preLen != inLen {
		panic("先,中序遍历数组长度不一致")
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

	panic("数组中没有找到要查找的值")
}
