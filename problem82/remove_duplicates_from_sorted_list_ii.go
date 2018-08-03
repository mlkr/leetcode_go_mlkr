package problem82

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
			head = head.Next
			if head == nil || head.Val > val {
				break
			}
		}

		return deleteDuplicates(head)
	}

	head.Next = deleteDuplicates(head.Next)

	return head
}
