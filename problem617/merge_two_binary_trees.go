package problem617

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	if root1 == nil && root2 == nil {
		return nil
	}

	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	if root1.Left == nil && root2.Left != nil {
		root1.Left = &TreeNode{Val: 0}
	}

	if root1.Right == nil && root2.Right != nil {
		root1.Right = &TreeNode{Val: 0}
	}

	// DLR
	root1.Val += root2.Val

	mergeTrees(root1.Left, root2.Left)
	mergeTrees(root1.Right, root2.Right)

	return root1
}

func mergeTrees2(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	// DLR
	root := &TreeNode{Val: root1.Val + root2.Val}
	root.Left = mergeTrees2(root1.Left, root2.Left)
	root.Right = mergeTrees2(root1.Right, root2.Right)

	return root
}
