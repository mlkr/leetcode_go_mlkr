package problem384

import "math/rand"

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.nums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	size := len(this.nums)
	nums := make([]int, size)
	copy(nums, this.nums)

	for i := range nums {
		idx := rand.Intn(size)
		nums[i], nums[idx] = nums[idx], nums[i]
	}

	return nums
}
