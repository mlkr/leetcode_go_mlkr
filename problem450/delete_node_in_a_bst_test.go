package problem450

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	pre []int
	key int
	ans [][]int
}{
	{
		[]int{5, 3, 2, 4, 6, 7},
		3,
		[][]int{
			{5, 2, 4, 6, 7},
			{5, 4, 2, 6, 7},
		},
	},

	{
		[]int{5, 3, 2, 4, 6, 7},
		6,
		[][]int{
			{5, 3, 2, 4, 7},
		},
	},

	{
		[]int{5, 4, 2, 6, 7},
		4,
		[][]int{
			{5, 2, 6, 7},
		},
	},

	{
		[]int{5, 4, 2, 6, 9, 7},
		6,
		[][]int{
			{5, 4, 2, 9, 7},
		},
	},

	{
		[]int{5, 2, 4, 6, 7},
		0,
		[][]int{
			{5, 2, 4, 6, 7},
		},
	},

	{
		[]int{},
		0,
		[][]int{},
	},
}

func Test_deleteNode(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		root := preToBST(q.pre)
		root1 := deleteNode(root, q.key)

		if root == nil {
			ast.Nil(root, "输入为nil")
			return
		}

		ans := []*TreeNode{}
		for _, a := range q.ans {
			ans = append(ans, preToBST(a))
		}
		ast.Contains(ans, root1, "输入 %v", q.pre)
	}
}

func Test_deleteNode2(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		root := preToBST(q.pre)
		root1 := deleteNode2(root, q.key)

		if root == nil {
			ast.Nil(root, "输入为nil")
			return
		}

		ans := []*TreeNode{}
		for _, a := range q.ans {
			ans = append(ans, preToBST(a))
		}
		ast.Contains(ans, root1, "输入 %v", q.pre)
	}
}

func Benchmark_deleteNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			deleteNode(preToBST(q.pre), q.key)
		}
	}
}

func Benchmark_deleteNode2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			deleteNode2(preToBST(q.pre), q.key)
		}
	}
}
