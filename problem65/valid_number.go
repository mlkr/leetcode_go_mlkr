package problem65

func isNumber(s string) bool {
	sLen := len(s)
	if sLen == 0 {
		return false
	}

	b := ([]byte)(s)

	// 去除前后空格
	front := 0
	for i := 0; i < sLen; i++ {
		if b[i] == 32 {
			front++
			continue
		}
		break
	}

	back := sLen - 1
	for j := sLen - 1; j >= 0; j-- {
		if b[j] == 32 {
			back--
			continue
		}
		break
	}

	if front > back {
		return false
	}

	b = b[front : back+1]

	pointNum := 0
	lastPoint := make([]int, 1)

	negativeNum := 0
	lastNegative := make([]int, 1)

	plusNum := 0
	lastPlus := make([]int, 1)

	for i := range b {

		// +, -, ., 0-9, e, 不在这个范围返回false
		if (b[i] < 48 && b[i] != 43 && b[i] != 45 && b[i] != 46) || (b[i] > 57 && b[i] != 101) {
			return false
		}

		// e, 科学表达式
		if b[i] == 101 {
			return isEnum(string(b))
		}

		// +号
		if b[i] == 43 {
			lastPlus[0] = i
			plusNum++
		}

		// 负号
		if b[i] == 45 {
			lastNegative[0] = i
			negativeNum++
		}

		// 小数点
		if b[i] == 46 {
			lastPoint[0] = i
			pointNum++
		}

	}

	// 检查 +号
	if plusNum > 0 {
		// 多于一个 +号
		if plusNum > 1 {
			return false
		}

		plusPos := lastPlus[0]
		// +号前面还有数字
		if len(b[:plusPos]) > 0 {
			return false
		}

		return isNumber(string(b[plusPos+1:]))
	}

	// 检查负号
	if negativeNum > 0 {
		// 多于一个负号
		if negativeNum > 1 {
			return false
		}

		negativePos := lastNegative[0]
		// 负号前面还有数字
		if len(b[:negativePos]) > 0 {
			return false
		}

		// 除了负号,没有别的
		b = b[negativePos+1:]
		if len(b) == 0 {
			return false
		}

		return isNumber(string(b))
	}

	// 检查小数点
	if pointNum > 0 {
		// 多于一个小数点
		if pointNum > 1 {
			return false
		}

		pointPos := lastPoint[0]
		// 除了小数点,没有别的
		b = append(b[:pointPos], b[pointPos+1:]...)
		if len(b) == 0 {
			return false
		}
	}

	return true
}

// 是不是科学记数
func isEnum(s string) bool {
	sLen := len(s)
	if sLen == 0 {
		return false
	}

	b := ([]byte)(s)

	eNum := 0
	lastE := make([]int, 1)

	for i := range b {

		// +, -, ., 0-9, e, 不在这个范围返回false
		if (b[i] < 48 && b[i] != 43 && b[i] != 45 && b[i] != 46) || (b[i] > 57 && b[i] != 101) {
			return false
		}

		// e, 科学表达式
		if b[i] == 101 {
			eNum++
			lastE[0] = i
		}

	}

	// 检查 e
	if eNum > 0 {
		// 多于一个 e
		if eNum > 1 {
			return false
		}

		ePos := lastE[0]
		if !isNumber(string(b[:ePos])) || !isInteger(string(b[ePos+1:])) {
			return false
		}
	}

	return true
}

// 是不是整数
func isInteger(s string) bool {
	sLen := len(s)
	if sLen == 0 {
		return false
	}

	b := ([]byte)(s)

	negativeNum := 0
	lastNegative := make([]int, 1)

	plusNum := 0
	lastPlus := make([]int, 1)

	for i := range b {

		// +, -, ., 0-9, e, 不在这个范围返回false
		if (b[i] < 48 && b[i] != 43 && b[i] != 45) || (b[i] > 57) {
			return false
		}

		// +号
		if b[i] == 43 {
			lastPlus[0] = i
			plusNum++
		}

		// 负号
		if b[i] == 45 {
			lastNegative[0] = i
			negativeNum++
		}

	}

	// 检查 +号
	if plusNum > 0 {
		// 多于一个 +号
		if plusNum > 1 {
			return false
		}

		plusPos := lastPlus[0]
		// +号前面还有数字
		if len(b[:plusPos]) > 0 {
			return false
		}

		return isInteger(string(b[plusPos+1:]))
	}

	// 检查负号
	if negativeNum > 0 {
		// 多于一个负号
		if negativeNum > 1 {
			return false
		}

		negativePos := lastNegative[0]
		// 负号前面还有数字
		if len(b[:negativePos]) > 0 {
			return false
		}

		// 除了负号,没有别的
		b = b[negativePos+1:]
		if len(b) == 0 {
			return false
		}

		return isInteger(string(b))
	}

	return true
}
