package problem496

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	stack := make([]int, 0, len(nums2))

	for _, n := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < n {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			m[top] = n
		}

		stack = append(stack, n)
	}

	for _, n := range stack {
		m[n] = -1
	}

	res := make([]int, len(nums1))
	for i := range nums1 {
		res[i] = m[nums1[i]]
	}

	return res
}
