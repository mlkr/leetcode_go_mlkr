package problem258

func addDigits(num int) int {
	res := 0
	for {
		for num > 0 {
			n := num % 10
			num /= 10
			res += n
		}

		if res < 10 {
			break
		}

		num = res
		res = 0
	}

	return res
}

// 解法二
// 假如有这么一个数ABCDE，
// 有ABCDE=A*10000+B*1000+C*100+D*10+E=(A+B+C+D+E)+(A*9999+B*999+C*99+D*9)，
// 那么ABCDE%9=(A+B+C+D+E)%9，这就很好的实现了ABCDE%9的结果为一位数，
// 但是这里还需要注意一点，若 (A+B+C+D+E) = 9，则结果为0。因而可用ABCDE%9=(A+B+C+D+E-1)%9+1
func addDigits2(num int) int {
	return (num-1)%9 + 1
}
