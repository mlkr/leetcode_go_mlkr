package problem94

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := inorderTraversal(root.Left)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)

	return res
}

// 非递归解法
func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	roots := []*TreeNode{}
	res := []int{}
	for len(roots) != 0 || root != nil {

		if root != nil {
			roots = append(roots, root)

			for root.Left != nil {
				root = root.Left
				roots = append(roots, root)
			}
		}

		node := roots[len(roots)-1:]
		roots = roots[:len(roots)-1]
		res = append(res, node[0].Val)

		root = node[0].Right
	}

	return res
}
