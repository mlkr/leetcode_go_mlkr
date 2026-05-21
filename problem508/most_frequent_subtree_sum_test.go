package problem508

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	pre, in, ans []int
}{
	{
		[]int{5, 2, -3},
		[]int{2, 5, -3},
		[]int{2, -3, 4},
	},

	{
		[]int{5, 2, -5},
		[]int{2, 5, -5},
		[]int{2},
	},
}

func Test_findFrequentTreeSum(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		root := preIn2Tree(q.pre, q.in)
		act := findFrequentTreeSum(root)
		sort.Ints(act)
		sort.Ints(q.ans)
		ast.Equal(q.ans, act, "输入 %v", act)
	}

}

func Benchmark_findFrequentTreeSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			root := preIn2Tree(q.pre, q.in)
			findFrequentTreeSum(root)
		}
	}
}

// DLR
// LDR
func preIn2Tree(pre, in []int) *TreeNode {
	size := len(pre)
	if len(in) != size {
		panic("pre in 长度不一致！")
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

	index := getIndex(pre[0], in)

	root.Left = preIn2Tree(pre[1:index+1], in[:index])
	root.Right = preIn2Tree(pre[index+1:], in[index+1:])

	return root
}

func getIndex(val int, in []int) int {
	for k, v := range in {
		if v == val {
			return k
		}
	}

	panic("没有找到 index")
}
