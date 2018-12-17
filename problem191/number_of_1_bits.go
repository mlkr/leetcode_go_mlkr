package problem191

func hammingWeight(num uint32) int {
	count := 0
	for i := 0; i < 32; i++ {
		res := num & 1
		if res == 1 {
			count++
		}
		num = num >> 1
	}

	return count
}
