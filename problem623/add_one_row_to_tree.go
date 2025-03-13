package problem623

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		r := &TreeNode{
			Val:   val,
			Left:  root,
			Right: nil,
		}

		return r
	}

	dfs(root, val, depth)

	return root
}

func dfs(root *TreeNode, val int, depth int) {
	if root == nil {
		return
	}

	depth -= 1
	if depth == 1 {
		l := &TreeNode{
			Val:   val,
			Left:  root.Left,
			Right: nil,
		}

		r := &TreeNode{
			Val:   val,
			Left:  nil,
			Right: root.Right,
		}

		root.Left = l
		root.Right = r

		return
	}

	dfs(root.Left, val, depth)
	dfs(root.Right, val, depth)
}
