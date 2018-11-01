package problem143

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil ||
		head.Next == nil ||
		head.Next.Next == nil {
		return
	}

	last := head
	prev := head
	for {
		if last.Next == nil {
			break
		}

		prev = last
		last = last.Next
	}

	nextHead := head.Next
	head.Next = last
	prev.Next = nil
	last.Next = nextHead
	reorderList(nextHead)
}

// 另一解法
func reorderList2(head *ListNode) {
	if head == nil ||
		head.Next == nil ||
		head.Next.Next == nil {
		return
	}

	// 计算链表长度
	cur := head
	size := 1
	for {
		if cur.Next == nil {
			break
		}
		cur = cur.Next
		size++
	}

	// 移动到链表的中间
	// size 为奇数， cur 指向 list 的中间节点
	// size 为偶数， cur 指向 list 前一半的最后一个节点
	cur = head
	mid := (size - 1) / 2
	for i := 0; i < mid; i++ {
		cur = cur.Next
	}

	// 后半部链表方向取反
	// 1->2->3->4
	// 1->2<->3<-4
	next := cur.Next
	for next != nil {
		temp := next.Next
		next.Next = cur
		cur = next
		next = temp
	}

	// 处理链表
	// 1->2<->3<-4
	// 1->4->2->3
	end := cur
	cur = head
	for i := 0; i < mid; i++ {
		nextHead := cur.Next
		nextEnd := end.Next
		cur.Next = end
		end.Next = nextHead
		cur = nextHead
		end = nextEnd
	}

	end.Next = nil
}
