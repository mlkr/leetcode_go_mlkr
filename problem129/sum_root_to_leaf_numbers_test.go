package problem129

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumNumbers(t *testing.T) {
	ast := assert.New(t)

	questions := []struct {
		pre []int
		in  []int
		sum int
	}{
		{
			[]int{0, 1},
			[]int{1, 0},
			1,
		},

		{
			[]int{1, 2, 3},
			[]int{2, 1, 3},
			25,
		},

		{
			[]int{4, 9, 5, 1, 0},
			[]int{5, 9, 1, 4, 0},
			1026,
		},
	}

	for _, q := range questions {
		root := prein2Tree(q.pre, q.in)
		ast.Equal(sumNumbers(root), q.sum)
		ast.Equal(sumNumbers2(root), q.sum)
	}
}

func prein2Tree(pre, in []int) *TreeNode {
	preLen, inLen := len(pre), len(in)
	if preLen != inLen {
		panic("in pre 数组长度不一致")
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

	panic("数组中没有要找的值")
}
