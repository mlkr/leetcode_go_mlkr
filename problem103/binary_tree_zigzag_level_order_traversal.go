package problem103

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	var dfs func(root *TreeNode, idx int)

	dfs = func(root *TreeNode, idx int) {
		if root == nil {
			return
		}

		if len(res) == idx {
			res = append(res, []int{})
		}

		res[idx] = append(res[idx], root.Val)

		dfs(root.Left, idx+1)
		dfs(root.Right, idx+1)
	}

	dfs(root, 0)

	for k, v := range res {
		if k%2 != 0 {
			reverse(v)
		}
	}

	return res
}

func reverse(nums []int) {
	start := 0
	end := len(nums) - 1
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// 另一解法
func zigzagLevelOrder2(root *TreeNode) [][]int {
	res := [][]int{}
	var dfs func(root *TreeNode, idx int)

	dfs = func(root *TreeNode, idx int) {
		if root == nil {
			return
		}

		if len(res) == idx {
			res = append(res, []int{})
		}

		if idx%2 == 0 {
			res[idx] = append(res[idx], root.Val)
		} else {
			temp := []int{root.Val}
			res[idx] = append(temp, res[idx]...)
		}

		dfs(root.Left, idx+1)
		dfs(root.Right, idx+1)
	}

	dfs(root, 0)

	return res
}

func zigzagLevelOrder3(root *TreeNode) [][]int {
	res := [][]int{}
	dfs(root, 0, &res)

	return res
}

func dfs(root *TreeNode, idx int, res *[][]int) {
	if root == nil {
		return
	}

	if len(*res) == idx {
		*res = append(*res, []int{})
	}

	ar := *res
	if idx%2 == 0 {
		ar[idx] = append(ar[idx], root.Val)
	} else {
		temp := []int{root.Val}
		ar[idx] = append(temp, ar[idx]...)
	}

	dfs(root.Left, idx+1, res)
	dfs(root.Right, idx+1, res)
}
