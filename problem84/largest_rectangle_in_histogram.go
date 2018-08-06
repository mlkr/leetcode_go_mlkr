package problem84

func largestRectangleArea(heights []int) int {
	h := append(heights, -1)
	hLen := len(h)
	left, right := 0, 0
	stack := []int{}
	maxArea := 0

	// 从左到右遍历, 如果现在的值比前一个小, 计算前面的直方图的面积,直到下一个直方图比现在的小
	// 在末尾添加了 -1, 强制计算前面的面积
	for right < hLen {
		if len(stack) == 0 || h[stack[len(stack)-1]] <= h[right] {
			stack = append(stack, right)
			right++
			continue
		}

		height := h[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]

		if len(stack) == 0 {
			left = -1
		} else {
			left = stack[len(stack)-1]
		}

		// 如果当前值比前面的值小, 计算前面的面积
		area := (right - left - 1) * height
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}
