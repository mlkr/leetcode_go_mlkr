package problem530

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	pre, in []int
	ans     int
}{
	{
		[]int{4, 2, 1, 3, 6},
		[]int{1, 2, 3, 4, 6},
		1,
	},

	{
		[]int{1, 0, 48, 12, 49},
		[]int{0, 1, 12, 48, 49},
		1,
	},

	{
		[]int{1, 3, 2},
		[]int{1, 2, 3},
		1,
	},

	{
		[]int{236, 104, 227, 701, 911},
		[]int{104, 227, 236, 701, 911},
		9,
	},
}

func Test_getMinimumDifference(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		root := preIn2Tree(q.pre, q.in)
		ast.Equal(q.ans, getMinimumDifference(root))
	}
}

func Benchmark_getMinimumDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			root := preIn2Tree(q.pre, q.in)
			getMinimumDifference(root)
		}
	}
}

// pre DLR
// 4 2 1 3 6

// in LDR
// 1 2 3 4 6
// index=3

func preIn2Tree(pre, in []int) *TreeNode {
	if len(pre) != len(in) {
		panic("pre in 切片长度不一致！")
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

func getIndex(val int, in []int) int {
	for i, v := range in {
		if v == val {
			return i
		}
	}

	panic("没有找到 index")
}
