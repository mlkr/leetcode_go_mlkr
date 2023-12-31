package problem543

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	pre, in []int
	ans     int
}{
	{
		[]int{1, 2},
		[]int{1, 2},
		1,
	},

	{
		[]int{1, 2, 4, 5, 3},
		[]int{4, 2, 5, 1, 3},
		3,
	},
}

func Test_diameterOfBinaryTree(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		root := preIn2Tree(q.pre, q.in)
		ast.Equal(q.ans, diameterOfBinaryTree(root))
	}
}

func preIn2Tree(pre, in []int) *TreeNode {
	if len(pre) != len(in) {
		panic("pre in 长度不一致！")
	}

	if len(pre) == 0 {
		return nil
	}

	root := &TreeNode{Val: pre[0]}

	index := indexOf(root.Val, in)
	root.Left = preIn2Tree(pre[1:index+1], in[:index])
	root.Right = preIn2Tree(pre[index+1:], in[index+1:])

	return root
}

func indexOf(num int, nums []int) int {
	for i := 0; i < len(nums); i++ {
		if num == nums[i] {
			return i
		}
	}

	panic("没有找到值！")
}
