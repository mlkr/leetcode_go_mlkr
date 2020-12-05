package problem461

func hammingDistance(x int, y int) int {
	xor := x ^ y
	res := 0
	for xor > 0 {
		res += xor & 1
		xor >>= 1
	}

	return res
}
