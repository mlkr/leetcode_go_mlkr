package problem42

func trap(height []int) int {
	len := len(height)

	if len <= 2 {
		return 0
	}

	left := make([]int, len, len)
	right := make([]int, len, len)

	left[0] = height[0]
	right[len-1] = height[len-1]

	for i := 1; i < len; i++ {
		left[i] = bigger(left[i-1], height[i])
		right[len-i-1] = bigger(right[len-i], height[len-i-1])
	}

	water := 0
	for k, v := range height {
		water += smaller(left[k], right[k]) - v
	}

	return water
}

func bigger(para1 int, para2 int) int {
	if para1 > para2 {
		return para1
	}

	return para2
}

func smaller(para1 int, para2 int) int {
	if para1 > para2 {
		return para2
	}

	return para1
}
