package problem199

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	pre, in, ans []int
}{
	{
		[]int{1, 2, 5, 3, 4},
		[]int{2, 5, 1, 3, 4},
		[]int{1, 3, 4},
	},

	{
		[]int{1, 2, 4, 3},
		[]int{4, 2, 1, 3},
		[]int{1, 3, 4},
	},
}

func Test_rightSideView(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		root := prein2Tree(q.pre, q.in)
		ast.Equal(rightSideView(root), q.ans)
	}
}

func Benchmark_rightSideView(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			root := prein2Tree(q.pre, q.in)
			rightSideView(root)
		}
	}
}

func prein2Tree(pre, in []int) *TreeNode {
	preLen := len(pre)
	inLen := len(in)

	if preLen != inLen || preLen == 0 {
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

func indexOf(num int, nums []int) int {
	for k, v := range nums {
		if v == num {
			return k
		}
	}

	panic("数组中没有要找的值!")
}
