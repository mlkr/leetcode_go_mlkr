package problem482

func licenseKeyFormatting(s string, k int) string {
	b := []byte(s)
	counter := 0
	const diff = byte('a' - 'A')
	res := []byte{}

	for i := len(s) - 1; i >= 0; i-- {
		if b[i] == '-' {
			continue
		}

		if b[i] >= 'a' && b[i] <= 'z' {
			b[i] = b[i] - diff
		}

		if counter == k {
			res = append(res, '-')
			counter = 0
		}

		res = append(res, b[i])
		counter++
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return string(res)
}
