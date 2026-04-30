package problem671

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	// 遍历树
	min, secMin := root.Val, 1<<63-1

	var dlr func(root *TreeNode)
	dlr = func(root *TreeNode) {
		if root == nil {
			return
		}

		if min < root.Val && root.Val < secMin {
			secMin = root.Val
		}

		dlr(root.Left)
		dlr(root.Right)
	}

	dlr(root)

	if secMin == 1<<63-1 {
		return -1
	}

	return secMin
}
