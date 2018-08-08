package problem86

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	small := &ListNode{}
	big := &ListNode{}
	part := big
	begin := small

	for {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			big.Next = head
			big = big.Next
		}

		if head.Next == nil {
			break
		}

		head = head.Next
	}

	small.Next = part.Next
	big.Next = nil

	return begin.Next
}
