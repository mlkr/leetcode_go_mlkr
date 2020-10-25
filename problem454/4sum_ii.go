package problem454

func fourSumCount(A []int, B []int, C []int, D []int) int {
	size := len(A)
	res := 0
	if size == 0 {
		return res
	}

	m := make(map[int]int, size*size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			m[A[i]+B[j]]++
		}
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			res += m[-(C[i] + D[j])]
		}
	}

	return res
}
