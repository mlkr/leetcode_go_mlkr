package problem530

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	res := 100000
	last := -1

	var ldr func(root *TreeNode)
	ldr = func(root *TreeNode) {
		if root.Left != nil {
			ldr(root.Left)
		}

		if last > -1 {
			res = min(res, diff(root.Val, last))
		}
		last = root.Val

		if root.Right != nil {
			ldr(root.Right)
		}
	}

	ldr(root)
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}
