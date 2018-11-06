package problem145

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	pre  []int
	in   []int
	post []int
}{
	{
		[]int{1, 2, 3},
		[]int{1, 3, 2},
		[]int{3, 2, 1},
	},

	{
		[]int{1},
		[]int{1},
		[]int{1},
	},
}

func Test_postorderTraversal(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		root := prein2tree(q.pre, q.in)
		post := postorderTraversal(root)
		ast.Equal(post, q.post)
	}
}

func Test_postorderTraversal2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		root := prein2tree(q.pre, q.in)
		post := postorderTraversal2(root)
		ast.Equal(post, q.post)
	}
}

func Benchmark_postorderTraversal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			root := prein2tree(q.pre, q.in)
			postorderTraversal(root)
		}
	}
}

func prein2tree(pre, in []int) *TreeNode {
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

	idx := indexOf(root.Val, in)
	root.Left = prein2tree(pre[1:idx+1], in[:idx])
	root.Right = prein2tree(pre[idx+1:], in[idx+1:])

	return root
}

func indexOf(n int, nums []int) int {
	for k, v := range nums {
		if v == n {
			return k
		}
	}

	panic("数组中没有找到要找的值!")
}
