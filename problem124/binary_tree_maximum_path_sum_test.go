package problem124

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxPathSum(t *testing.T) {
	ast := assert.New(t)

	questions := []struct {
		pre, in []int
		pathSum int
	}{
		{
			[]int{5, 4, 11, 7, 2, 8, 13, 4, 1},
			[]int{7, 11, 2, 4, 5, 13, 8, 4, 1},
			48,
		},
		{
			[]int{1, 2, 3},
			[]int{2, 1, 3},
			6,
		},

		{
			[]int{-10, 9, 20, 15, 7},
			[]int{9, -10, 15, 20, 7},
			42,
		},

		{
			[]int{20, 9, -10, 15, 7},
			[]int{9, 20, 15, -10, 7},
			34,
		},
	}

	for _, q := range questions {
		ast.Equal(maxPathSum(prein2Tree(q.pre, q.in)), q.pathSum, "输入 %v", q)
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

	panic("数组中没有要找的值!")
}
