package problem540

func singleNonDuplicate(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := (hi-lo)/2 + lo
		if mid%2 == 1 {
			mid--
		}

		if nums[mid] == nums[mid+1] {
			// 所找的元素肯定不在左边
			lo = mid + 2
		} else {
			// 所找的元素肯定在左边
			hi = mid
		}
	}

	return nums[lo]
}
