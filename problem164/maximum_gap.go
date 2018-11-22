package problem164

import (
	"math/rand"
)

func maximumGap(nums []int) int {
	size := len(nums)
	if size < 2 {
		return 0
	}

	quickSort(nums)
	max := 0
	for i := 1; i < size; i++ {
		if nums[i]-nums[i-1] > max {
			max = nums[i] - nums[i-1]
		}
	}

	return max
}

func quickSort(nums []int) {
	size := len(nums)
	if size <= 1 {
		return
	}

	n := part(nums)
	quickSort(nums[:n])
	quickSort(nums[n+1:])
}

func part(nums []int) int {
	size := len(nums)
	n := rand.Intn(size)

	nums[0], nums[n] = nums[n], nums[0]
	i, j := 1, size-1

	for i < j {
		for nums[i] <= nums[0] && i < size-1 {
			i++
		}

		for nums[0] < nums[j] && j > 1 {
			j--
		}

		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	pos := 0
	if nums[j] <= nums[0] {
		nums[0], nums[j] = nums[j], nums[0]
		pos = j
	}

	return pos
}
