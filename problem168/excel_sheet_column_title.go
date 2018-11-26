package problem168

func convertToTitle(n int) string {
	res := []byte{}
	for {
		newN, b := convert(n)
		n = newN
		res = append([]byte{b}, res...)

		if n == 0 {
			break
		}
	}

	return string(res)
}

func convert(n int) (int, byte) {
	n--
	b := n % 26
	b += 65

	return n / 26, byte(b)
}

// 另一解法(最佳)
func convertToTitle2(n int) string {
	res := ""

	for n > 0 {
		n--
		res = string(byte(n%26)+'A') + res
		n /= 26
	}

	return res
}
