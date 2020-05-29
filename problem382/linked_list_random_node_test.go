package problem382

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Solution(t *testing.T) {
	ast := assert.New(t)

	head := ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	solution := Constructor(&head)
	ast.Contains([]int{1, 2, 3}, solution.GetRandom())
	ast.Contains([]int{1, 2, 3}, solution.GetRandom())
	ast.Contains([]int{1, 2, 3}, solution.GetRandom())
}

func Benchmark_Solution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		head := ListNode{Val: 1}
		head.Next = &ListNode{Val: 2}
		head.Next.Next = &ListNode{Val: 3}
		solution := Constructor(&head)
		solution.GetRandom()
	}
}
