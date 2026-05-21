package problem645

func findErrorNums(nums []int) []int {
	sum := 0
	rec := make([]bool, len(nums)+1)
	dup := 0
	for _, num := range nums {
		if rec[num] {
			dup = num
		}
		rec[num] = true

		sum += num
	}

	return []int{dup, len(nums)*(len(nums)+1)/2 - sum + dup}
}
