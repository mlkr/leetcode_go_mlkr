package problem624

// 想象成一个史莱姆吃另一个史莱姆，然后变大，返回最后的史莱姆的大小
func maxDistance(arrays [][]int) int {
	minN, maxN, diff := arrays[0][0], arrays[0][len(arrays[0])-1], 0

	for i := 1; i < len(arrays); i++ {
		arr := arrays[i]
		diff = max(diff, max(abs(arr[len(arr)-1]-minN), abs(maxN-arr[0])))

		maxN = max(maxN, arr[len(arr)-1])
		minN = min(minN, arr[0])
	}

	return diff
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}

	return -a
}
