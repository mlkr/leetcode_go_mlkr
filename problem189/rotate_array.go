package problem189

func rotate(nums []int, k int) {
	size := len(nums)
	if k > size {
		k %= size
	}

	if k == 0 {
		return
	}

	point := size - k
	left := nums[point:]
	left = append(left, nums[:point]...)
	for i := range nums {
		nums[i] = left[i]
	}
}

// 解法二
func rotate2(nums []int, k int) {
	size := len(nums)
	if k > size {
		k %= size
	}

	if k == 0 {
		return
	}

	reverse(nums, 0, size-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, size-1)
}

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

// 解法三(解法二修改版 最佳)
func rotate3(nums []int, k int) {
	size := len(nums)
	if k > size {
		k %= size
	}

	if k == 0 {
		return
	}

	for i := 0; i < size/2; i++ {
		nums[i], nums[size-1-i] = nums[size-1-i], nums[i]
	}

	for i := 0; i < k/2; i++ {
		nums[i], nums[k-1-i] = nums[k-1-i], nums[i]
	}

	for i := k; i < (size-k)/2+k; i++ {
		nums[i], nums[size-1-(i-k)] = nums[size-1-(i-k)], nums[i]
	}

}
