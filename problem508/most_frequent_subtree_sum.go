package problem508

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	sumCount := make(map[int]int)
	maxCount := 0

	var getSubTreeSum func(root *TreeNode) int
	getSubTreeSum = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		sum := root.Val + getSubTreeSum(root.Left) + getSubTreeSum(root.Right)
		sumCount[sum]++

		if sumCount[sum] > maxCount {
			maxCount = sumCount[sum]
		}

		return sum
	}

	getSubTreeSum(root)

	res := make([]int, 0, len(sumCount))
	for sum, count := range sumCount {
		if maxCount == count {
			res = append(res, sum)
		}
	}

	return res
}
