package problem546

func removeBoxes(boxes []int) int {
	n := len(boxes)
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
		}
	}

	l := 0
	r := n - 1
	k := 0

	res := remove(boxes, dp, l, r, k)
	return res
}

func remove(boxes []int, dp [][][]int, l, r, k int) int {
	if l > r {
		return 0
	}

	if dp[l][r][k] > 0 {
		return dp[l][r][k]
	}

	// 不能这样把r提前往前移动
	// dp[l][r][k] 会失去记录计算结果的作用, 导致重复计算
	// for l < r && boxes[r-1] == boxes[r] {
	// 	r--
	// 	k++
	// }

	// 直接去掉相同颜色的盒子，把右侧的 k 个元素和 boxes[r] 现在就清除掉
	res := (k+1)*(k+1) + remove(boxes, dp, l, r-1, 0)

	for i := r - 1; i >= l; i-- {
		if boxes[i] == boxes[r] {
			// 参看 "不能这样把r提前往前移动" 这个注释
			// if i+1 < r && boxes[i+1] == boxes[r] {
			// 	continue
			// }

			// 或者拼接前面相同颜色的盒子
			res = max(res, remove(boxes, dp, i+1, r-1, 0)+remove(boxes, dp, l, i, k+1))
		}
	}

	dp[l][r][k] = res
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
