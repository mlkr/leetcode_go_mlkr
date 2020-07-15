package problem415

func addStrings(num1 string, num2 string) string {
	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}

	size1 := len(num1)
	size2 := len(num2)

	res := make([]byte, size2+1)
	res[0] = '1'

	carry := byte(0)
	for i := 1; i <= size2; i++ {
		if size1-i >= 0 {
			res[size2+1-i] = num1[size1-i] - '0'
		}

		res[size2+1-i] += num2[size2-i] + carry

		carry = byte(0)
		if res[size2+1-i] > '9' {
			carry = byte(1)
			res[size2+1-i] -= 10
		}
	}

	if carry == 1 {
		return string(res)
	}

	return string(res[1:])
}
