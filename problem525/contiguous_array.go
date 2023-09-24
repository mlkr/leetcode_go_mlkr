package problem525

func findMaxLength(nums []int) int {
	sum := 0
	res := 0
	m := make(map[int]int)
	for i, num := range nums {
		if num == 0 {
			sum--
		} else {
			sum++
		}

		// 0~i区间内出现的 0 1 数量相等
		if sum == 0 {
			res = i + 1
			continue
		}

		// pre~i区间内出现的 0 1 数量相等
		if pre, ok := m[sum]; ok {
			res = max(res, i-pre)
			continue
		}

		m[sum] = i
	}

	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
