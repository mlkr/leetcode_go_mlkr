package problem220

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t < 0 || k < 0 || len(nums) <= 1 {
		return false
	}

	t++
	used := make(map[int]int, k+1)

	for i, n := range nums {
		// 例:t=3,n=1或n=2, index=n/t都是同一个
		index := n / t
		if n < 0 {
			index--
		}

		if _, ok := used[index]; ok {
			return true
		}

		if val, ok := used[index-1]; ok && abs(n, val) < t {
			return true
		}

		if val, ok := used[index+1]; ok && abs(n, val) < t {
			return true
		}

		used[index] = n

		// 删除超出范围的节点
		if i >= k {
			key := nums[i-k] / t
			if nums[i-k] < 0 {
				key--
			}
			delete(used, key)
		}
	}

	return false
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}

	return y - x
}
