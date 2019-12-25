package problem349

func intersection(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]struct{})
	for _, num := range nums1 {
		m1[num] = struct{}{}
	}

	m2 := make(map[int]struct{})
	for _, num := range nums2 {
		if _, ok := m1[num]; ok {
			m2[num] = struct{}{}
		}
	}

	res := []int{}
	for k := range m2 {
		res = append(res, k)
	}

	return res
}
