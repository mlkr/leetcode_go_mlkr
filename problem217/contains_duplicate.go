package problem217

func containsDuplicate(nums []int) bool {
	used := make(map[int]bool, len(nums))
	for _, num := range nums {
		if used[num] {
			return true
		}

		used[num] = true
	}

	return false
}
