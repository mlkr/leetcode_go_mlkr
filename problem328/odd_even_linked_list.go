package problem328

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	oddNode, evenHead, evenNode := head, head.Next, head.Next
	preNode := oddNode

	for {
		// 奇节点
		oddNode.Next = oddNode.Next.Next
		preNode = oddNode
		oddNode = oddNode.Next
		if oddNode == nil {
			break
		}

		// 偶节点
		evenNode.Next = evenNode.Next.Next
		evenNode = evenNode.Next
		if evenNode == nil {
			break
		}

	}

	if oddNode == nil {
		oddNode = preNode
	}
	oddNode.Next = evenHead

	return head
}
