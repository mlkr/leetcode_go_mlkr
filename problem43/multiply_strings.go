package problem43

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	num1B := []byte(num1)
	num2B := []byte(num2)

	temp := make([]int, len(num1B)+len(num2B))

	for i := 0; i < len(num1B); i++ {
		for j := 0; j < len(num2B); j++ {
			temp[i+j+1] += (int)(num1B[i]-'0') * (int)(num2B[j]-'0')
		}
	}

	for k := len(temp) - 1; k > 0; k-- {
		temp[k-1] += temp[k] / 10
		temp[k] = temp[k] % 10
	}

	if temp[0] == 0 {
		temp = temp[1:]
	}

	res := make([]byte, len(temp))
	for k, v := range temp {
		res[k] = byte(v) + '0'
	}

	return string(res)

}
