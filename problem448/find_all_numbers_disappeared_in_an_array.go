package problem448

func findDisappearedNumbers(nums []int) []int {
	size := len(nums)
	for i := 0; i < size; i++ {
		// x != nums[x-1]
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	res := make([]int, 0, size)
	for i, num := range nums {
		if i+1 != num {
			res = append(res, i+1)
		}
	}

	return res
}
