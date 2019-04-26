package problem263

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}

	if num == 1 {
		return true
	}

	switch {
	case num%2 == 0:
		return isUgly(num / 2)
	case num%3 == 0:
		return isUgly(num / 3)
	case num%5 == 0:
		return isUgly(num / 5)
	}

	return false
}
