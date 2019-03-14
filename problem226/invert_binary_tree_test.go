package problem226

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	pre, in, aPre, ain []int
}{
	{
		[]int{4, 2, 1, 3, 7, 6, 9},
		[]int{1, 2, 3, 4, 6, 7, 9},
		[]int{4, 7, 9, 6, 2, 3, 1},
		[]int{9, 7, 6, 4, 3, 2, 1},
	},
}

func Test_invertTree(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		root := prein2Tree(q.pre, q.in)
		root = invertTree(root)
		invertRoot := prein2Tree(q.aPre, q.ain)
		ast.True(reflect.DeepEqual(root, invertRoot))
	}
}

func Benchmark_invertTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			root := prein2Tree(q.pre, q.in)
			invertTree(root)
		}
	}
}

func prein2Tree(pre, in []int) *TreeNode {
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

	panic("数组中没有找到值!")
}
