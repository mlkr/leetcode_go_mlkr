package problem32

import "sort"

func LongestValidParentheses(s string) int {
	s_length := len(s)
	if s_length == 0 {
		return 0
	}

	break_point := []int{-1, s_length}

	// 从左到右确定断点
	var left int
	for k, v := range s {

		if v == '(' {
			left++
		} else if left > 0 {
			left--
		} else {
			break_point = append(break_point, k)
		}
	}

	// 从右到左确定断点
	var right int
	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == ')' {
			right++
		} else if right > 0 {
			right--
		} else {
			break_point = append(break_point, i)
		}
	}

	// 得断点间的最大距离
	sort.Ints(break_point)
	max := 0
	count := 0
	for i := 1; i < len(break_point); i++ {
		count = break_point[i] - break_point[i-1]
		if count > max {
			max = count
		}
	}

	return max - 1
}
