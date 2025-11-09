package problem653

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	str string
	k   int
	ans bool
}{
	{
		"2,1,3",
		4,
		true,
	},

	{
		"2,null,3",
		6,
		false,
	},

	{
		"1",
		2,
		false,
	},

	{
		"5,3,6",
		9,
		true,
	},

	{
		"5,3,6,2,4,null,7",
		9,
		true,
	},

	{
		"5,3,6,2,4,null,7",
		28,
		false,
	},
}

func Test_findTarget(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		root := str2Tree(q.str)
		ast.Equal(q.ans, findTarget(root, q.k), "输入 %v %v", q.str, q.k)
	}
}

func Benchmark_findTarget(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			root := str2Tree(q.str)
			findTarget(root, q.k)
		}
	}
}

type Queue struct {
	nodes []*TreeNode
}

func NewQue() *Queue {
	return &Queue{
		nodes: make([]*TreeNode, 0),
	}
}

func (this *Queue) EnQue(node *TreeNode) {
	this.nodes = append(this.nodes, node)
}

func (this *Queue) isEmpty() bool {
	return len(this.nodes) == 0
}

func (this *Queue) deQue() *TreeNode {
	if this.isEmpty() {
		return nil
	}

	node := this.nodes[0]
	this.nodes = this.nodes[1:]

	return node
}

func str2Tree(str string) *TreeNode {
	str = strings.TrimSpace(str)
	if len(str) == 0 || str == "null" {
		return nil
	}

	ss := strings.Split(str, ",")

	root := &TreeNode{}
	root.Val, _ = strconv.Atoi(ss[0])
	que := NewQue()
	que.EnQue(root)

	index := 1
	for index < len(ss) && !que.isEmpty() {
		now := que.deQue()

		// 左子树
		if ss[index] == "null" {
			now.Left = nil
		} else {
			left := &TreeNode{}
			left.Val, _ = strconv.Atoi(ss[index])
			now.Left = left
			que.EnQue(left)
		}

		index++
		if index == len(ss) {
			break
		}

		// 右子树
		if ss[index] == "null" {
			now.Right = nil
		} else {
			right := &TreeNode{}
			right.Val, _ = strconv.Atoi(ss[index])
			now.Right = right
			que.EnQue(right)
		}

		index++

	}

	return root
}

func tree2str(root *TreeNode) string {
	if root == nil {
		return "null"
	}

	que := NewQue()
	que.EnQue(root)

	res := []string{}
	for !que.isEmpty() {
		now := que.deQue()
		if now == nil {
			res = append(res, "null")
		} else {
			res = append(res, strconv.Itoa(now.Val))
			que.EnQue(now.Left)
			que.EnQue(now.Right)
		}
	}

	// 去掉尾部多余的"null"
	i := len(res) - 1
	for res[i] == "null" {
		i--
	}
	res = res[:i+1]

	return strings.Join(res, ",")
}
