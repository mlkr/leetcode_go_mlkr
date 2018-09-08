package problem102

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	recur(root, &res, 0)

	return res
}

func recur(root *TreeNode, res *[][]int, idx int) {
	if root == nil {
		return
	}

	if (len(*res) - 1) < idx {
		*res = append(*res, []int{})
	}

	temp := *res
	temp[idx] = append(temp[idx], root.Val)

	recur(root.Left, res, idx+1)
	recur(root.Right, res, idx+1)
}

// 另一中解法
func levelOrder2(root *TreeNode) [][]int {
	res := [][]int{}
	var dfs func(*TreeNode, int)

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		// 出现了新的 level
		if level == len(res) {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return res
}
