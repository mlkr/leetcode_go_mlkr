package problem103

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	pre, in []int
}

type ans struct {
	one [][]int
}

type question struct {
	para
	ans
}

func TestZigzagLevelOrder(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{
				[]int{1, 2, 4, 3, 5},
				[]int{2, 4, 1, 3, 5},
			},
			ans{
				[][]int{
					[]int{1},
					[]int{3, 2},
					[]int{4, 5},
				},
			},
		},

		question{
			para{
				[]int{3, 9, 20, 15, 7},
				[]int{9, 3, 15, 20, 7},
			},
			ans{
				[][]int{
					[]int{3},
					[]int{20, 9},
					[]int{15, 7},
				},
			},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(zigzagLevelOrder(prein2Tree(para.pre, para.in)), ans.one, "输入 %v", para)
		ast.Equal(zigzagLevelOrder2(prein2Tree(para.pre, para.in)), ans.one, "输入 %v", para)
		ast.Equal(zigzagLevelOrder3(prein2Tree(para.pre, para.in)), ans.one, "输入 %v", para)
	}
}

func prein2Tree(pre, in []int) *TreeNode {
	preLen, inLen := len(pre), len(in)
	if preLen != inLen {
		panic("pre in 数组不相等!")
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

	panic("数组里没右要找的值!")
}
