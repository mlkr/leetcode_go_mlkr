package problem268

//  a^b^b =a
func missingNumber(nums []int) int {
	xor := 0
	for i := 0; i < len(nums); i++ {
		xor = xor ^ i ^ nums[i]
	}

	return xor ^ len(nums)
}
