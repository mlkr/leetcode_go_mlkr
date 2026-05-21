package problem654

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  string
}{
	{
		[]int{3, 2, 1, 6, 0, 5},
		"6,3,5,null,2,0,null,null,1",
	},
}

func Test_constructMaximumBinaryTree(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		root := constructMaximumBinaryTree(q.nums)
		ast.Equal(q.ans, tree2str(root))
	}
}

func tree2str(root *TreeNode) string {
	nodes := []*TreeNode{}
	nodes = append(nodes, root)

	ss := []string{}
	for len(nodes) > 0 {
		now := nodes[0]
		nodes = nodes[1:]

		if now == nil {
			ss = append(ss, "null")
			continue
		}

		ss = append(ss, strconv.Itoa(now.Val))
		nodes = append(nodes, now.Left)
		nodes = append(nodes, now.Right)
	}

	i := len(ss) - 1
	for ss[i] == "null" {
		i--
	}
	ss = ss[:i+1]

	return strings.Join(ss, ",")
}
