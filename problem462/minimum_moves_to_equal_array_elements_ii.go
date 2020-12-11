package problem462

import (
	"math/rand"
	"sort"
)

func minMoves2(nums []int) int {
	sort.Ints(nums)
	size := len(nums)
	mid := nums[size/2]

	if size&1 == 0 {
		// 数量为偶数
		mid += (nums[size/2-1] - mid) / 2
	}

	// 向中位数靠拢是步骤最少的
	res := 0
	for _, n := range nums {
		res += diff(mid, n)
	}

	return res
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}

// 更快的解法
func minMoves22(nums []int) int {
	size := len(nums)
	mid := size / 2
	midVal := quickSelect(nums, mid+1)

	if size&1 == 0 {
		// 偶数
		midVal += (quickSelect(nums, mid) - midVal) / 2
	}

	res := 0
	for _, n := range nums {
		res += diff(midVal, n)
	}

	return res
}

func quickSelect(nums []int, midSize int) int {
	size := len(nums)
	if size == 1 {
		return nums[0]
	}

	l, r := 0, size-1
	pivot := nums[rand.Intn(size)]
	pCount := 0

	for l <= r {
		for l <= r && nums[l] <= pivot {
			if nums[l] == pivot {
				pCount++
			}
			l++
		}

		for l <= r && nums[r] > pivot {
			r--
		}

		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}

	if l >= midSize && l-pCount < midSize {
		return pivot
	}

	if l-pCount >= midSize {
		return quickSelect(nums[:l], midSize)
	}

	return quickSelect(nums[l:], midSize-l)
}
