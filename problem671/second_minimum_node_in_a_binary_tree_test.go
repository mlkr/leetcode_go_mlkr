package problem671

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s   string
	ans int
}{
	{
		"1,1,3,1,1,3,4,3,1,1,1,3,8,4,8,3,3,1,6,2,1",
		2,
	},

	{
		"5,8,5",
		8,
	},

	{
		"2,2,5,null,null,5,7",
		5,
	},

	{
		"2,2,2,null,null,2,7",
		7,
	},

	{
		"2,2,2",
		-1,
	},
}

func Test_findSecondMinimumValue(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		root := str2Tree(q.s)
		ast.Equal(q.ans, findSecondMinimumValue(root))
	}
}

func BenchMark_findSecondMinimumValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			root := str2Tree(q.s)
			findSecondMinimumValue(root)
		}
	}
}

func str2Tree(s string) *TreeNode {
	ss := strings.Split(s, ",")
	if len(ss) == 0 || ss[0] == "null" {
		return nil
	}

	n, _ := strconv.Atoi(ss[0])
	root := &TreeNode{Val: n}
	nodes := []*TreeNode{root}

	index := 1
	for len(nodes) > 0 && index < len(ss) {
		now := nodes[0]
		nodes = nodes[1:]

		if ss[index] != "null" {
			leftVal, _ := strconv.Atoi(ss[index])
			left := &TreeNode{Val: leftVal}
			now.Left = left
			nodes = append(nodes, left)
		}
		index++

		if ss[index] != "null" {
			rightVal, _ := strconv.Atoi(ss[index])
			right := &TreeNode{Val: rightVal}
			now.Right = right
			nodes = append(nodes, right)
		}
		index++

	}

	return root
}
