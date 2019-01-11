package problem206

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var question = []struct {
	para []int
	ans  []int
}{
	{
		[]int{1, 2, 3, 4, 5},
		[]int{5, 4, 3, 2, 1},
	},
}

func Test_reverseList(t *testing.T) {
	ast := assert.New(t)
	for _, q := range question {
		head := s2l(q.para)
		newHead := reverseList(head)
		ast.Equal(l2s(newHead), q.ans, "输入 %q", q.para)
	}
}

func Test_reverseList2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range question {
		head := s2l(q.para)
		newHead := reverseList2(head)
		ast.Equal(l2s(newHead), q.ans, "输入 %q", q.para)
	}
}

func Benchmark_reverseList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range question {
			head := s2l(q.para)
			reverseList(head)
		}
	}
}

func Benchmark_reverseList2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range question {
			head := s2l(q.para)
			reverseList2(head)
		}
	}
}

func l2s(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

func s2l(nums []int) *ListNode {
	temp := &ListNode{}
	head := temp
	for _, num := range nums {
		head.Next = &ListNode{
			Val: num,
		}

		head = head.Next
	}

	return temp.Next
}
