package problem403

func canCross(stones []int) bool {
	size := len(stones)

	var backtrace func(idx, dif int) bool
	backtrace = func(idx, dif int) bool {
		if idx == 0 && dif == 1 {
			return true
		}

		for i := idx - 1; i >= 0; i-- {
			// x+1=dif dif-1=x
			// x=dif   dif=x
			// x-1=dif dif+1=x
			// dif-1 <= x <= dif+1
			if stones[idx]-stones[i] < dif-1 {
				continue
			}

			if stones[idx]-stones[i] > dif+1 {
				return false
			}

			if backtrace(i, stones[idx]-stones[i]) {
				return true
			}
		}

		return false
	}

	for i := size - 2; i >= 0; i-- {
		if backtrace(i, stones[size-1]-stones[i]) {
			return true
		}
	}

	return false
}
