package problem35

func searchInsert(nums []int, target int) int {
	low, mid, height := 0, 0, len(nums)-1

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

	if nums[mid] > target {
		return mid
	} else {
		return mid + 1
	}
}
