package problem206

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	headNext := head.Next
	head.Next = nil
	re := reverseList(headNext)

	temp := re
	for re.Next != nil {
		re = re.Next
	}
	re.Next = head

	return temp
}

// 解法二(最佳)
func reverseList2(head *ListNode) *ListNode {
	var pre *ListNode

	for head != nil {
		temp := head.Next
		head.Next = pre
		pre = head
		head = temp
	}

	return pre
}
