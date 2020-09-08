package problem445

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n1 := list2Num(l1)
	n2 := list2Num(l2)

	size1 := len(n1)
	size2 := len(n2)
	if size1 > size2 {
		n1, n2 = n2, n1
		size1, size2 = size2, size1
	}

	res := byte(0)
	head := &ListNode{}
	for i := 1; i <= size1 || i <= size2; i++ {
		if i <= size1 {
			res += n1[size1-i]
		}

		if i <= size2 {
			res += n2[size2-i]
		}

		node := &ListNode{Val: int(res)%10}
		if head.Next != nil {
			node.Next = head.Next
		}
		head.Next = node

		res /= 10
	}

	if res == 1 {
		pre := head.Next
		head.Next = &ListNode{Val:1, Next:pre}
	}

	return head.Next
}

func list2Num(root *ListNode) []byte {
	res := []byte{}
	for root != nil {
		res = append(res, byte(root.Val))
		root = root.Next
	}

	return res
}