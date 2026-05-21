package problem503

func nextGreaterElements(nums []int) []int {
	size := len(nums)
	res := make([]int, size)
	stack := make([]int, size)

	for i := 0; i < size; i++ {
		res[i] = -1
		stack[i] = i
	}

	for i := 0; i < size; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			res[stack[len(stack)-1]] = nums[i]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	return res
}
