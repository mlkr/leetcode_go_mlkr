package problem129

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	sum := 0
	if root == nil {
		return sum
	}

	var dfs func(root *TreeNode, path []int)
	dfs = func(root *TreeNode, path []int) {
		path = append(path, root.Val)

		if root.Left == nil && root.Right == nil {
			sum += getNum(path)
			return
		}

		if root.Left != nil {
			dfs(root.Left, path)
		}

		if root.Right != nil {
			dfs(root.Right, path)
		}

	}

	dfs(root, []int{})

	return sum
}

func getNum(nums []int) int {
	size := len(nums)
	res := 0
	for k, v := range nums {
		res += v * pow10(size-1-k)
	}

	return res
}

func pow10(time int) int {
	if time == 0 {
		return 1
	}

	res := 1
	for i := 1; i <= time; i++ {
		res *= 10
	}

	return res
}
