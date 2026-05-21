package problem572

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	if isEque(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)

}

func isEque(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return subRoot == nil
	}

	if subRoot == nil {
		return false
	}

	if root.Val != subRoot.Val {
		return false
	}

	return isEque(root.Left, subRoot.Left) &&
		isEque(root.Right, subRoot.Right)
}
