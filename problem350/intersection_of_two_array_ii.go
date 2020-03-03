package problem350

func intersect(nums1 []int, nums2 []int) []int {
	m1 := map[int]int{}
	for _, num := range nums1 {
		m1[num] = m1[num] + 1
	}

	res := []int{}
	for _, num := range nums2 {
		if v, ok := m1[num]; ok && v > 0 {
			res = append(res, num)
			m1[num] = v - 1
		}
	}

	return res
}
