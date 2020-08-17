package problem437

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	cnt := 0
	prefixTracker := make(map[int]int)
	prefixTracker[0] = 1

	var dfs func(root *TreeNode, curSum int)
	dfs = func(root *TreeNode, curSum int) {
		if root == nil {
			return
		}

		curSum += root.Val
		if count, ok := prefixTracker[curSum-sum]; ok {
			cnt += count
		}
		prefixTracker[curSum]++

		dfs(root.Left, curSum)
		dfs(root.Right, curSum)

		if prefixTracker[curSum] == 1 {
			delete(prefixTracker, curSum)
		} else {
			prefixTracker[curSum]--
		}
	}

	dfs(root, 0)

	return cnt
}
