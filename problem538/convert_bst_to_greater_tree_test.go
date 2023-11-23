package problem538

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	pre, in []int
	ans     []int
}{
	{
		[]int{4, 1, 0, 2, 3, 6, 5, 7, 8},
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 8},
		[]int{8, 15, 21, 26, 30, 33, 35, 36, 36},
	},
}

func Test_convertBST(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		root := preIn2Tree(q.pre, q.in)
		root = convertBST(root)
		nums := getRDL(root)
		ast.Equal(q.ans, nums)
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
	if index == -1 {
		panic("pre in 元素不一致！")
	}

	root.Left = preIn2Tree(pre[1:index+1], in[:index])
	root.Right = preIn2Tree(pre[index+1:], in[index+1:])

	return root
}

func getIndex(num int, in []int) int {
	for k, v := range in {
		if v == num {
			return k
		}
	}

	return -1
}

func getRDL(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	nums := getRDL(root.Right)
	nums = append(nums, root.Val)
	nums = append(nums, getRDL(root.Left)...)

	return nums
}
