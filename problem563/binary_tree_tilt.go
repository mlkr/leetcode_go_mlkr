package problem563

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var tiltSum int

func findTilt(root *TreeNode) int {
	tiltSum = 0
	getSumAndTilt(root)

	return tiltSum
}

func getSumAndTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftSum := getSumAndTilt(root.Left)
	rightSum := getSumAndTilt(root.Right)

	sum := root.Val + leftSum + rightSum
	tiltSum += getAbs(leftSum, rightSum)

	return sum
}

func getAbs(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}
