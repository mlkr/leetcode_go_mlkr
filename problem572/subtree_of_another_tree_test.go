package problem572

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	pre1, in1 []int
	pre2, in2 []int
	ans       bool
}{
	{
		[]int{3, 4, 1, 2, 5},
		[]int{1, 4, 2, 3, 5},
		[]int{4, 1, 2},
		[]int{1, 4, 2},
		true,
	},

	{
		[]int{3, 4, 1, 2, 0, 5},
		[]int{1, 4, 0, 2, 3, 5},
		[]int{4, 1, 2},
		[]int{1, 4, 2},
		false,
	},
}

func Test_isSubtree(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		root := preIn2Tree(q.pre1, q.in1)
		subRoot := preIn2Tree(q.pre2, q.in2)
		ast.Equal(q.ans, isSubtree(root, subRoot))
	}
}

func Benchmark_isSubtree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			root := preIn2Tree(q.pre1, q.in1)
			subRoot := preIn2Tree(q.pre2, q.in2)
			isSubtree(root, subRoot)
		}

	}
}

func preIn2Tree(pre, in []int) *TreeNode {
	if len(pre) != len(in) {
		panic("pre in 长度不一致！")
	}

	if len(pre) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: pre[0],
	}

	index := getIndex(root.Val, in)

	root.Left = preIn2Tree(pre[1:index+1], in[:index])
	root.Right = preIn2Tree(pre[index+1:], in[index+1:])

	return root
}

func getIndex(num int, in []int) int {
	for i, n := range in {
		if num == n {
			return i
		}
	}

	panic("in 中找不到对应值！")
}
