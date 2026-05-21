package problem515

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	pre, in []int
	ans     []int
}{
	{
		[]int{},
		[]int{},
		nil,
	},

	{
		[]int{1, 3, 5, 3, 2, 9},
		[]int{5, 3, 3, 1, 2, 9},
		[]int{1, 3, 9},
	},

	{
		[]int{1, 3, 2},
		[]int{3, 1, 2},
		[]int{1, 3},
	},
}

func Test_largestValues(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		root := preIn2Tree(q.pre, q.in)
		values := largestValues(root)
		ast.Equal(q.ans, values, "%v", values)
	}
}

func Benchmark_largestValues(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			root := preIn2Tree(q.pre, q.in)
			largestValues(root)
		}
	}
}

func preIn2Tree(pre, in []int) *TreeNode {
	size := len(pre)

	if size != len(in) {
		panic("pre in 数组长度不一致！")
	}

	if size == 0 {
		return nil
	}

	root := &TreeNode{Val: pre[0]}
	if size == 1 {
		return root
	}

	index := getIndex(root.Val, in)

	root.Left = preIn2Tree(pre[1:index+1], in[:index])
	root.Right = preIn2Tree(pre[index+1:], in[index+1:])

	return root
}

func getIndex(val int, in []int) int {
	for index, v := range in {
		if v == val {
			return index
		}
	}

	panic("in 数组没有找到对应值！")
}
