package problem75

// func sortColors(nums []int) {
// 	quickSort(nums)
// }

// func quickSort(nums []int) {
// 	nLen := len(nums)
// 	if nLen <= 1 {
// 		return
// 	}

// 	k := rand.Intn(nLen)
// 	nums[k], nums[0] = nums[0], nums[k]
// 	point := part(nums)

// 	quickSort(nums[:point])
// 	quickSort(nums[point+1:])
// }

// func part(nums []int) int {
// 	nLen := len(nums)

// 	i := 1
// 	j := nLen - 1

// 	for {
// 		for nums[i] <= nums[0] && i < nLen-1 {
// 			i++
// 		}

// 		for nums[0] <= nums[j] && j > 0 {
// 			j--
// 		}

// 		if i >= j {
// 			break
// 		}

// 		nums[i], nums[j] = nums[j], nums[i]

// 	}

// 	nums[0], nums[j] = nums[j], nums[0]

// 	return j
// }

func sortColors(nums []int) {
	numLen := len(nums)
	if numLen == 0 {
		return
	}

	temp := nums[0]
	nums[0] = 1

	i := 0
	j := 1
	k := numLen - 1
	for j <= k {
		switch {
		case nums[j] < 1:
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		case 1 < nums[j]:
			nums[j], nums[k] = nums[k], nums[j]
			k--
		default:
			j++
		}
	}

	switch temp {
	case 0:
		nums[i] = 0
	case 2:
		nums[k] = 2
	}

	return
}
