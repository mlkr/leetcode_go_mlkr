package problem390

func lastRemaining(n int) int {
	leftStart := true
	min := 1
	step := 1

	for n>1 {
		if leftStart || n%2 == 1{
			min += step
		}

		leftStart = !leftStart
		step *= 2
		n /=2
	}

	return min
}
