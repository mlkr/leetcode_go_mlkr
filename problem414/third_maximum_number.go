package problem414

import "math"

func thirdMax(nums []int) int {
	first, second, third := math.MinInt64, math.MinInt64, math.MinInt64

	for _, num := range nums {
		if num == first || num == second {
			continue
		}

		switch {
		case num > first:
			first, second, third = num, first, second
		case num > second:
			second, third = num, second
		case num > third:
			third = num
		}
	}

	if third == math.MinInt64 {
		return first
	}

	return third
}
