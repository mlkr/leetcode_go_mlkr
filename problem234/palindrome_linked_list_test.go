package problem234

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  bool
}{
	{
		[]int{1, 2},
		false,
	},

	{
		[]int{1, 2, 2, 1},
		true,
	},
}

func Test_isPalindrome(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		head := nums2list(q.nums)
		ast.Equal(q.ans, isPalindrome(head))
	}
}

func Benchmark_isPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			head := nums2list(q.nums)
			isPalindrome(head)
		}

	}
}

func nums2list(nums []int) *ListNode {
	head := &ListNode{}
	temp := head
	for _, num := range nums {
		temp.Next = &ListNode{
			Val: num,
		}

		temp = temp.Next
	}

	return head.Next
}
