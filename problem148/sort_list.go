package problem148

import "math/rand"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 插入排序
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	headPre := &ListNode{
		Next: head,
	}

	for head.Next != nil {
		if head.Val <= head.Next.Val {
			head = head.Next
			continue
		}
		node := head.Next
		head.Next = head.Next.Next

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

// 快速排序
func sortList2(head *ListNode) *ListNode {
	nums := listToSlice(head)
	quickSort(nums)
	return sliceToList(nums)
}

// 归并排序 (最佳)
func sortList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	secHead := split(head)
	return merge(sortList3(head), sortList3(secHead))
}

// 二分
func split(head *ListNode) *ListNode {
	var pre *ListNode
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	pre.Next = nil
	return slow
}

// 合并
func merge(left, right *ListNode) *ListNode {
	cur := &ListNode{}
	headPre := cur

	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}

		cur = cur.Next
	}

	if left == nil {
		cur.Next = right
	} else {
		cur.Next = left
	}

	return headPre.Next
}

func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	devide := part(nums)
	quickSort(nums[:devide])
	quickSort(nums[devide+1:])
}

func part(nums []int) int {
	n := rand.Intn(len(nums))
	nums[0], nums[n] = nums[n], nums[0]

	end := len(nums) - 1
	i, j := 1, end
	for i < j {
		for nums[i] <= nums[0] && i < end {
			i++
		}

		for nums[0] < nums[j] && j > 1 {
			j--
		}

		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	devide := 0
	if nums[0] >= nums[j] {
		nums[0], nums[j] = nums[j], nums[0]
		devide = j
	}
	return devide
}

func sliceToList(s []int) *ListNode {
	sLen := len(s)
	if sLen == 0 {
		return nil
	}

	head := &ListNode{
		Val: s[0],
	}

	node := head
	for i := 1; i < sLen; i++ {
		node.Next = &ListNode{
			Val: s[i],
		}

		node = node.Next
	}

	return head
}

func listToSlice(head *ListNode) []int {
	if head == nil {
		return nil
	}

	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}
