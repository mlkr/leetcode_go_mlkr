package problem337

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	var dfs func(root *TreeNode) (int, int)

	// a 为抢劫该节点的最大值
	// b 为不抢劫该节点的最大值
	dfs = func(root *TreeNode) (a, b int) {
		if root == nil {
			return 0, 0
		}

		la, lb := dfs(root.Left)
		ra, rb := dfs(root.Right)

		a = root.Val + lb + rb
		b = max(la, lb) + max(ra, rb)

		return a, b
	}

	a, b := dfs(root)
	return max(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
