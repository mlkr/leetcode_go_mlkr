package problem575

func distributeCandies(candyType []int) int {
	eats := len(candyType) >> 1
	m := make(map[int]bool, eats)
	for _, t := range candyType {
		m[t] = true
	}

	types := len(m)

	if types < eats {
		return types
	}

	return eats
}
