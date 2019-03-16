package problem227

func calculate(s string) int {
	// n1 opt1 n2 opt2 n3
	// ↓   ↓   ↓   ↓   ↓
	// 1   +   2   *   3
	var n1, n2, n3 int
	var opt1, opt2 byte
	idx := 0
	sLen := len(s)
	n1 = 0
	opt1 = '+'

	nextN := func() int {
		num := 0

		// 跳过空格
		for idx < sLen && s[idx] == ' ' {
			idx++
		}

		for idx < sLen && s[idx] >= '0' && s[idx] <= '9' {
			num = num*10 + int(s[idx]-'0')
			idx++
		}

		return num
	}

	nextOpt := func() byte {
		var opt byte
		opt = '+'
		// 跳过空格
		for idx < sLen && s[idx] == ' ' {
			idx++
		}

		if idx < sLen {
			opt = s[idx]
		}
		idx++

		return opt
	}

	n2 = nextN()
	for idx < sLen {
		opt2 = nextOpt()
		n3 = nextN()

		if opt2 == '*' || opt2 == '/' {
			n2 = operate(n2, n3, opt2)
		} else {
			n1 = operate(n1, n2, opt1)
			n2 = n3
			opt1 = opt2
		}
	}

	return operate(n1, n2, opt1)
}

func operate(a, b int, opt byte) int {
	res := 0
	switch opt {
	case '+':
		res = a + b
	case '-':
		res = a - b
	case '*':
		res = a * b
	case '/':
		res = a / b
	}

	return res
}
