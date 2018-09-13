package problem109

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	nums := l2s(head)

	return sortedArrayToBST(nums)
}

func sortedArrayToBST(nums []int) *TreeNode {
	numsLen := len(nums)
	if numsLen == 0 {
		return nil
	}

	mid := numsLen / 2

	return &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
}

func l2s(head *ListNode) []int {
	res := []int{}

	for head != nil {
		res = append(res, head.Val)

		head = head.Next
	}

	return res
}
