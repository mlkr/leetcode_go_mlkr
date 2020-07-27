package problem424

func characterReplacement(s string, k int) int {
	count := [128]int{}
	max := 0
	left := 0
	for right, b := range s {
		count[b]++
		max = Max(max, count[b])

		// 不能覆盖
		if max+k < right-left+1 {
			count[s[left]]--
			left++
		}
	}

	return len(s) - left
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
