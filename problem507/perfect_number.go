package problem507

func checkPerfectNumber(num int) bool {
	if num%2 != 0 {
		return false
	}

	sum := 1
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			sum += i + num/i
		}
	}

	return sum == num
}
