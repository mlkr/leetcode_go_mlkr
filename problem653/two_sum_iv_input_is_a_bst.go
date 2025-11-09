package problem653

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}

	nodes := make([]*TreeNode, 0)
	nodes = append(nodes, root)
	for len(nodes) > 0 {
		now := nodes[0]
		nodes = nodes[1:]

		if now.Val != k-now.Val && find(root, k-now.Val) {
			return true
		}

		if now.Left != nil {
			nodes = append(nodes, now.Left)
		}
		if now.Right != nil {
			nodes = append(nodes, now.Right)
		}

	}

	return false
}

func find(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}

	if root.Val == k {
		return true
	}

	if root.Val < k {
		return find(root.Right, k)
	}

	return find(root.Left, k)
}
