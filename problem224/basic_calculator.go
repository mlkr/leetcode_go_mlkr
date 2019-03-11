package problem224

func calculate(s string) int {
	sLen := len(s)
	if sLen == 0 {
		return 0
	}

	stack := make([]int, 0, sLen)
	res := 0
	sign := 1
	for i := 0; i < sLen; i++ {
		switch s[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num := 0
			for ; i < sLen && s[i] >= '0' && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			i--

			res += sign * num
		case '+':
			sign = 1
		case '-':
			sign = -1
		case '(':
			stack = append(stack, res, sign)
			res = 0
			sign = 1
		case ')':
			sign = stack[len(stack)-1]
			lastRes := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			res = lastRes + sign*res
		}
	}

	return res
}
