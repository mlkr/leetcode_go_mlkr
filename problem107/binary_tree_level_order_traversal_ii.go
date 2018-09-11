package problem107

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}

	var getLevelOrder func(root *TreeNode, level int)
	getLevelOrder = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if len(res) == level {
			res = append(res, []int{})
		}

		res[level] = append(res[level], root.Val)

		getLevelOrder(root.Left, level+1)
		getLevelOrder(root.Right, level+1)
	}

	getLevelOrder(root, 0)

	reverse(res)

	return res
}

func reverse(nums [][]int) {
	i, j := 0, len(nums)-1

	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

// 另一解法
func levelOrderBottom2(root *TreeNode) [][]int {
	res := [][]int{}

	var getLevelOrder func(root *TreeNode, level int)
	getLevelOrder = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if len(res) == level {
			res = append([][]int{[]int{}}, res...)
		}

		resLen := len(res)
		res[resLen-level-1] = append(res[resLen-level-1], root.Val)

		getLevelOrder(root.Left, level+1)
		getLevelOrder(root.Right, level+1)
	}

	getLevelOrder(root, 0)

	return res
}
