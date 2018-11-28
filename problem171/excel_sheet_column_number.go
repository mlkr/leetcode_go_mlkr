package problem171

func titleToNumber(s string) int {
	size := len(s)
	res := 0
	for i := 0; i < size; i++ {
		num := int(s[i] - 'A' + 1)
		res = res*26 + num
	}

	return res
}
