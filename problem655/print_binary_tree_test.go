package problem655

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	str string
	ans [][]string
}{
	{
		"3,null,30,10,null,null,15,null,45",
		[][]string{
			[]string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "3", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
			[]string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "30", "", "", "", "", "", "", ""},
			[]string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "10", "", "", "", "", "", "", "", "", "", "", ""},
			[]string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "15", "", "", "", "", "", "", "", "", ""},
			[]string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "45", "", "", "", "", "", "", "", ""},
		},
	},

	{
		"1,2,5,3,null,null,null,4",
		[][]string{
			[]string{"", "", "", "", "", "", "", "1", "", "", "", "", "", "", ""},
			[]string{"", "", "", "2", "", "", "", "", "", "", "", "5", "", "", ""},
			[]string{"", "3", "", "", "", "", "", "", "", "", "", "", "", "", ""},
			[]string{"4", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
		},
	},
	{
		"5,3,6,2,4,null,7",
		[][]string{
			[]string{"", "", "", "5", "", "", ""},
			[]string{"", "3", "", "", "", "6", ""},
			[]string{"2", "", "4", "", "", "", "7"},
		},
	},

	{
		"1,2",
		[][]string{
			[]string{"", "1", ""},
			[]string{"2", "", ""},
		},
	},

	{
		"1,2,3,null,4",
		[][]string{
			[]string{"", "", "", "1", "", "", ""},
			[]string{"", "2", "", "", "", "3", ""},
			[]string{"", "", "4", "", "", "", ""},
		},
	},
}

func Test_printTree(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		root := str2Tree(q.str)
		ast.Equal(q.ans, printTree(root))
	}
}

func Benchmark_printTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			root := str2Tree(q.str)
			printTree(root)
		}
	}
}

func str2Tree(str string) *TreeNode {
	str = strings.TrimSpace(str)
	if str == "null" || len(str) == 0 {
		return nil
	}

	ss := strings.Split(str, ",")
	root := &TreeNode{}
	root.Val, _ = strconv.Atoi(ss[0])

	nodes := []*TreeNode{}
	nodes = append(nodes, root)

	index := 1
	for index < len(ss) && len(nodes) > 0 {
		now := nodes[0]
		nodes = nodes[1:]

		if ss[index] != "null" {
			left := &TreeNode{}
			left.Val, _ = strconv.Atoi(ss[index])
			now.Left = left
			nodes = append(nodes, left)
		}

		index++
		if index == len(ss) {
			break
		}

		if ss[index] != "null" {
			right := &TreeNode{}
			right.Val, _ = strconv.Atoi(ss[index])
			now.Right = right
			nodes = append(nodes, right)
		}
		index++
	}

	return root
}
