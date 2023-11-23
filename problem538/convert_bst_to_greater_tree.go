package problem538

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	RDL(root, &sum)
	return root
}

func RDL(root *TreeNode, sum *int) {
	if root == nil {
		return
	}

	RDL(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	RDL(root.Left, sum)
}
