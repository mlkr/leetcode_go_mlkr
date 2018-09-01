package problem98

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	pre []int
	in  []int
}

type ans struct {
	one bool
}

type question struct {
	para
	ans
}

func TestIsValidBST(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		question{
			para{
				[]int{1, 1},
				[]int{1, 1},
			},
			ans{false},
		},

		question{
			para{
				[]int{},
				[]int{},
			},
			ans{true},
		},

		question{
			para{
				[]int{2, 1, 3},
				[]int{1, 2, 3},
			},
			ans{true},
		},

		question{
			para{
				[]int{5, 1, 4, 3, 6},
				[]int{1, 5, 3, 4, 6},
			},
			ans{false},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(isValidBST(prein2Tree(para.pre, para.in)), ans.one, "输入 %v", para)
		ast.Equal(isValidBST2(prein2Tree(para.pre, para.in)), ans.one, "输入 %v", para)
	}
}

// 先序遍历,中序遍历转二叉树
func prein2Tree(pre, in []int) *TreeNode {
	preLen, inLen := len(pre), len(in)
	if preLen != inLen {
		panic("先序遍历,中序遍历 数组长度不一样")
	}

	if preLen <= 0 {
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

func indexOf(val int, nums []int) int {
	for idx, v := range nums {
		if val == v {
			return idx
		}
	}

	msg := fmt.Sprintf("数组中没有 %d", val)
	panic(msg)
}
