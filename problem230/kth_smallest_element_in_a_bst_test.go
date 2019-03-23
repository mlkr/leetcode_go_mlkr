package problem230

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	pre, in []int
	k, ans  int
}{
	{
		[]int{3, 1, 2, 4},
		[]int{1, 2, 3, 4},
		1,
		1,
	},

	{
		[]int{3, 1, 2, 4},
		[]int{1, 2, 3, 4},
		4,
		4,
	},

	{
		[]int{5, 3, 2, 1, 4, 6},
		[]int{1, 2, 3, 4, 5, 6},
		3,
		3,
	},
}

func Test_kthSmallest(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		root := prein2Tree(q.pre, q.in)
		ast.Equal(q.ans, kthSmallest(root, q.k))
	}
}

func Test_kthSmallest2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		root := prein2Tree(q.pre, q.in)
		ast.Equal(q.ans, kthSmallest2(root, q.k))
	}
}

func Test_kthSmallest3(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		root := prein2Tree(q.pre, q.in)
		ast.Equal(q.ans, kthSmallest3(root, q.k))
	}
}

func Benchmark_kthSmallest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			root := prein2Tree(q.pre, q.in)
			kthSmallest(root, q.k)
		}
	}
}

func Benchmark_kthSmallest2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			root := prein2Tree(q.pre, q.in)
			kthSmallest2(root, q.k)
		}
	}
}

func Benchmark_kthSmallest3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			root := prein2Tree(q.pre, q.in)
			kthSmallest3(root, q.k)
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

	root.Left = prein2Tree(pre[1:idx+1], in[0:idx])
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
