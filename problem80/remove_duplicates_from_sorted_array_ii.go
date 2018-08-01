package problem80

func removeDuplicates(nums []int) int {

	// 用两个指针，将第i个新出现的数字，放倒第j个上
	i, j := 0, 0
	numsLen := len(nums)

	for ; i < numsLen; i++ {
		if i < 2 || nums[i] > nums[j-2] {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}
