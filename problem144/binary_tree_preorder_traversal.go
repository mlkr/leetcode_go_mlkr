package problem144

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DLR
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := []int{root.Val}
	leftRes := preorderTraversal(root.Left)
	rightRes := preorderTraversal(root.Right)

	res = append(res, leftRes...)
	res = append(res, rightRes...)

	return res
}

// DLR(非递归解法)
func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := []int{}
	nodes := []*TreeNode{}
	for root != nil {
		res = append(res, root.Val)
		if root.Right != nil {
			nodes = append(nodes, root.Right)
		}

		root = root.Left
		if root == nil && len(nodes) > 0 {
			root = nodes[len(nodes)-1]
			nodes = nodes[:len(nodes)-1]
		}
	}

	return res
}
