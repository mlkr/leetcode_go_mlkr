package problem561

func arrayPairSum(nums []int) int {
	nums = sortArr(nums)

	res := 0
	for i := 0; i < len(nums); i += 2 {
		res += nums[i]
	}

	return res
}

// 计数排序
// https://www.runoob.com/w3cnote/counting-sort.html
func sortArr(nums []int) []int {
	// 计算出现的次数
	max, min := getMaxMin(nums)
	freq := make([]int, max-min+1)
	for _, n := range nums {
		freq[n-min]++
	}

	// 累计
	for i := 1; i < len(freq); i++ {
		freq[i] += freq[i-1]
	}

	// 根据次数排序
	res := make([]int, len(nums))
	for _, n := range nums {
		freq[n-min]--
		res[freq[n-min]] = n
	}

	return res
}

func getMaxMin(nums []int) (int, int) {
	max, min := -1000, 1000
	for _, n := range nums {
		if n > max {
			max = n
		}

		if n < min {
			min = n
		}
	}

	return max, min
}
