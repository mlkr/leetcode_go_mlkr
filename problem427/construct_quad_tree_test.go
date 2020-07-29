package problem427

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	grid [][]int
	ans  [][]int
}{
	{
		[][]int{
			{1, 1, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 1, 0, 0, 0, 0},
		},
		[][]int{
			{0, 1},
			{1, 1},
			{0, 1},
			{1, 1},
			{1, 0},
			{1, 0},
			{1, 0},
			{1, 1},
			{1, 1},
		},
	},

	{
		[][]int{
			{0, 1},
			{1, 0},
		},
		[][]int{
			{0, 1},
			{1, 0},
			{1, 1},
			{1, 1},
			{1, 0},
		},
	},

	{
		[][]int{},
		[][]int{},
	},
}

func Test_construct(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		root := construct(q.grid)
		res := bfs(root)
		ast.Equal(q.ans, res, "输入 %v", q.grid)
	}
}

func bfs(root *Node) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	queue := []*Node{}
	queue = append(queue, root)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		isLeaf, val := 0, 0
		if node.IsLeaf {
			isLeaf = 1
		}
		if node.Val {
			val = 1
		}
		res = append(res, []int{isLeaf, val})

		if node.IsLeaf {
			continue
		}

		queue = append(queue, node.TopLeft)
		queue = append(queue, node.TopRight)
		queue = append(queue, node.BottomLeft)
		queue = append(queue, node.BottomRight)
	}

	return res
}
