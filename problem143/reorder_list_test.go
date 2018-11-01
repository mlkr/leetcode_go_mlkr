package problem143

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para []int
	ans  []int
}{
	{
		[]int{1, 2, 3, 4},
		[]int{1, 4, 2, 3},
	},

	{
		[]int{1, 2, 3, 4, 5},
		[]int{1, 5, 2, 4, 3},
	},
}

func nums2link(nums []int) *ListNode {
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

func link2nums(head *ListNode) []int {
	res := []int{}
	for {
		if head == nil {
			break
		}

		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

func Test_link2nums(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		head := nums2link(q.para)
		nums := link2nums(head)
		ast.Equal(q.para, nums)
	}
}

func Test_reorderList(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		head := nums2link(q.para)
		reorderList(head)
		nums := link2nums(head)
		ast.Equal(q.ans, nums)
	}
}

func Test_reorderList2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		head := nums2link(q.para)
		reorderList2(head)
		nums := link2nums(head)
		ast.Equal(q.ans, nums)
	}
}

func Benchmark_reorderList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			head := nums2link(q.para)
			reorderList(head)
		}
	}
}

func Benchmark_reorderList2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			head := nums2link(q.para)
			reorderList2(head)
		}
	}
}
