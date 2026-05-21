package problem560

func subarraySum(nums []int, k int) int {
	res, sum := 0, 0
	sumCount := make(map[int]int, len(nums))
	sumCount[0] = 1
	for i := range nums {
		sum += nums[i]
		res += sumCount[sum-k]
		sumCount[sum]++
	}

	return res
}
