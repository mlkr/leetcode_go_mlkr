package problem404

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	pre, in []int
	ans     int
}{
	{
		[]int{3, 9, 20, 15, 7},
		[]int{9, 3, 15, 20, 7},
		24,
	},

	{
		[]int{1},
		[]int{1},
		0,
	},

	{
		[]int{3, 9, 15, 7, 20},
		[]int{15, 9, 7, 3, 20},
		15,
	},
}

func Test_sumOfLeftLeaves(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		root := preIn2Tree(q.pre, q.in)
		ast.Equal(q.ans, sumOfLeftLeaves(root), "输入 %d %d", q.pre, q.in)
	}
}

func preIn2Tree(pre, in []int) *TreeNode {
	if len(pre) != len(in) {
		panic("pre len 长度不一致!")
	}

	if len(pre) == 0 {
		return nil
	}

	root := &TreeNode{Val: pre[0]}
	if len(pre) == 1 {
		return root
	}

	idx := indexOf(pre[0], in)
	root.Left = preIn2Tree(pre[1:idx+1], in[:idx])
	root.Right = preIn2Tree(pre[idx+1:], in[idx+1:])

	return root
}

func indexOf(n int, nums []int) int {
	for k, v := range nums {
		if v == n {
			return k
		}
	}

	panic("数组中没有要找的数")
}
