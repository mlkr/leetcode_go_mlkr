package problem128

import (
	"sort"
)

func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	sort.Ints(nums)
	maxLength, length := 0, 1
	for i := 1; i < len(nums); i++ {
		if nums[i]-1 == nums[i-1] {
			length++
		} else if nums[i] != nums[i-1] {
			length = 1
		}

		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}

func longestConsecutive2(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	numsMap := make(map[int]bool, len(nums))
	for _, v := range nums {
		numsMap[v] = true
	}

	max := 0
	for _, v := range nums {
		length := 0
		temp := v
		for numsMap[v] {
			delete(numsMap, v)
			length++
			v++
		}

		v = temp - 1
		for numsMap[v] {
			delete(numsMap, v)
			length++
			v--
		}

		if length > max {
			max = length
		}
	}

	return max
}

// 最佳
func longestConsecutive3(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	max := 0
	numsMap := make(map[int]int, len(nums))
	for _, v := range nums {
		if numsMap[v] == 0 {
			left := numsMap[v-1]
			right := numsMap[v+1]

			length := left + right + 1

			numsMap[v-left] = length
			numsMap[v] = length
			numsMap[v+right] = length

			if length > max {
				max = length
			}
		}
	}

	return max
}
