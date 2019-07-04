package problem306

import "strconv"

func isAdditiveNumber(num string) bool {
	size := len(num)
	if size < 3 {
		return false
	}

	res := false

	var dfs func(num string, pre int)
	dfs = func(num string, pre int) {

		for i := 1; i < len(num); i++ {
			if res {
				return
			}

			// 没有前导0
			if i >= 2 && num[0] == '0' {
				return
			}

			cur, _ := strconv.Atoi(num[:i])
			remain, _ := strconv.Atoi(num[i:])
			if pre+cur == remain && size != len(num) {
				if len(num[i:]) < 2 || num[i] != '0' {
					res = true
				}
				return
			}

			if size == len(num) {
				dfs(num[i:], cur)
			}

			want := pre + cur
			digit := getDigit(want)
			if i+digit > len(num) {
				break
			}

			// 没有前导0
			if digit >= 2 && num[i] == '0' {
				continue
			}

			next, _ := strconv.Atoi(num[i : i+digit])
			if next == want {
				dfs(num[i:], cur)
			}
		}
	}

	dfs(num, 0)
	return res
}

// 得数字的位数
func getDigit(num int) int {
	digit := 0
	for {
		digit++
		num /= 10
		if num == 0 {
			break
		}
	}

	return digit
}

// 解法二
func isAdditiveNumber2(num string) bool {
	size := len(num)
	if size < 3 {
		return false
	}

	var dfs func(n1, n2 int, num string) bool
	dfs = func(n1, n2 int, num string) bool {
		n3 := n1 + n2
		s3 := strconv.Itoa(n3)
		step := len(s3)

		if num == s3 {
			return true
		}

		if len(num) < step || num[:step] != s3 {
			return false
		}

		return dfs(n2, n3, num[step:])
	}

	jMax := size / 3 * 2
	for i := 1; i <= jMax-1; i++ {
		// 没有前导0
		if i > 1 && num[0] == '0' {
			return false
		}

		for j := i + 1; j <= jMax; j++ {
			// 没有前导0
			if j-i > 1 && num[i] == '0' {
				break
			}

			n1, _ := strconv.Atoi(num[:i])
			n2, _ := strconv.Atoi(num[i:j])

			if dfs(n1, n2, num[j:]) {
				return true
			}
		}
	}

	return false
}
