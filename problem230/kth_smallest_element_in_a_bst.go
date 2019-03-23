package problem230

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	res := 0
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inorder(root.Left)
		k--
		if k == 0 {
			res = root.Val
			return
		}
		inorder(root.Right)

	}

	inorder(root)
	return res
}

// 第二个解法
func kthSmallest2(root *TreeNode, k int) int {
	leftSize := getSize(root.Left)

	switch {
	case leftSize >= k:
		return kthSmallest2(root.Left, k)
	case leftSize+1 < k:
		return kthSmallest2(root.Right, k-leftSize-1)
	default:
		return root.Val
	}
}

func getSize(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + getSize(root.Left) + getSize(root.Right)
}

// 解法三
func kthSmallest3(root *TreeNode, k int) int {
	c := make(chan int)
	go func() {
		Traverse(c, root)
		close(c)
	}()
	x, ok := <-c
	for ; k > 1 && ok; x, ok = <-c {
		k--
	}
	return x
}

func Traverse(c chan int, root *TreeNode) {
	if root == nil {
		return
	}
	Traverse(c, root.Left)
	c <- root.Val
	Traverse(c, root.Right)
}
