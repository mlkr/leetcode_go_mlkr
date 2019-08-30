package problem318

func maxProduct(words []string) int {
	size := len(words)

	masks := make([]int, size)
	for i := 0; i < size; i++ {
		for _, b := range words[i] {
			masks[i] |= (1 << uint32(b-'a'))
		}
	}

	max := 0
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if masks[i]&masks[j] != 0 {
				continue
			}

			product := len(words[i]) * len(words[j])
			if product > max {
				max = product
			}
		}
	}

	return max
}
