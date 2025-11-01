package problem652

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	res := make([]*TreeNode, 0)
	count := make(map[string]int)

	helper(root, &res, count)

	return res
}

func helper(root *TreeNode, res *[]*TreeNode, count map[string]int) string {
	if root == nil {
		return "*"
	}

	left := helper(root.Left, res, count)
	right := helper(root.Right, res, count)
	key := strconv.Itoa(root.Val) + "," + left + "," + right

	count[key]++
	if count[key] == 2 {
		*res = append(*res, root)
	}

	return key
}
