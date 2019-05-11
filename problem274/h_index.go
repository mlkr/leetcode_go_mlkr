package problem274

import (
	"sort"
)

func hIndex(citations []int) int {
	quickSort(citations)
	left, right := 0, len(citations)-1

	// 从大到小排
	for left < right {
		citations[left], citations[right] = citations[right], citations[left]
		left++
		right--
	}

	// 二分查找
	lo, hi := 0, len(citations)
	for lo <= hi {
		mi := (lo + hi) / 2
		if citations[mi] > mi {
			lo = mi + 1
		} else {
			hi = mi - 1
		}
	}

	return lo
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

// 解法2
func hIndex2(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))

	// 二分查找
	lo, hi := 0, len(citations)-1
	for lo <= hi {
		mi := (lo + hi) / 2
		if citations[mi] > mi {
			lo = mi + 1
		} else {
			hi = mi - 1
		}
	}

	return lo
}
