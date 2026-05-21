package problem655

import (
	"math"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Helper struct {
	r, begin, end int
}

func printTree(root *TreeNode) [][]string {
	if root == nil {
		return nil
	}

	rows := getRows(root)
	cols := int(math.Pow(float64(2), float64(rows))) - 1
	res := make([][]string, rows)
	for i := 0; i < rows; i++ {
		res[i] = make([]string, cols)
	}

	nodes := []*TreeNode{root}
	helpers := []Helper{Helper{r: 0, begin: 0, end: cols}}

	for len(nodes) > 0 {
		now := nodes[0]
		nodes = nodes[1:]

		helper := helpers[0]
		helpers = helpers[1:]

		slicePrint(res[helper.r][helper.begin:helper.end], now)
		mid := (helper.begin + helper.end) / 2

		if now.Left != nil {
			nodes = append(nodes, now.Left)
			helpers = append(helpers, Helper{r: helper.r + 1, begin: helper.begin, end: mid})
		}

		if now.Right != nil {
			nodes = append(nodes, now.Right)
			helpers = append(helpers, Helper{r: helper.r + 1, begin: mid + 1, end: helper.end})
		}

	}

	return res
}

func slicePrint(s []string, root *TreeNode) {
	s[len(s)/2] = strconv.Itoa(root.Val)
}

func getRows(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(getRows(root.Left), getRows(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
