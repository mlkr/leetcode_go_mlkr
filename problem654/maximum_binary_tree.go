package problem654

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	max := 0
	index := 0
	for i, num := range nums {
		if num > max {
			max = num
			index = i
		}
	}

	root := &TreeNode{}
	root.Val = max
	root.Left = constructMaximumBinaryTree(nums[:index])
	root.Right = constructMaximumBinaryTree(nums[index+1:])

	return root
}
