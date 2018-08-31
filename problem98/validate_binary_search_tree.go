package problem98

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	// 得二叉树的中序遍历
	in := getInorder(root)
	inLen := len(in)
	if inLen == 0 {
		return true
	}

	for i := 1; i < inLen; i++ {
		if in[i] <= in[i-1] {
			return false
		}
	}

	return true

}

// 得二叉树的中序遍历
func getInorder(root *TreeNode) []int {
	roots := []*TreeNode{}
	res := []int{}
	for root != nil || len(roots) != 0 {

		if root != nil {
			roots = append(roots, root)

			for root.Left != nil {
				root = root.Left
				roots = append(roots, root)
			}
		}

		root = roots[len(roots)-1]
		roots = roots[:len(roots)-1]

		res = append(res, root.Val)

		root = root.Right
	}

	return res
}
