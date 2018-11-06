package problem145

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LRD
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := []int{}
	leftRes := postorderTraversal(root.Left)
	rightRes := postorderTraversal(root.Right)

	res = append(res, leftRes...)
	res = append(res, rightRes...)
	res = append(res, root.Val)

	return res
}

// LRD(非递归解法)
// 		1
// 			2
// 		3
func postorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := []int{}
	nodes := []*TreeNode{}
	for root != nil || len(nodes) > 0 {

		if root == nil {
			root = nodes[len(nodes)-1]
			nodes = nodes[:len(nodes)-1]
		}

		if root.Left == nil && root.Right == nil {
			res = append(res, root.Val)
			root = nil
		} else {
			left, right := root.Left, root.Right
			root.Left, root.Right = nil, nil

			nodes = append(nodes, root)
			if right != nil {
				nodes = append(nodes, right)
			}

			root = left
		}

	}

	return res
}
