package problem669

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	rootStr    string
	low, hight int
	ans        string
}{
	{
		"1,0,2",
		1, 2,
		"1,null,2",
	},

	{
		"3,0,4,null,2,null,null,1",
		1, 3,
		"3,2,null,1",
	},
}

func Test_trimBST(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		root := str2Tree(q.rootStr)
		trimRoot := trimBST(root, q.low, q.hight)
		trimStr := tree2Str(trimRoot)
		ast.Equal(q.ans, trimStr)
	}
}

func str2Tree(s string) *TreeNode {
	ss := strings.Split(s, ",")
	nodes := make([]*TreeNode, 1, len(ss))

	rootVal, _ := strconv.Atoi(ss[0])
	root := &TreeNode{Val: rootVal}
	nodes[0] = root
	ss = ss[1:]

	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]

		left := ss[0]
		if left != "null" {
			leftVal, _ := strconv.Atoi(left)
			node.Left = &TreeNode{Val: leftVal}
			nodes = append(nodes, node.Left)
		}

		ss = ss[1:]
		if len(ss) == 0 {
			break
		}

		right := ss[0]
		if right != "null" {
			rightVal, _ := strconv.Atoi(right)
			node.Right = &TreeNode{Val: rightVal}
			nodes = append(nodes, node.Right)
		}

		ss = ss[1:]
		if len(ss) == 0 {
			break
		}

	}

	return root
}

func tree2Str(root *TreeNode) string {
	ss := []string{}
	nodes := []*TreeNode{root}

	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]

		if node == nil {
			ss = append(ss, "null")
			continue
		}

		ss = append(ss, strconv.Itoa(node.Val))

		nodes = append(nodes, node.Left)
		nodes = append(nodes, node.Right)
	}

	i := len(ss) - 1
	for ss[i] == "null" {
		i--
	}
	ss = ss[:i+1]

	res := strings.Join(ss, ",")

	return res
}
