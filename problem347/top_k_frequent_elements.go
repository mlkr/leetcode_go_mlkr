package problem347

type candidates []int

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 1 {
		return nums
	}

	// 统计次数
	record := make(map[int]int, len(nums))
	for _, num := range nums {
		record[num]++
	}

	counts := make([]candidates, len(nums))
	for num, count := range record {
		counts[count-1] = append(counts[count-1], num)
	}

	res := make([]int, 0, k)

loop:
	for i := len(counts) - 1; i >= 0; i-- {
		if k == 0 {
			break
		}

		for _, num := range counts[i] {
			res = append(res, num)
			k--

			if k == 0 {
				break loop
			}

		}
	}

	return res
}
