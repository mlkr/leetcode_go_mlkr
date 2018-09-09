package problem104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	max := 0
	if root == nil {
		return max
	}

	var getMaxDepth func(root *TreeNode, idx int)
	getMaxDepth = func(root *TreeNode, idx int) {
		if idx > max {
			max = idx
		}

		if root == nil {
			return
		}

		getMaxDepth(root.Left, idx+1)
		getMaxDepth(root.Right, idx+1)
	}

	getMaxDepth(root, 0)

	return max
}

// 不用闭包
func maxDepth3(root *TreeNode) int {
	max := 0
	if root == nil {
		return max
	}

	getMaxDepth(root, 0, &max)

	return max
}

func getMaxDepth(root *TreeNode, idx int, max *int) {
	if idx > *max {
		*max = idx
	}

	if root == nil {
		return
	}

	getMaxDepth(root.Left, idx+1, max)
	getMaxDepth(root.Right, idx+1, max)
}

// 另一解法
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth2(root.Left), maxDepth2(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
