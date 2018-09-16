package problem112

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	is := false

	var runPath func(root *TreeNode, count int)
	runPath = func(root *TreeNode, count int) {
		if root == nil || is == true {
			return
		}

		if root.Left == nil && root.Right == nil {
			count += root.Val
			if count == sum {
				is = true
			}

			return
		}

		runPath(root.Left, count+root.Val)
		runPath(root.Right, count+root.Val)
	}

	runPath(root, 0)

	return is
}

// 另一解法
func hasPathSum2(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	sum -= root.Val

	if root.Left == nil && root.Right == nil {
		return sum == 0
	}

	return hasPathSum2(root.Left, sum) || hasPathSum2(root.Right, sum)
}
