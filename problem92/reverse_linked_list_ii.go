package problem92

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}

	// 前面增加一个节点, 防止 {3,5} m=1,n=2 这种情况 不能分离节点
	// 分离节点
	headPre := &ListNode{Next: head}
	mPre, nNext, mL := split(headPre, m+1, n+1)

	// 翻转
	mL, nL := reverse(mL)

	// 和原来分离的地方拼接起来
	mPre.Next = mL
	nL.Next = nNext

	return headPre.Next
}

// 分离要翻转的 ListNode
func split(head *ListNode, m, n int) (mPre, nNext, mL *ListNode) {
	for i := 1; i <= n; i++ {
		if i == m-1 {
			mPre = head
			mL = head.Next
		}

		if i == n {
			nNext = head.Next
			head.Next = nil
		}

		head = head.Next
	}

	return
}

// 翻转 ListNode
func reverse(head *ListNode) (h, e *ListNode) {
	if head == nil || head.Next == nil {
		return head, head
	}

	h, end := reverse(head.Next)
	end.Next = head
	e = head

	return
}
