package problem532

import "sort"

func findPairs(nums []int, k int) int {
	sort.Ints(nums)
	size := len(nums)

	res := 0
	for i := 0; i < size-1; i++ {
		// 去掉重复的
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		want := nums[i] + k
		if j := binarySearch(want, nums, i+1); j != -1 {
			res++
		}

	}

	return res
}

func binarySearch(want int, nums []int, left int) int {
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == want {
			return mid
		}

		if nums[mid] < want {
			left = mid + 1
		}

		if nums[mid] > want {
			right = mid - 1
		}
	}

	return -1
}

func findPairs2(nums []int, k int) int {
	rec := make(map[int]int)
	for _, n := range nums {
		rec[n]++
	}

	res := 0
	if k == 0 {
		for n := range rec {
			if rec[n] > 1 {
				res++
			}
		}

		return res
	}

	for n := range rec {
		if rec[n-k] > 0 {
			res++
		}
	}

	return res
}
