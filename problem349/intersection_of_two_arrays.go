package problem349

func intersection(nums1 []int, nums2 []int) []int {
	m1 := map[int]bool{}
	for _, num := range nums1 {
		m1[num] = true
	}

	res := []int{}
	for _, num := range nums2 {
		if v, ok := m1[num]; ok && v {
			res = append(res, num)
			m1[num] = false
		}
	}

	return res
}
