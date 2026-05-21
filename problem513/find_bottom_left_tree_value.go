package problem513

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	que := []*TreeNode{root}

	for len(que) > 0 {
		root = que[0]
		que = que[1:]

		if root.Right != nil {
			que = append(que, root.Right)
		}

		if root.Left != nil {
			que = append(que, root.Left)
		}

	}

	return root.Val
}
