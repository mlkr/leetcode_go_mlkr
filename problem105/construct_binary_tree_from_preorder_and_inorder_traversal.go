package problem105

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	preLen := len(preorder)

	if preLen == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	if preLen == 1 {
		return root
	}

	idx := indexOf(root.Val, inorder)

	root.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	root.Right = buildTree(preorder[idx+1:], inorder[idx+1:])

	return root
}

func indexOf(n int, nums []int) int {
	for k, v := range nums {
		if v == n {
			return k
		}
	}

	return -1
}
