package problem235

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	pre, in   []int
	p, q, ans int
}{
	{
		[]int{6, 2, 0, 4, 3, 5, 8, 7, 9},
		[]int{0, 2, 3, 4, 5, 6, 7, 8, 9},
		2,
		8,
		6,
	},

	{
		[]int{6, 2, 0, 4, 3, 5, 8, 7, 9},
		[]int{0, 2, 3, 4, 5, 6, 7, 8, 9},
		2,
		4,
		2,
	},

	{
		[]int{6, 2, 0, 4, 3, 5, 8, 7, 9},
		[]int{0, 2, 3, 4, 5, 6, 7, 8, 9},
		7,
		9,
		8,
	},
}

func Test_lowestCommonAncestor(t *testing.T) {
	ast := assert.New(t)
	for _, que := range questions {
		root := prein2Tree(que.pre, que.in)

		p := &TreeNode{
			Val: que.p,
		}

		q := &TreeNode{
			Val: que.q,
		}

		node := lowestCommonAncestor(root, p, q)
		ast.Equal(que.ans, node.Val)
	}
}

func Benchmark_lowestCommonAncestor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, que := range questions {
			root := prein2Tree(que.pre, que.in)

			p := &TreeNode{
				Val: que.p,
			}

			q := &TreeNode{
				Val: que.q,
			}

			lowestCommonAncestor(root, p, q)
		}
	}
}

func prein2Tree(pre, in []int) *TreeNode {
	preLen, inLen := len(pre), len(in)
	if preLen != inLen {
		panic("pre, in 数组长度不一致")
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

	panic("数组中没有找到值")
}
