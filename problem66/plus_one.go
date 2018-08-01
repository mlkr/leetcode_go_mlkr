package problem66

// func plusOne(digits []int) []int {
// 	dLen := len(digits)
// 	if dLen == 0 {
// 		return []int{}
// 	}

// 	res := make([]int, dLen)

// 	carryFlag := true
// 	for i := range digits {
// 		if digits[i] != 9 {
// 			carryFlag = false
// 			break
// 		}
// 	}

// 	if carryFlag {
// 		dLen++
// 		res = make([]int, dLen)
// 		res[0] = 1

// 		return res
// 	}

// 	carray := 0
// 	for j := dLen - 1; j >= 0; j-- {
// 		temp := digits[j]

// 		if j == dLen-1 {
// 			temp++
// 		}

// 		if carray > 0 {
// 			temp += carray
// 		}

// 		carray = temp / 10
// 		res[j] = temp % 10

// 	}

// 	return res
// }

func plusOne(digits []int) []int {
	dLen := len(digits)
	if dLen == 0 {
		return []int{1}
	}

	digits[dLen-1]++

	for j := dLen - 1; j > 0; j-- {
		if digits[j] < 10 {
			break
		}

		digits[j] -= 10
		digits[j-1]++
	}

	if digits[0] > 9 {
		digits[0] -= 10
		digits = append([]int{1}, digits...)
	}

	return digits
}
