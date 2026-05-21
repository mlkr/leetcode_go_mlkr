package problem662

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	str string
	ans int
}{
	{
		"1,3,2,5,3,null,9",
		4,
	},

	{
		"1,3,2,5,null,null,9,6,null,7",
		7,
	},

	{
		"1,3,2,5",
		2,
	},
}

func Test_widthOfBinaryTree(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		root := str2Tree(q.str)
		ast.Equal(q.ans, widthOfBinaryTree(root))
	}
}

func str2Tree(str string) *TreeNode {
	ss := strings.Split(str, ",")

	if ss[0] == "null" {
		return nil
	}

	root := &TreeNode{}
	root.Val, _ = strconv.Atoi(ss[0])

	nodes := []*TreeNode{root}
	index := 1
	for len(nodes) > 0 && index < len(ss) {
		now := nodes[0]
		nodes = nodes[1:]

		if ss[index] != "null" {
			node := &TreeNode{}
			node.Val, _ = strconv.Atoi(ss[index])
			now.Left = node
			nodes = append(nodes, node)
		}

		index++
		if index == len(ss) {
			break
		}

		if ss[index] != "null" {
			node := &TreeNode{}
			node.Val, _ = strconv.Atoi(ss[index])
			now.Right = node
			nodes = append(nodes, node)
		}

		index++
	}

	return root
}
