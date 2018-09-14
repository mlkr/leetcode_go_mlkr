package problem110

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

func TestIsBalanced(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{
				[]int{3, 9, 20, 15, 7},
				[]int{9, 3, 15, 20, 7},
			},
			ans{true},
		},

		question{
			para{
				[]int{1, 2, 3, 4, 4, 3, 2},
				[]int{4, 3, 4, 2, 3, 1, 2},
			},
			ans{false},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		// ast.Equal(isBalanced(prein2Tree(para.pre, para.in)), ans.one, "输入 %v", para)
		ast.Equal(isBalanced2(prein2Tree(para.pre, para.in)), ans.one, "输入 %v", para)
	}
}

func prein2Tree(pre, in []int) *TreeNode {
	preLen, inLen := len(pre), len(in)
	if preLen != inLen {
		panic("pre in 数组长度不一致!")
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

	panic("要找的值,数组中没有!")
}
