package problem81

func search(nums []int, target int) bool {
	numsLen := len(nums)
	if numsLen == 0 {
		return false
	}

	i := 1
	for ; i < numsLen; i++ {
		if nums[i] < nums[i-1] {
			break
		}
	}

	return binarySearch(nums[:i], target) || binarySearch(nums[i:], target)
}

func binarySearch(nums []int, target int) bool {
	numsLen := len(nums)
	if numsLen == 0 {
		return false
	}

	mid := numsLen / 2

	if nums[mid] == target {
		return true
	}

	if nums[mid] > target {
		return binarySearch(nums[:mid], target)
	}

	return binarySearch(nums[mid+1:], target)
}
