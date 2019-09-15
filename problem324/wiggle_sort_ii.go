package problem324

// 参看 https://leetcode.com/problems/wiggle-sort-ii/discuss/77682/Step-by-step-explanation-of-index-mapping-in-Java
func wiggleSort(nums []int) {
	size := len(nums)
	pivot := findKthLargest(nums, (size+1)/2)

	i := 0
	left, right := 0, size-1
	for i <= right {
		switch {
		case nums[newIndex(i, size)] > pivot:
			nums[newIndex(i, size)], nums[newIndex(left, size)] = nums[newIndex(left, size)], nums[newIndex(i, size)]
			i++
			left++
		case nums[newIndex(i, size)] < pivot:
			nums[newIndex(i, size)], nums[newIndex(right, size)] = nums[newIndex(right, size)], nums[newIndex(i, size)]
			right--
		default:
			i++
		}
	}
}

func newIndex(i, size int) int {
	return (i*2 + 1) % (size | 1)
}

// 取第 k 个大的数 (215 题)
func findKthLargest(nums []int, k int) int {
	nums = append([]int{}, nums...)
	return partition(nums, k)
}

func partition(nums []int, k int) int {
	size := len(nums)
	mid := size / 2
	pivot := nums[mid]
	nums[0], nums[mid] = nums[mid], nums[0]

	n := size - 1
	for i := size - 1; i > 0; i-- {
		if nums[i] >= pivot {
			nums[i], nums[n] = nums[n], nums[i]
			n--
		}
	}

	nums[0], nums[n] = nums[n], nums[0]

	if size-n == k {
		return pivot
	}

	if size-n > k {
		nums = append([]int{}, nums[n+1:]...)
		return partition(nums, k)
	}

	k = k - (size - n)
	nums = append([]int{}, nums[:n]...)

	return partition(nums, k)
}
