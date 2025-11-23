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
	cols := int(math.Pow(2, float64(rows))) - 1
	res := make([][]string, rows)
	for i := range res {
		res[i] = make([]string, cols)
	}

	nodes := []*TreeNode{}
	nodes = append(nodes, root)

	sss := [][]string{}
	sss = append(sss, res[0])

	helper := Helper{}
	helper.r = 0      // 在第几行
	helper.begin = 0  // 起点
	helper.end = cols // 终点

	helpers := []Helper{}
	helpers = append(helpers, helper)

	for len(nodes) > 0 {
		now := nodes[0]
		nodes = nodes[1:]

		ss := sss[0]
		sss = sss[1:]
		assign(now, ss)

		helper = helpers[0]
		helpers = helpers[1:]

		r := helper.r
		begin := helper.begin
		end := helper.end
		mid := (begin + end) / 2

		if now.Left != nil {
			helper = Helper{}
			helper.r = r + 1
			helper.begin = begin
			helper.end = mid

			helpers = append(helpers, helper)
			nodes = append(nodes, now.Left)
			sss = append(sss, res[helper.r][helper.begin:helper.end])
		}

		if now.Right != nil {
			helper = Helper{}
			helper.r = r + 1
			helper.begin = mid + 1
			helper.end = end

			helpers = append(helpers, helper)
			nodes = append(nodes, now.Right)
			sss = append(sss, res[helper.r][helper.begin:helper.end])
		}

	}

	return res
}

func assign(root *TreeNode, row []string) {
	mid := len(row) / 2
	row[mid] = strconv.Itoa(root.Val)
}

func getRows(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftR := getRows(root.Left)
	rightR := getRows(root.Right)

	r := leftR
	if rightR > r {
		r = rightR
	}

	return r + 1
}
