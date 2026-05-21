package problem543

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, res := helper(root)
	return res
}

// leng 为root节点到叶子节点最长距离所经过的节点数
func helper(root *TreeNode) (leng, diam int) {
	if root == nil {
		return
	}

	leftLeng, leftDiam := helper(root.Left)
	rightLeng, rightDiam := helper(root.Right)

	leng = max(leftLeng, rightLeng) + 1
	diam = max(max(leftDiam, rightDiam), leftLeng+rightLeng)

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
