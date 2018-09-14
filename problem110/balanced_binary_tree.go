package problem110

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	sub := getDep(root.Left) - getDep(root.Right)
	return sub >= -1 && sub <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func getDep(root *TreeNode) int {
	max := 0
	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			if level > max {
				max = level
			}

			return
		}

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return max
}

// 另一解法
func isBalanced2(root *TreeNode) bool {
	_, is := recur(root)
	return is
}

func recur(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftDep, leftIs := recur(root.Left)
	rightDep, rightIs := recur(root.Right)

	sub := leftDep - rightDep
	if leftIs && rightIs && sub <= 1 && sub >= -1 {
		return max(leftDep, rightDep) + 1, true
	}

	return 0, false
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
