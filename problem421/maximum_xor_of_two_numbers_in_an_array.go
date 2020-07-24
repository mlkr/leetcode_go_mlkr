package problem421

// 参看
// https://kingsfish.github.io/2017/12/15/Leetcode-421-Maximum-XOR-of-Two-Numbers-in-an-Array/
func findMaximumXOR(nums []int) int {
	max, mask := 0, 0

	maxNum := 0
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}

	length := 31
	for i := 0; i < 32; i++ {
		if 1<<i > maxNum {
			length = i - 1
			break
		}
	}

	for i := length; i >= 0; i-- {
		mask |= 1 << i

		m := make(map[int]bool)
		for _, num := range nums {
			m[num&mask] = true
		}

		tmp := max | 1<<i
		for i := range m {
			if m[tmp^i] {
				max = tmp
				break
			}
		}
	}

	return max
}
