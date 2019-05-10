package problem274

func hIndex(citations []int) int {
	quickSort(citations)
	left, right := 0, len(citations)-1

	// 从大到小排
	for left < right {
		citations[left], citations[right] = citations[right], citations[left]
		left++
		right--
	}

	i := 0
	for ; i < len(citations); i++ {
		if i+1 > citations[i] {
			break
		}
	}

	return i
}

func quickSort(nums []int) {
	size := len(nums)
	if size <= 1 {
		return
	}

	left, right := 0, size-1
	base := nums[0]
	i := 1
	for i <= right {
		if nums[i] >= base {
			nums[i], nums[right] = nums[right], nums[i]
			right--
		}

		if nums[i] < base {
			nums[i], nums[left] = nums[left], nums[i]
			left++
			i++
		}
	}

	quickSort(nums[:left])
	quickSort(nums[left+1:])
}
