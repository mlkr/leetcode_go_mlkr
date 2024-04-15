package problem559

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	depth := 0
	for _, c := range root.Children {
		depth = max(depth, maxDepth(c))
	}

	return depth + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
