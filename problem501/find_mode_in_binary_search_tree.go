package problem501

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	rec := map[int]int{}
	search(root, rec)

	max := -1
	res := []int{}
	for v, r := range rec {
		if max <= r {
			if max < r {
				max = r
				res = res[:0]
			}
			res = append(res, v)
		}
	}

	return res
}

func search(root *TreeNode, rec map[int]int) {
	if root == nil {
		return
	}

	rec[root.Val]++
	search(root.Left, rec)
	search(root.Right, rec)
}
