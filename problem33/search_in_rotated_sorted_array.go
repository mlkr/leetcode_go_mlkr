package problem33

import (
	"sort"
)

func search(nums []int, target int) int {
	nums_length := len(nums)
	var nums_new = make([]int, nums_length)
	copy(nums_new, nums)

	sort.Ints(nums_new)
	index_new := binary_search(nums_new, target, 0)
	if index_new == -1 {
		return -1
	}

	index_new_map := binary_search(nums_new, nums[index_new], 0)

	index_off := index_new - index_new_map
	if index_off < 0 {
		index_off = index_off + nums_length
	}
	return (index_new + index_off) % nums_length

}

func binary_search(nums []int, target int, drop int) int {
	nums_length := len(nums)
	if nums_length == 0 {
		return -1
	}
	mid := (nums_length - 1) / 2

	if nums[mid] == target {
		return drop + mid
	} else if nums[mid] < target {
		return binary_search(nums[mid+1:nums_length], target, drop+mid+1)
	} else {
		return binary_search(nums[0:mid], target, drop)
	}
}
