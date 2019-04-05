package problem236

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	pre, in   []int
	p, q, ans int
}{
	{
		[]int{3, 5, 6, 2, 7, 4, 1, 0, 8},
		[]int{6, 5, 7, 2, 4, 3, 0, 1, 8},
		5,
		1,
		3,
	},

	{
		[]int{3, 5, 6, 2, 7, 4, 1, 0, 8},
		[]int{6, 5, 7, 2, 4, 3, 0, 1, 8},
		5,
		4,
		5,
	},
}

func Test_lowestCommonAncestor(t *testing.T) {
	ast := assert.New(t)
	for _, que := range questions {
		root := prein2Tree(que.pre, que.in)
		p, q := &TreeNode{Val: que.p}, &TreeNode{Val: que.q}
		node := lowestCommonAncestor(root, p, q)
		ast.Equal(que.ans, node.Val)
	}
}

func Benchmark_lowestCommonAncestor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, que := range questions {
			root := prein2Tree(que.pre, que.in)
			p, q := &TreeNode{Val: que.p}, &TreeNode{Val: que.q}
			lowestCommonAncestor(root, p, q)
		}
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

func indexOf(num int, nums []int) int {
	for k, v := range nums {
		if v == num {
			return k
		}
	}

	panic("数组中没有要找的值!")
}
