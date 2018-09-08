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
