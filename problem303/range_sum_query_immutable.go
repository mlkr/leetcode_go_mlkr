package problem303

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	return NumArray{nums: nums}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.nums[j]
	}

	return this.nums[j] - this.nums[i-1]
}
