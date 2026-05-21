package problem515

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	que := []*TreeNode{root}
	res := []int{}
	n := 1
	max := math.MinInt64

	for n > 0 {
		root = que[0]
		que = que[1:]
		n--

		if root.Val > max {
			max = root.Val
		}

		if root.Left != nil {
			que = append(que, root.Left)
		}

		if root.Right != nil {
			que = append(que, root.Right)
		}

		if n == 0 {
			n = len(que)
			res = append(res, max)
			max = math.MinInt64
		}

	}

	return res
}
