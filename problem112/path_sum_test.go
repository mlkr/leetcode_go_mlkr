package problem112

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	pre, in []int
	sum     int
}

type ans struct {
	one bool
}

type question struct {
	para
	ans
}

func TestHasPathSun(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{
				[]int{5, 4, 11, 7, 2, 8, 13, 4, 1},
				[]int{7, 11, 2, 4, 5, 13, 8, 4, 1},
				22,
			},
			ans{
				true,
			},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(hasPathSum(prein2Tree(para.pre, para.in), para.sum), ans.one, "输入 %v", para)
		ast.Equal(hasPathSum2(prein2Tree(para.pre, para.in), para.sum), ans.one, "输入 %v", para)
	}
}

func prein2Tree(pre, in []int) *TreeNode {
	preLen, inLen := len(pre), len(in)
	if preLen != inLen {
		panic("pre, in 数组长度不一致!")
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

	panic("数组中没有要找的值!")
}
