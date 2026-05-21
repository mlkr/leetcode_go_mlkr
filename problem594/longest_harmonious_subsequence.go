package problem594

import "sort"

func findLHS(nums []int) int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}

	res := 0
	for n, t1 := range m {
		if t2, ok := m[n+1]; ok {
			res = max(res, t1+t2)
		}
	}

	return res
}

func findLHS2(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	sort.Ints(nums)

	length := 0
	left, right := 0, 0
	leftLen, rightLen := 0, 0
	for right < len(nums) {
		if nums[right] == nums[left] {
			leftLen++
		}

		if nums[right]-nums[left] == 1 {
			rightLen++
		}

		if nums[right]-nums[left] > 1 {
			if rightLen != 0 {
				length = max(length, leftLen+rightLen)
			}

			left += leftLen
			leftLen = rightLen

			rightLen = 0
			right--
		}

		right++
	}

	if rightLen != 0 {
		length = max(length, leftLen+rightLen)
	}

	return length
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
