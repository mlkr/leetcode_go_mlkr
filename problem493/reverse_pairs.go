package problem493

var res int
var temp []int

func reversePairs(nums []int) int {
	res = 0
	temp = make([]int, len(nums))
	sort(nums, 0, len(nums))

	return res
}

func sort(nums []int, begin, end int) {
	if begin+1 >= end {
		return
	}

	mid := (begin + end) / 2

	sort(nums, begin, mid)
	sort(nums, mid, end)

	copy(temp[begin:end], nums[begin:end])
	merge(nums, begin, mid, end)
}

func merge(nums []int, begin, mid, end int) {
	i, j := begin, mid
	for i < mid && j < end {
		if nums[i] > 2*nums[j] {
			res += mid - i
			j++
		} else {
			i++
		}
	}

	i, j, k := begin, mid, begin
	for i < mid && j < end {
		if temp[i] < temp[j] {
			nums[k] = temp[i]
			i++
		} else {
			nums[k] = temp[j]
			j++
		}
		k++
	}

	if i == mid {
		copy(nums[k:end], temp[j:end])
	}

	if j == end {
		copy(nums[k:end], temp[i:mid])
	}
}

// func reversePairs(nums []int) int {
// 	return count(nums, 0, len(nums))
// }

// func count(nums []int, begin, end int) int {
// 	if begin+1 >= end {
// 		return 0
// 	}

// 	mid := (begin + end) / 2
// 	res := count(nums, begin, mid) + count(nums, mid, end)

// 	i, j := begin, mid
// 	for i < mid && j < end {
// 		if nums[i] > 2*nums[j] {
// 			res += mid - i
// 			j++
// 		} else {
// 			i++
// 		}
// 	}

// 	quickSort(nums, begin, end-1)

// 	return res
// }

// func quickSort(nums []int, i, j int) {
// 	if i >= j {
// 		return
// 	}

// 	tempI, tempJ := i, j
// 	devide := i

// 	for i < j {
// 		for nums[j] >= nums[devide] && i < j {
// 			j--
// 		}

// 		for nums[i] <= nums[devide] && i < j {
// 			i++
// 		}

// 		nums[i], nums[j] = nums[j], nums[i]
// 	}

// 	nums[i], nums[devide] = nums[devide], nums[i]
// 	devide = i

// 	quickSort(nums, tempI, devide-1)
// 	quickSort(nums, devide+1, tempJ)
// }
