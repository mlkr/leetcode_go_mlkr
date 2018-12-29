package problem203

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para []int
	val  int
	ans  []int
}{
	{
		[]int{1, 2, 6, 3, 4, 5, 6},
		6,
		[]int{1, 2, 3, 4, 5},
	},
}

func Test_removeElements(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		head := numsToLink(q.para)
		res := removeElements(head, q.val)
		ans := linkToNums(res)
		ast.Equal(ans, q.ans, "输入 %v", q.para)
	}
}

func numsToLink(nums []int) *ListNode {
	head := &ListNode{}
	temp := head
	for _, num := range nums {
		head.Next = &ListNode{
			Val: num,
		}
		head = head.Next
	}

	return temp.Next
}

func linkToNums(head *ListNode) []int {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	return nums
}
