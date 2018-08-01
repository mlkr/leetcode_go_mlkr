package problem61

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}

	fast := head
	for i := 0; i < k; i++ {
		if fast.Next == nil {
			return rotateRight(head, k%(i+1))
		}

		fast = fast.Next
	}

	slow := head
	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}

	nowHead := slow.Next
	slow.Next, fast.Next = nil, head

	return nowHead
}
