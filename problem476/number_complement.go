package problem476

func findComplement(num int) int {
	res := 0
	i := 0

	for num > 0 {
		bit := num & 1
		bit ^= 1
		res |= bit << i

		i++
		num >>= 1
	}

	return res
}
