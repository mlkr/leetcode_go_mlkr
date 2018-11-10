package problem147

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	sort := head
	head = head.Next
	if head == nil {
		return sort
	}
	sort.Next = nil

	for {
		if head == nil {
			break
		}

		node := head
		head = head.Next
		node.Next = nil
		sort = insert(sort, node)
	}

	return sort
}

func insert(sort, node *ListNode) *ListNode {
	head := sort
	if head.Val > node.Val {
		node.Next = head
		head = node
		return head
	}

	for {
		if sort.Next == nil || sort.Next.Val > node.Val {
			break
		}

		sort = sort.Next
	}

	if sort.Next != nil {
		node.Next = sort.Next
	}
	sort.Next = node

	return head
}

// 解法2(最佳)
func insertionSortList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	headPre := &ListNode{Next: head}
	cur := head
	for cur.Next != nil {
		if cur.Val <= cur.Next.Val {
			cur = cur.Next
			continue
		}
		node := cur.Next
		cur.Next = cur.Next.Next

		pre, preNext := headPre, headPre.Next
		for preNext.Val < node.Val {
			pre = preNext
			preNext = preNext.Next
		}

		pre.Next = node
		node.Next = preNext

	}

	return headPre.Next
}
