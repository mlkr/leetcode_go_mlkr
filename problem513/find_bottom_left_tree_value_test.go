package problem513

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	pre, in []int
	ans     int
}{
	{
		[]int{2, 1, 3},
		[]int{1, 2, 3},
		1,
	},

	{
		[]int{1, 2, 4, 3, 5, 7, 6},
		[]int{4, 2, 1, 7, 5, 3, 6},
		7,
	},
}

func Test_findBottomLeftValue(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		root := preIn2Tree(q.pre, q.in)
		ast.Equal(q.ans, findBottomLeftValue(root), "输入 %v %v", q.pre, q.in)
	}
}

func Benchmark_findBottomLeftValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			root := preIn2Tree(q.pre, q.in)
			findBottomLeftValue(root)
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

	root := &TreeNode{
		Val: pre[0],
	}

	if size == 1 {
		return root
	}

	index := getIndex(root.Val, in)

	root.Left = preIn2Tree(pre[1:index+1], in[:index])
	root.Right = preIn2Tree(pre[index+1:], in[index+1:])

	return root
}

func getIndex(val int, nums []int) int {
	for k, v := range nums {
		if v == val {
			return k
		}
	}

	panic("数组中没有对应值！")
}
