package problem106

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	postLen := len(postorder)

	if postLen == 0 {
		return nil
	}

	root := &TreeNode{
		Val: postorder[postLen-1],
	}

	if postLen == 1 {
		return root
	}

	idx := indexOf(root.Val, inorder)

	root.Left = buildTree(inorder[:idx], postorder[:idx])
	root.Right = buildTree(inorder[idx+1:], postorder[idx:postLen-1])

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
