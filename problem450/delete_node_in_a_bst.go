package problem450

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {

	nums := []int{}
	var preOrder func(root *TreeNode)
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}

		nums = append(nums, root.Val)
		preOrder(root.Left)
		preOrder(root.Right)
	}

	preOrder(root)
	i := 0
	for ; i < len(nums); i++ {
		if nums[i] == key {
			break
		}
	}

	if i == len(nums) {
		return root
	}

	nums = append(nums[:i], nums[i+1:]...)

	return preToBST(nums)
}

func preToBST(pre []int) *TreeNode {

	idx := 0
	var buildTree func(min, max int) *TreeNode
	buildTree = func(min, max int) *TreeNode {
		if idx == len(pre) {
			return nil
		}

		val := pre[idx]
		if val < min || val > max {
			return nil
		}

		idx++
		root := &TreeNode{Val: val}
		root.Left = buildTree(min, val)
		root.Right = buildTree(val, max)

		return root
	}

	return buildTree(-1<<31, (1<<31)-1)
}

// 最佳
func deleteNode2(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode2(root.Left, key)
		return root
	}

	if key > root.Val {
		root.Right = deleteNode2(root.Right, key)
		return root
	}

	if root.Right == nil {
		return root.Left
	}

	rightSmallest := root.Right
	for rightSmallest.Left != nil {
		rightSmallest = rightSmallest.Left
	}
	rightSmallest.Left = root.Left

	return root.Right
}
