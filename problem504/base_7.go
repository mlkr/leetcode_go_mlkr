package problem504

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	isNeg := false
	if num < 0 {
		num = -num
		isNeg = true
	}

	res := []byte{}
	for num != 0 {
		res = append(res, byte(num%7)+'0')
		num /= 7
	}

	if isNeg {
		res = append(res, '-')
	}

	reverse(res)

	return string(res)
}

func reverse(b []byte) {
	size := len(b)
	i := 0
	j := size - 1

	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
}
