package problem144

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	pre []int
	in  []int
}{
	{
		[]int{1, 2, 3},
		[]int{1, 3, 2},
	},

	{
		[]int{3, 1, 2},
		[]int{1, 3, 2},
	},
}

func Test_preorderTraversal(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		root := prein2tree(q.pre, q.in)
		ast.Equal(preorderTraversal(root), q.pre)
	}
}

func Test_preorderTraversal2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		root := prein2tree(q.pre, q.in)
		ast.Equal(preorderTraversal2(root), q.pre)
	}
}

func Benchmark_preorderTraversal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			root := prein2tree(q.pre, q.in)
			preorderTraversal(root)
		}
	}
}

func Benchmark_preorderTraversal2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			root := prein2tree(q.pre, q.in)
			preorderTraversal2(root)
		}
	}
}

func prein2tree(pre, in []int) *TreeNode {
	preLen := len(pre)
	inLen := len(in)

	if preLen != inLen {
		panic("pre in 数组长度不一致")
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

	index := indexOf(root.Val, in)

	root.Left = prein2tree(pre[1:index+1], in[:index])
	root.Right = prein2tree(pre[index+1:], in[index+1:])

	return root
}

func indexOf(n int, nums []int) int {
	for k, v := range nums {
		if n == v {
			return k
		}
	}

	panic("数组中没有找到要找的值!")
}
