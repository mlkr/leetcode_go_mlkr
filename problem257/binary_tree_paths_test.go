package problem257

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	pre, in []int
	ans     []string
}{
	{
		[]int{},
		[]int{},
		nil,
	},

	{
		[]int{1, 2, 5, 3},
		[]int{2, 5, 1, 3},
		[]string{"1->2->5", "1->3"},
	},
}

func Test_binaryTreePaths(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		root := prein2Tree(q.pre, q.in)
		ast.Equal(q.ans, binaryTreePaths(root))
	}
}

func prein2Tree(pre, in []int) *TreeNode {
	preLen := len(pre)
	inLen := len(in)
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

	i := indexOf(root.Val, in)
	root.Left = prein2Tree(pre[1:i+1], in[:i])
	root.Right = prein2Tree(pre[i+1:], in[i+1:])

	return root
}

func indexOf(num int, nums []int) int {
	for k, v := range nums {
		if num == v {
			return k
		}
	}

	panic("数组中没有找到值!")
}
