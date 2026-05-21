package problem606

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DLR
func tree2str(root *TreeNode) string {
	res := ""
	if root == nil {
		return res
	}

	res += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return res
	}

	res += "(" + tree2str(root.Left) + ")"

	if root.Right != nil {
		res += "(" + tree2str(root.Right) + ")"
	}

	return res
}
