package problem396

func maxRotateFunction(A []int) int {
	sum := 0
	sumF := 0
	size := len(A)

	for i := 0; i < size; i++ {
		sum += A[i]
		sumF += i * A[i]
	}

	res := sumF
	for i := size - 1; i > 0; i-- {
		sumF += sum - size*A[i]
		res = max(res, sumF)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
