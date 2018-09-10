package problem106

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	in, post []int
}

type ans struct {
	level [][]int
}

type question struct {
	para
	ans
}

func TestBuildTree(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{
				[]int{9, 3, 15, 20, 7},
				[]int{9, 15, 7, 20, 3},
			},
			ans{
				[][]int{
					[]int{3},
					[]int{9, 20},
					[]int{15, 7},
				},
			},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(getLevelOrder(buildTree(para.in, para.post)), ans.level, "输入 %v", para)
	}
}

func getLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}

	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if len(res) == level {
			res = append(res, []int{})
		}

		res[level] = append(res[level], root.Val)

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return res
}
