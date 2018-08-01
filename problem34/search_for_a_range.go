package problem34

func searchRange(nums []int, target int) []int {
	index := binarySearch(nums, target)
	len := len(nums)

	if index == -1 {
		return []int{-1, -1}
	}

	pre := index
	next := index

	for pre > 0 && nums[index] == nums[pre-1] {
		pre--
	}

	for next < len-1 && nums[index] == nums[next+1] {
		next++
	}

	return []int{pre, next}

}

func binarySearch(nums []int, target int) int {
	len := len(nums)
	low, mid, height := 0, 0, len-1

	for low <= height {
		mid = (low + height) / 2

		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] > target:
			height = mid - 1
		case nums[mid] < target:
			low = mid + 1
		}
	}

	return -1
}
