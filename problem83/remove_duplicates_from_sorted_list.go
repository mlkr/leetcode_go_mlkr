package problem83

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if head.Val == head.Next.Val {
		val := head.Val

		for {
			if head.Next == nil || head.Next.Val > val {
				break
			}
			head = head.Next
		}
	}

	head.Next = deleteDuplicates(head.Next)

	return head
}
