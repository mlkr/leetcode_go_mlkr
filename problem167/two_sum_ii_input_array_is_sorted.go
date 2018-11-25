package problem167

func twoSum(numbers []int, target int) []int {
	res := []int{}
	size := len(numbers)

loop:
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if numbers[i]+numbers[j] == target {
				res = []int{i + 1, j + 1}
				break loop
			}
		}
	}

	return res
}

// 第二种解法(最佳)
func twoSum2(numbers []int, target int) []int {
	m := make(map[int]int)

	for k, v := range numbers {
		if m[target-v] != 0 {
			return []int{m[target-v], k + 1}
		}
		m[v] = k + 1
	}

	return nil
}
