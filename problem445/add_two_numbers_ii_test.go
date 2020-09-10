package problem445

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	l1, l2, ans []byte
}{
	{
		[]byte{5},
		[]byte{5},
		[]byte{1, 0},
	},

	{
		[]byte{7, 2, 4, 3},
		[]byte{5, 6, 4},
		[]byte{7, 8, 0, 7},
	},

	{
		[]byte{},
		[]byte{5, 6, 4},
		[]byte{5, 6, 4},
	},
}

func Test_addTwoNumbers(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		list1 := num2List(q.l1)
		list2 := num2List(q.l2)
		res := addTwoNumbers(list1, list2)

		ast.Equal(q.ans, list2Num(res), "输入 %v %v", q.l1, q.l2)
	}
}

func Benchmark_addTwoNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			list1 := num2List(q.l1)
			list2 := num2List(q.l2)
			res := addTwoNumbers(list1, list2)
			list2Num(res)
		}
	}
}

func num2List(nums []byte) *ListNode {
	var head *ListNode
	head = nil
	for i := len(nums) - 1; i >= 0; i-- {
		node := &ListNode{Val: int(nums[i])}
		node.Next = head
		head = node
	}

	return head
}
