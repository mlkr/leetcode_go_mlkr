package problem652

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums string
	ans  []string
}{
	{
		"1,2,3,4,null,2,4,null,null,4",
		[]string{"2,4", "4"},
	},

	{
		"2,1,11,11,null,1",
		[]string{},
	},
}

func Test_findDuplicateSubtrees(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		root := str2Tree(q.nums)
		subtrees := findDuplicateSubtrees(root)

		ss := []string{}
		for _, subtree := range subtrees {
			ss = append(ss, tree2Str(subtree))
		}

		ast.ElementsMatch(q.ans, ss)
	}

}

// 层序遍历转二叉树

// 队列实现
type Queue struct {
	nodes []*TreeNode
}

func NewQue() *Queue {
	return &Queue{nodes: make([]*TreeNode, 0)}
}

func (this *Queue) isEmpty() bool {
	if len(this.nodes) == 0 {
		return true
	}

	return false
}

func (this *Queue) Enqueue(root *TreeNode) {
	this.nodes = append(this.nodes, root)
}

func (this *Queue) Dequeue() *TreeNode {
	if this.isEmpty() {
		return nil
	}

	root := this.nodes[0]
	this.nodes = this.nodes[1:]

	return root
}

// "2,1,11,11,null,1"
func str2Tree(str string) *TreeNode {
	str = strings.TrimSpace(str)
	if str == "" || str == "null" {
		return nil
	}

	parts := strings.Split(str, ",")

	rootVal, _ := strconv.Atoi(parts[0])
	root := &TreeNode{Val: rootVal}

	que := NewQue()
	que.Enqueue(root)

	index := 1
	for !que.isEmpty() && index < len(parts) {
		now := que.Dequeue()

		// 左子树
		if parts[index] == "null" {
			now.Left = nil
		} else {
			leftVal, _ := strconv.Atoi(parts[index])
			left := &TreeNode{Val: leftVal}
			now.Left = left
			que.Enqueue(left)
		}

		index++
		if index == len(parts) {
			break
		}

		// 右子树
		if parts[index] == "null" {
			now.Right = nil
		} else {
			rightVal, _ := strconv.Atoi(parts[index])
			right := &TreeNode{Val: rightVal}
			now.Right = right
			que.Enqueue(right)
		}

		index++
	}

	return root
}

// "2,1,11,11,null,1"
func tree2Str(root *TreeNode) string {
	if root == nil {
		return "null"
	}

	que := NewQue()
	que.Enqueue(root)

	res := []string{}
	for !que.isEmpty() {
		now := que.Dequeue()
		if now == nil {
			res = append(res, "null")
			continue
		}

		res = append(res, strconv.Itoa(now.Val))

		que.Enqueue(now.Left)
		que.Enqueue(now.Right)
	}

	// 去掉末尾多余的“null”
	i := len(res) - 1
	for res[i] == "null" {
		i--
	}
	res = res[:i+1]

	return strings.Join(res, ",")
}
