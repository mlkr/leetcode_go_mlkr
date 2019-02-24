package problem219

func containsNearbyDuplicate(nums []int, k int) bool {
	if k <= 0 {
		return false
	}

	used := make(map[int]int, len(nums))
	for i, n := range nums {
		index, ok := used[n]
		if ok && i-index <= k {
			return true
		}

		used[n] = i
	}

	return false
}
