package problem257

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	res := make([]string, 0)

	var dfs func(node *TreeNode, path string)
	dfs = func(node *TreeNode, path string) {
		if path == "" {
			path = strconv.Itoa(node.Val)
		} else {
			path += "->" + strconv.Itoa(node.Val)
		}

		if node.Left != nil {
			dfs(node.Left, path)
		}

		if node.Right != nil {
			dfs(node.Right, path)
		}

		if node.Left == nil && node.Right == nil {
			res = append(res, path)
			return
		}
	}

	dfs(root, "")
	return res
}
