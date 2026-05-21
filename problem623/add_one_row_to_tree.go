package problem623

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		node := &TreeNode{
			Val:  val,
			Left: root,
		}

		return node
	}

	dfs(root, val, depth)
	return root
}

func dfs(root *TreeNode, val, depth int) {
	if root == nil {
		return
	}

	depth--
	if depth == 1 {
		left := &TreeNode{
			Val:  val,
			Left: root.Left,
		}

		right := &TreeNode{
			Val:   val,
			Right: root.Right,
		}

		root.Left = left
		root.Right = right
		return
	}

	dfs(root.Left, val, depth)
	dfs(root.Right, val, depth)
}
