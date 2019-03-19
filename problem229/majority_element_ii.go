package problem229

func majorityElement(nums []int) []int {
	count := make(map[int]int)
	have := make(map[int]bool)
	size := len(nums) / 3
	res := []int{}
	for _, num := range nums {
		times := count[num]
		times++
		if times > size && !have[num] {
			have[num] = true
			res = append(res, num)
		}
		count[num] = times
	}

	return res
}

// 解法二
func majorityElement2(nums []int) []int {
	size := len(nums)
	if size <= 1 {
		return nums
	}

	// 出现次数大于 1/3 的数最多只有两个
	num1, num2, times1, times2 := 0, 1, 0, 0
	for _, num := range nums {
		switch {
		case num == num1:
			times1++
		case num == num2:
			times2++
		case times1 == 0:
			num1 = num
			times1 = 1
		case times2 == 0:
			num2 = num
			times2 = 1
		default:
			times1--
			times2--
		}
	}

	// 检查两个数是否出现多余 1/3
	times1, times2 = 0, 0
	for _, num := range nums {
		if num == num1 {
			times1++
			continue
		}

		if num == num2 {
			times2++
		}
	}

	res := make([]int, 0, 2)
	if times1 > size/3 {
		res = append(res, num1)
	}

	if times2 > size/3 {
		res = append(res, num2)
	}

	return res
}
