package problem108

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	numsLen := len(nums)
	if numsLen == 0 {
		return nil
	}

	mid := numsLen / 2

	root := &TreeNode{
		Val: nums[mid],
	}

	if numsLen == 1 {
		return root
	}

	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])

	return root
}
