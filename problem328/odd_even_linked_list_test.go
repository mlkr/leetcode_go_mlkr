package problem328

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para, ans []int
}{
	{
		[]int{1, 2, 3, 4, 5},
		[]int{1, 3, 5, 2, 4},
	},

	{
		[]int{1, 2, 3, 4, 5, 6},
		[]int{1, 3, 5, 2, 4, 6},
	},

	{
		[]int{2, 1, 3, 5, 6, 4, 7},
		[]int{2, 3, 6, 7, 1, 5, 4},
	},

	{
		[]int{},
		[]int{},
	},

	{
		[]int{1},
		[]int{1},
	},
}

func Test_oddEvenList(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		head := arrToList(q.para)
		newHead := oddEvenList(head)
		ast.Equal(q.ans, listToArr(newHead), "输入 %v", q.para)
	}
}

func Benchmark_oddEvenList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			head := arrToList(q.para)
			newHead := oddEvenList(head)
			listToArr(newHead)
		}
	}
}

func arrToList(nums []int) *ListNode {
	head := new(ListNode)
	temp := head
	for _, num := range nums {
		head.Next = &ListNode{
			Val: num,
		}

		head = head.Next
	}

	return temp.Next
}

func listToArr(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}
